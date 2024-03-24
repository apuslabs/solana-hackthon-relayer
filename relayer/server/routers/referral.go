package routers

import (
	"apus-relayer/relayer/db"
	"apus-relayer/relayer/model"
	"apus-relayer/relayer/server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initReferral(r *gin.Engine) {
	//r.GET("/ranking", RankingHandler)
	//r.GET("/currentRanking", CurrentRankingHandler)
	r.GET("/referral/referrals", ReferralsHandler)
}

// 直接邀请的用户列表，根据邀请列表获取银河积分，根据银河积分算出提供给当前用户的bouns
// 第一个是用户自己的积分和获得到的总bouns
func ReferralsHandler(c *gin.Context) {
	pubkey := c.PostForm("publickey")
	result, err := db.Referrals(pubkey)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}

	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: result})
}

func RankingHandler(c *gin.Context) {
	result, err := db.Ranking()
	if err != nil {
		c.JSON(http.StatusOK, model.Response{Code: 500, Msg: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: result})
}

func CurrentRankingHandler(c *gin.Context) {
	userInfo := handler.GetUser(c)
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "", Data: db.CurrentRanking(userInfo.PubKey)})
}
