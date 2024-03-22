package routers

import (
	"apus-relayer/relayer/db"
	"apus-relayer/relayer/model"
	"apus-relayer/relayer/server/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// 用户注册 & 登录
func LoginHandler(c *gin.Context) {

	key := c.PostForm("publickey")
	code := c.PostForm("code")
	if key == "" {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "publickey must be full", Data: ""})
		return
	}

	user, err := db.GetUser(key)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "", Data: err})
		return
	}

	// 用户未注册
	if user.Code == "" {
		user.Code, err = db.GetCode(key)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{Code: 401, Msg: err.Error(), Data: ""})
			return
		}
		user.PubKey = key
		user.Ctime = time.Now().Unix()
		if err := db.AddUser(user); err != nil {
			c.JSON(http.StatusOK, model.Response{Code: 401, Msg: err.Error(), Data: ""})
			return
		}
		if code != "" {
			// add referral
			fmt.Printf("referral code: %s, new user code: %s", code, user.Code)
			db.AddReferral(code, user.PubKey)
		}
	}
	claims := handler.UserClaims{
		User:             user,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString(handler.SIGNKEY)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: token})
}
