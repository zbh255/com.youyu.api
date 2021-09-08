// 封装了一些Jwt处理的方法
package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type Jwt interface {
	GetToken(claims *MyClaims) (string, error)
	ParseToken(token string) (*MyClaims, error)
}

type myJwt struct {
	signingKey []byte
}

type MyClaims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

// New 返回一个新的Jwt interface
func New(signingKey []byte) Jwt {
	return Jwt(&myJwt{signingKey: signingKey})
}

func (m *myJwt) GetToken(claims *MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.signingKey)
}

func (m *myJwt) ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return m.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &MyClaims{
			Uid: int64(claims["uid"].(float64)),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: int64(claims["exp"].(float64)),
				Issuer:    claims["iss"].(string),
			},
		}, nil
	} else {
		return nil, errors.New("token is valid")
	}
}
