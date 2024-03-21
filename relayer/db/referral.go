package db

import "apus-relayer/relayer/utils"

var codeCache = make(map[string]string)
var referrals = make(map[string][]string)
var refSum = make(map[string]int64)
var referralRevers = make(map[string]string)

// 邀请码：pubkey映射
func GetCode(pubkey string) string {
	code := ""
	for {
		code = utils.GenerateCode(8)
		// 判断重复
		if codeExist(code) {
			continue
		}
		codeCache[code] = pubkey
		return code
	}
}

// 用户的邀请列表      pubkey---[pubkeys]
func AddReferral(code string, pubkey string) {
	master := codeCache[code]
	followers := referrals[master]
	followers = append(followers, pubkey)
	referrals[master] = followers
	refSum[pubkey] = refSum[pubkey] + 1
}

// 用户邀请排行        zset pubkey-sumscope

// 邀请人被邀请人映射(结算)   被邀请人pubkey --- 邀请人pubkey
func GetMaster(pubkey string) string {
	return referralRevers[pubkey]
}

func codeExist(code string) bool {
	if _, ok := codeCache[code]; ok {
		return true
	}
	return false
}
