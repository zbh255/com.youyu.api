package secretKey_rpc

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/auth"
	"com.youyu.api/lib/ecode"
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
		return &rpc.User{}, err
	}
	if err := s.RedisClient.Set(context.Background(), path.TokenKeyPrefix+token, strconv.FormatInt(int64(user.Uid), 10), time.Duration(user.ExpTime)).Err(); err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.User{}, err
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
		return &rpc.User{}, err
	} else {
		uid, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			s.Logger.Error(errors.WithStack(err))
			return &rpc.User{}, err
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
		return &rpc.Null{}, err
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
				return &rpc.RsaKey{}, rsa.Err()
			} else {
				// json
				result, err := json.Marshal(&KeyJSON{
					PublicKey:  string(rsa.GetPublicKeyPemBytes()),
					PrivateKey: string(rsa.GetPrivateKeyPemBytes()),
				})
				if err != nil {
					s.Logger.Error(errors.WithStack(err))
					return nil, err
				}
				s.RedisClient.Set(context.Background(), path.PubAndPriKeyPrefix+strconv.Itoa(i), string(result), KeyTimeout)
			}
		}
	}
	// 添加完在获取密钥
	rand.Seed(time.Now().UnixNano())
	result, err := s.RedisClient.Get(context.Background(), path.PubAndPriKeyPrefix+strconv.Itoa(rand.Intn(10)+1)).Result()
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.RsaKey{}, err
	}
	// 分解密钥对
	keyJson := &KeyJSON{}
	err = json.Unmarshal([]byte(result), &keyJson)
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.RsaKey{}, err
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
		return &rpc.RsaKey{}, err
	}
	return &rpc.RsaKey{PublicKey: key.PublicKey, PrivateKey: privateKey, ClientId: key.ClientId}, nil
}
