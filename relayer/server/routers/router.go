package routers

import (
	"apus-relayer/relayer/db"
	"apus-relayer/relayer/model"
	"apus-relayer/relayer/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	loadRouters(r)
	return r
}

func loadRouters(r *gin.Engine) {
	r.POST("/login", LoginHandler)

	r.POST("/getNode", getNodeHandler)
	r.POST("/submitMetrics", submitMetricsHandler)
	r.POST("/submitMetrics", submitMetricsHandler)
}

func getNodeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}

func submitMetricsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: "hello"})
}

func LoginHandler(c *gin.Context) {

	key := c.PostForm("publickey")
	user := db.GetUser(key)
	if user.Code == "" {
		user.Code = db.GetCode(key)
		user.PubKey = key
		user.Ctime = time.Now().Unix()
		if err := db.AddUser(user); err != nil {
			c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "", Data: err})
			return
		}
	}
	claims := handler.UserClaims{
		User:             user,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString(handler.SIGNKEY)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 401, Msg: "", Data: err})
		return
	}
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: token})
}
