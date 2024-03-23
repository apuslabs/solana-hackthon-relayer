package handler

import (
	"apus-relayer/relayer/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

var SIGNKEY = []byte("apus-relayer")

type UserClaims struct {
	User model.UserInfo
	jwt.RegisteredClaims
}

func JWTAuthHandler(c *gin.Context) {
	if c.Request.URL.Path == "/login" {
		return
	}

	tokenString := c.Request.Header.Get("x-token")
	if tokenString == "" {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "permission denied, pls login", Data: ""})
		c.Abort()
		return
	}

	parser := jwt.NewParser()
	token, err := parser.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SIGNKEY, nil
	})
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: err.Error(), Data: ""})
		c.Abort()
		return
	}
	if !token.Valid {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "token invalid", Data: ""})
		c.Abort()
		return
	}

	if userClaims, ok := token.Claims.(*UserClaims); ok {
		c.Set("claims", userClaims)
	} else {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "claim not found", Data: ""})
		c.Abort()
		return
	}
}

func GetUser(c *gin.Context) model.UserInfo {
	claims, _ := c.Get("claims")
	result, _ := claims.(UserClaims)
	return result.User
}
