package middleware

import (
	"LdapAdmin/common/constant"
	"LdapAdmin/common/model"
	"LdapAdmin/config"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type CustomerClaims struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Active   string `json:"active"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func JwtHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		claims, err := ParseToken(token)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				model.ResponseErr{
					Code:   constant.TokenError,
					ErrMsg: err.Error(),
				},
			)
			context.Abort()
			return
		}
		context.Set("role", claims.Role)
		context.Set("account", claims.Account)
		context.Set("active", claims.Active)
		context.Next()
	}
}

func GenerateToken(account string, password string, active string, role string) (string, error) {
	claims := CustomerClaims{
		Account:  account,
		Password: password,
		Active:   active,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Conf.System.TokenExpired) * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Conf.System.TokenSecret)
}

func ParseToken(tokenString string) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.System.TokenSecret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle the token")
			}
		}
	}
	if claims, ok := token.Claims.(*CustomerClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle the token")
}
