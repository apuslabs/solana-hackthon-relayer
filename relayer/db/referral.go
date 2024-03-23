package db

import (
	"apus-relayer/relayer/model"
	"apus-relayer/relayer/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// 当前生成的CODE记录，用于去重 和 CODE-PUBKEY映射查找  code: pubkey
const CODE_CACHE = "code_cache"

// 被邀请人和邀请人关系映射集合 pubkey:pubkey
const REFERRAL_RELATION = "referral_relation"

// 当前用户的邀请列表，  pubkey: pubkeys
const REFERRAL_LIST = "referral_list"

// 排名列表   pubkey:scope
const REFERRAL_RANK = "referral_rank"

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

func AddReferral(referralCode string, user model.UserInfo) {
	ctx := context.Background()
	// 验证code合法性
	referralPubKey, err := rdb.HGet(ctx, CODE_CACHE, referralCode).Result()
	if err != nil || referralPubKey == "" {
		fmt.Printf("AddReferral:  获取%s对应公钥失败, 失败原因: %s\n", referralCode, err.Error())
		return
	}
	// 记录邀请关系 pubkey:pubkey
	_, err = rdb.HSet(ctx, REFERRAL_RELATION, user.PubKey, referralPubKey).Result()
	if err != nil {
		fmt.Printf("AddReferral: %s:%s 邀请关系记录失败, 失败原因: %s\n", user.PubKey, referralPubKey, err.Error())
	}
	// 记录用户邀请列表 pubkey: pubkeys
	key := fmt.Sprintf("%s_%s", REFERRAL_LIST, referralPubKey)
	_, err = rdb.RPush(ctx, key, user.PubKey).Result()
	if err != nil {
		fmt.Printf("AddReferral: %s:%s 邀请列表记录失败, 失败原因: %s\n", referralPubKey, user.PubKey, err.Error())
	}
	// 更新用户Rank排名list pubkey:scope
	count, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		fmt.Printf("AddReferral: %s 获取用户邀请count失败, 失败原因: %s\n", key, err.Error())
		return
	}
	_, err = rdb.ZAdd(ctx, REFERRAL_RANK, &redis.Z{Member: referralPubKey, Score: float64(count)}).Result()
	if err != nil {
		fmt.Printf("AddReferral: 更新CODE%s排名数据失败, 失败原因: %s\n", referralPubKey, err.Error())
		return
	}

}

// 用户邀请排行        zset pubkey-sumscope
func Ranking() ([]model.Ranking, error) {
	// 获取当前邀请排名list
	ctx := context.Background()
	op := redis.ZRangeBy{
		Min:    "0",
		Max:    "-1",
		Offset: 0,
		Count:  500,
	}

	ranks, err := rdb.ZRevRangeByScoreWithScores(ctx, REFERRAL_RANK, &op).Result()
	if err != nil {
		fmt.Printf("Ranking: 获取邀请排行失败, 失败原因: %s\n", err.Error())
		return []model.Ranking{}, err
	}
	result := make([]model.Ranking, 0)
	for _, val := range ranks {
		result = append(result, model.Ranking{Scope: int64(val.Score), Pubkey: fmt.Sprintf("%s", val.Member)})
	}
	return result, nil
}

// 用户当前邀请名次
func CurrentRanking(pubkey string) int64 {
	ctx := context.Background()
	index, err := rdb.ZRevRank(ctx, REFERRAL_RANK, pubkey).Result()
	if err != nil {
		fmt.Printf("CurrentRanking: 获取%s邀请排行失败, 失败原因: %s\n", pubkey, err.Error())
		return 0
	}
	return index
}

// 查看邀请列表
func Referrals(pubkey string) ([]string, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s_%s", REFERRAL_LIST, pubkey)
	result, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

// 查看推荐人(邀请人) 邀请人被邀请人映射(结算)   被邀请人pubkey --- 邀请人pubkey
func GetMaster(pubkey string) (string, error) {
	ctx := context.Background()
	referralPubkey, err := rdb.HGet(ctx, REFERRAL_RELATION, pubkey).Result()
	if err != nil {
		fmt.Printf("GetMaster: 获取%s邀请人信息失败, 失败原因: %s\n", pubkey, err.Error())
		return "", err
	}
	return referralPubkey, nil
}
