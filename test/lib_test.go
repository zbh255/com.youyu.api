package test

import (
	"com.youyu.api/app/rpc/model"
	"com.youyu.api/lib/auth"
	"com.youyu.api/lib/ecode"
	cl "com.youyu.api/lib/log"
	"com.youyu.api/lib/path"
	"com.youyu.api/lib/utils"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
	"github.com/wumansgy/goEncrypt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestPKGERRORS(t *testing.T) {
	_, err := ioutil.ReadFile("./jgp.lks")
	pkgErr := errors.WithStack(err)
	log.Error().Timestamp().Msgf("%+v", pkgErr)
	err = errors.Cause(pkgErr)
	t.Errorf("%+v", pkgErr)
	err = errors.New(err.Error())
	t.Log(err)
}

func TestModelError(t *testing.T) {
	err := errors.WithStack(model.UserNameAlreadyExists)
	switch errors.Cause(err) {
	case model.UserNameAlreadyExists:
		t.Logf("%+v", err)
	case model.UserDoesNotExist:
		t.Logf("%+v", err)
	}
}

func TestMUtilsPKGERRORS(t *testing.T) {
	err := E3(E2(E1()))
	t.Errorf("%+v\n", err)
	cerr := errors.Cause(err)
	t.Errorf("%+v", cerr)
}

func E1() error {
	return errors.New("my err1")
}

func E2(err error) error {
	return errors.Wrap(err, "my err2")
}

func E3(err error) error {
	return errors.Wrap(err, "my err3")
}

func TestEcode(t *testing.T) {
	file, err := os.Open(path.InfoFileDefaultPath + "/" + path.ErrMsgJsonFileName)
	if err != nil {
		t.Error(file)
	}
	codesMap, err := utils.ReadErrJsonToCodesMap(io.Reader(file))
	if err != nil {
		t.Error(err)
	}
	s := codesMap[int(ecode.NothingFound)]
	t.Log(s)
}

func TestCustomLogger(t *testing.T) {
	zLogger := &cl.ZLogger{
		Level:  zerolog.ErrorLevel,
		Logger: log.Output(os.Stderr),
	}
	logger := cl.Logger(zLogger)
	_, err := os.Open("/tutu")
	logger.Panic(err)
}

func TestUtilsTag(t *testing.T) {
	strings := utils.TagListToSplitStrings([]string{"1", "2", "3"})
	if strings != "" {
		t.Log(strings)
	}
	tags := utils.SplitStringsToTagList(strings)
	if tags != nil {
		t.Log(tags)
	}
}

func TestJwtParse(t *testing.T) {
	var hmacSampleSecret []byte
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}

func TestNewJwt(t *testing.T) {
	mySigningKey := []byte("AllYourBase")
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000000000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	token2, err := jwt.Parse(ss, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		t.Error(err)
	}
	if claims, ok := token2.Claims.(jwt.MapClaims); ok && token2.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}

func TestJWTERR(t *testing.T) {
	// 构造一个错误
	_, err := jwt.Parse("llll", func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hello"), nil
	})
	val := err.(*jwt.ValidationError)
	t.Log(val.Errors == jwt.ValidationErrorMalformed)
}

func TestCustomJwt(t *testing.T) {
	authJwt := auth.New([]byte("hello world"))
	token, err := authJwt.GetToken(&auth.MyClaims{
		Uid: 19,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000000000,
			Id:        "Unmarshal",
			Issuer:    "test",
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
	myClaims, err := authJwt.ParseToken(token)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(myClaims.Uid)
	}
}

func TestGoRedis(t *testing.T) {

}

func TestGoEncrypt(t *testing.T) {
	goEncrypt.GetRsaKey()
}

func TestUUid(t *testing.T) {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4(), nil)
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2 := uuid.NewV4()

	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err := uuid.FromString("2f396d98-329a-40e5-a58b-704a360868d3")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
