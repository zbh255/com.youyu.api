// 签钥服务，全面统一项目的错误处理
package option

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/auth"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/path"
	"com.youyu.api/lib/utils"
	"context"
	"encoding/json"
	encrypt "github.com/abingzo/go-encrypt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type SecretKeyApiServer struct {
	rpc.UnimplementedCentApiServer
	// redis 连接池
	RedisClient *redis.Client
	// BindToken Mutex
	BindTokenLock sync.Mutex
	// 日志接口
	Logger log.Logger
}

const (
	KeyTimeout = time.Hour * 24
	// KeyLoginAndSignTimeOut 注册和登录保存的密钥副本超时时间
	KeyLoginAndSignTimeOut     = time.Minute * 10
	KeyNum                 int = 10
)

type KeyJSON struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

// NOTE: 废弃方法
// GetUserSigningKey 从redis数据库中获取存储的用户签名密钥,没有则添加
func (s *SecretKeyApiServer) GetUserSigningKey(uid string) (string, error) {
	if result, err := s.RedisClient.Get(context.Background(), path.SigningKeyPrefix+uid).Result(); err == redis.Nil {
		signingKey := utils.CreateSigningKey(uid)
		if err = s.RedisClient.Set(context.Background(), path.SigningKeyPrefix+uid, signingKey, 0).Err(); err != nil {
			s.Logger.Error(errors.Wrap(err, "get user signing key failed"))
			return "", err
		} else {
			return signingKey, nil
		}
	} else {
		return result, nil
	}
}

// 默认策略一个账号对应一个登录状态
func (s *SecretKeyApiServer) BindTokenToUser(ctx context.Context, user *rpc.User) (*rpc.User, error) {
	s.BindTokenLock.Lock()
	defer s.BindTokenLock.Unlock()
	// TODO:签名密钥直接获取，原有Api已废弃
	signingKey := TokenSigntureKey
	// 为该用户生成一个token
	authJwt := auth.New(signingKey)
	token, err := authJwt.GetToken(&auth.MyClaims{
		Uid: int64(user.Uid),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: user.ExpTime,
			Issuer:    "test",
		},
	})
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.User{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	if err := s.RedisClient.Set(context.Background(), path.TokenKeyPrefix+token, strconv.FormatInt(int64(user.Uid), 10), time.Duration(user.ExpTime)).Err(); err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.User{}, status.Error(ecode.RedisServerErr,ecode.RedisServerErr.Message())
	}
	return &rpc.User{
		Code:    int32(ecode.OK.Code()),
		Message: ecode.OK.Message(),
		Uid:     user.Uid,
		ExpTime: user.ExpTime,
		Token:   token,
	}, nil
}

func (s *SecretKeyApiServer) ForTokenGetBindUser(ctx context.Context, user *rpc.User) (*rpc.User, error) {
	if result, err := s.RedisClient.Get(context.Background(), path.TokenKeyPrefix+user.Token).Result(); err == redis.Nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.User{}, status.Error(ecode.AccessTokenExpires,ecode.AccessTokenExpires.Message())
	} else {
		uid, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			s.Logger.Error(errors.WithStack(err))
			return &rpc.User{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
		}
		return &rpc.User{
			Code:    int32(ecode.OK.Code()),
			Message: ecode.OK.Message(),
			Uid:     int32(uid),
			ExpTime: user.ExpTime,
			Token:   user.Token,
		}, nil
	}
}

func (s *SecretKeyApiServer) DeleteBindUser(ctx context.Context, user *rpc.User) (*rpc.Null, error) {
	if err := s.RedisClient.Del(context.Background(), path.TokenKeyPrefix+user.Token).Err(); err == redis.Nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.Null{}, status.Error(ecode.AccessTokenExpires,ecode.AccessTokenExpires.Message())
	} else {
		return &rpc.Null{}, nil
	}
}

// 公钥从公钥池中派发
// 客户端id -> 私钥
func (s *SecretKeyApiServer) GetPublicKey(ctx context.Context, null *rpc.RsaKey) (*rpc.RsaKey, error) {
	// 查询是否有空余的密钥
	// 没有则添加10个密钥
	if s.RedisClient.Get(context.Background(), path.PubAndPriKeyPrefix+"1").Err() == redis.Nil {
		for i := 1; i <= KeyNum; i++ {
			if rsa := encrypt.NewCoder().GetEncrypted().RsaCoder(encrypt.BitSize1024, nil, nil).CreateKeyPairPem(); rsa.Err() != nil {
				s.Logger.Error(errors.WithStack(rsa.Err()))
				return &rpc.RsaKey{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
			} else {
				// json
				result, err := json.Marshal(&KeyJSON{
					PublicKey:  string(rsa.GetPublicKeyPemBytes()),
					PrivateKey: string(rsa.GetPrivateKeyPemBytes()),
				})
				if err != nil {
					s.Logger.Error(errors.WithStack(err))
					return &rpc.RsaKey{}, status.Error(ecode.JsonParseError,ecode.JsonParseError.Message())
				}
				err = s.RedisClient.Set(context.Background(), path.PubAndPriKeyPrefix+strconv.Itoa(i), string(result), KeyTimeout).Err()
				if err != nil {
					s.Logger.Error(err)
					return &rpc.RsaKey{},status.Error(ecode.RedisServerErr,ecode.RedisServerErr.Message())
				}
			}
		}
	}
	// 添加完在获取密钥
	rand.Seed(time.Now().UnixNano())
	result, err := s.RedisClient.Get(context.Background(), path.PubAndPriKeyPrefix+strconv.Itoa(rand.Intn(10)+1)).Result()
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.RsaKey{}, status.Error(ecode.RedisServerErr,ecode.RedisServerErr.Message())
	}
	// 分解密钥对
	keyJson := &KeyJSON{}
	err = json.Unmarshal([]byte(result), &keyJson)
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.RsaKey{}, status.Error(ecode.JsonParseError,ecode.JsonParseError.Message())
	}
	// 创建副本以供查询
	// prefix+client_id -> privateKey
	s.RedisClient.Set(context.Background(), path.PubAndPriKeyPrefix+null.ClientId, keyJson.PrivateKey, KeyLoginAndSignTimeOut)
	return &rpc.RsaKey{PublicKey: keyJson.PublicKey, ClientId: null.ClientId}, nil
}

