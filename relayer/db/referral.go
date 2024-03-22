package db

import (
	"apus-relayer/relayer/utils"
	"context"
)

const CODE_CACHE = "code_cache"

var codeCache = make(map[string]string)
var referrals = make(map[string][]string)
var refSum = make(map[string]int64)
var referralRevers = make(map[string]string)

// 邀请码：pubkey映射
func GetCode(pubkey string) (string, error) {
	ctx := context.Background()
	code := ""
	for {
		code = utils.GenerateCode(8)
		// 判断重复
		result, err := rdb.HGet(ctx, CODE_CACHE, code).Result()
		if err != nil {
			return "", err
		}
		// code exist
		if result != "" {
			continue

		}
		_, err = rdb.HSet(ctx, CODE_CACHE, code, pubkey).Result()
		if err != nil {
			return "", err
		}
		return code, nil
	}
}

// 用户的邀请列表      pubkey---[pubkeys]
func AddReferral(referralCode string, pubkey string) {
	// 增加用户邀请人数
	// 记录用户邀请列表
	// 更新用户排名list
}

// 用户邀请排行        zset pubkey-sumscope
func rank() {
	// 获取当前邀请排名list
}

// 查看邀请列表

// 查看推荐人(邀请人) 邀请人被邀请人映射(结算)   被邀请人pubkey --- 邀请人pubkey
func GetMaster(pubkey string) string {
	return referralRevers[pubkey]
}
