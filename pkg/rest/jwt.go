package rest

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Operator ...
type Operator struct {
	jwt.StandardClaims
	Name string
}

// MySecret ...
var MySecret = []byte("#1024Helper")

// TokenExpireDuration expireDuration
const TokenExpireDuration = time.Hour * 2

// AuthCros ....
func AuthCros() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator, err := authorization(c.Request.Header.Get("authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, &ResultReply{
				Msg: err.Error(),
			})
			return
		}
		c.Set("operator", operator)
		c.Next()
	}
}

// authorization ...
func authorization(tokenStr string) (o *Operator, err error) {
	if tokenStr == "" {
		err = errors.New("authorization is null")
		return
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Operator{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return
	}
	o, ok := token.Claims.(*Operator)
	if ok && token.Valid {
		return
	}
	err = errors.New("authorization is invalid")
	return
}

// GenToken 生成JWT
func GenToken(name string) (string, error) {
	c := Operator{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "1024Helper",                               // 签发人
		},
		name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}
