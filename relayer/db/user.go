package db

import (
	"apus-relayer/relayer/model"
	"context"
	"errors"
)

const USER_KEY = "user"

func GetUser(pubKey string) (model.UserInfo, error) {
	ctx := context.Background()
	user := &model.UserInfo{}
	err := rdb.HGet(ctx, USER_KEY, pubKey).Scan(user)
	if err != nil {
		return model.UserInfo{}, err
	}
	return *user, nil
}

func AddUser(user model.UserInfo) error {
	ctx := context.Background()
	oUser, err := GetUser(user.PubKey)
	if oUser.PubKey != "" {
		return errors.New("user exist")
	}
	_, err = rdb.HSet(ctx, USER_KEY, user.PubKey, &user).Result()
	if err != nil {
		return err
	}
	return nil
}