// key : value
// prefix+client_id : privateKey
func (s *SecretKeyApiServer) GetPrivateKey(ctx context.Context, key *rpc.RsaKey) (*rpc.RsaKey, error) {
	privateKey, err := s.RedisClient.Get(context.Background(), path.PubAndPriKeyPrefix+key.ClientId).Result()
	// 不存在则添加,调用完成即刻退出
	if err == redis.Nil {
		_, _ = s.GetPublicKey(context.Background(), &rpc.RsaKey{})
		s.Logger.Error(errors.Wrap(err,"client_id get private key failed"))
		return &rpc.RsaKey{}, status.Error(ecode.SecretKeyTimeout,ecode.SecretKeyTimeout.Message())
	}
	return &rpc.RsaKey{PublicKey: key.PublicKey, PrivateKey: privateKey, ClientId: key.ClientId}, nil
}

// code 不为0请勿绑定
func (s *SecretKeyApiServer) BindWechatToken(ctx context.Context, info *rpc.WechatTokenInfo) (*rpc.User, error) {
	bytes, err := json.Marshal(info)
	if err != nil {
		return &rpc.User{},err
	}
	// 生成一个token
	signingKey := TokenSigntureKey
	// 为该用户生成一个token
	sha512Key, err := encrypt.NewCoder().GetAbstract().Sha512Coder(encrypt.BASE64).SetJoinStr(";").
		Append(string(signingKey)).Append(time.Now().String()).Append(info.SessionKey).Result()
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.User{},status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	if err = s.RedisClient.Set(context.Background(), path.TokenKeyWechatLogin + string(sha512Key),string(bytes),KeyLoginAndSignTimeOut).Err(); err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.User{},status.Error(ecode.RedisServerErr,ecode.RedisServerErr.Message())
	}
	return &rpc.User{Token: string(sha512Key),ExpTime: int64(KeyLoginAndSignTimeOut)},nil
}

func (s *SecretKeyApiServer) ForWechatTokenGetInfo(ctx context.Context, user *rpc.User) (*rpc.WechatTokenInfo, error) {
	wechatToken,err := s.RedisClient.Get(context.Background(),path.TokenKeyWechatLogin + user.Token).Result()
	// token不存在则返回
	if err == redis.Nil {
		return &rpc.WechatTokenInfo{}, status.Error(ecode.AccessTokenExpires,ecode.AccessTokenExpires.Message())
	}
	jsons := rpc.WechatTokenInfo{}
	err = json.Unmarshal([]byte(wechatToken),&jsons)
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.WechatTokenInfo{},status.Error(ecode.JsonParseError,ecode.JsonParseError.Message())
	} else {
		return &jsons,nil
	}
}

func (s *SecretKeyApiServer) BindUserVcCode(ctx context.Context, code *rpc.UserVcCode) (*rpc.Null, error) {
	// 参数校验
	err := code.Validate()
	if err != nil {
		return &rpc.Null{},status.Error(ecode.ParaMeterErr,err.Error())
	}
	err = s.RedisClient.Set(context.Background(),path.VcCodePrefix + code.BindInfo,code.VcCode,KeyLoginAndSignTimeOut).Err()
	if err != nil {
		return &rpc.Null{},status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	} else {
		return &rpc.Null{},nil
	}
}

// 不进行参数校验，因为使用不到全部参数
func (s *SecretKeyApiServer) GetUserVcCode(ctx context.Context, code *rpc.UserVcCode) (*rpc.UserVcCode, error) {
	result,err := s.RedisClient.Get(context.Background(),path.VcCodePrefix + code.BindInfo).Result()
	if err == redis.Nil {
		return &rpc.UserVcCode{},status.Error(ecode.VcCodeTimeout,ecode.VcCodeTimeout.Message())
	} else if err != nil {
		return &rpc.UserVcCode{},status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	} else {
		return &rpc.UserVcCode{BindInfo: code.BindInfo,VcCode: result},nil
	}
}
