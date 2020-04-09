package tools

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

const (
	TokenSecretKey string = "token_secret_key"
	ExpireCount int = 2
)

var OmissionPath []string = []string{"/api/register", "/api/login", "/api/search/book"}

// 创建加密Token
func CreateJWTToken(accountId int, account string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(ExpireCount)).Unix(),
		"iat": time.Now().Unix(),
		"account": account,
		"accountId": accountId,
	})
	tokenString, _ := token.SignedString([]byte(TokenSecretKey))
	return tokenString
}

// 验证token的中间件
func JwtMiddleWare() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return TokenSecretKey, nil
		},
		ContextKey:          "",
		ErrorHandler:        jwtValidateError,
		CredentialsOptional: false,
		Extractor:           nil,
		EnableAuthOnOptions: false,
		SigningMethod:       jwt.SigningMethodHS256,
		Expiration:          false,
	})
}
// 处理验证失败问题
func jwtValidateError(ctx iris.Context, err error) {
	if err == nil {
		return
	}
	requestPath := ctx.Path()
	for _, path := range OmissionPath {
		if strings.Contains(requestPath, path){
			ctx.Next()
			return
		}
	}

	ctx.StopExecution()
	ctx.JSON(ApiResource("fail", nil, err.Error()))
}