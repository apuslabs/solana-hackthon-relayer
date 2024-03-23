package routers

import (
	"apus-relayer/relayer/db"
	"apus-relayer/relayer/model"
	"apus-relayer/relayer/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initReferral(r *gin.Engine) {
	r.GET("/ranking", Ranking)
	r.GET("/currentRanking", CurrentRanking)
	r.GET("/referrals", Referrals)
}

func Ranking(c *gin.Context) {
	result, err := db.Ranking()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: result})
}

func CurrentRanking(c *gin.Context) {
	userInfo := handler.GetUser(c)
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: db.CurrentRanking(userInfo.PubKey)})
}

func Referrals(c *gin.Context) {
	userInfo := handler.GetUser(c)
	result, err := db.Referrals(userInfo.PubKey)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: result})
}
