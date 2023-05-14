package tgRepository

import (
	"bot_forge_back/internal/repository"
	"context"
	"github.com/go-redis/redis/v8"
)

type userRepo struct {
	rdb *redis.Client
}

func NewUserRepo(r *redis.Client) *userRepo {
	return &userRepo{rdb: r}
}

func (r userRepo) SetUserState(botId string, userId string, blockId string) error {
	err := r.rdb.Set(context.Background(), repository.BotUserStateKey(botId, userId), blockId, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r userRepo) GetUserState(botId string, userId string) (string, error) {
	blockId, err := r.rdb.Get(context.Background(), repository.BotUserStateKey(botId, userId)).Result()
	if err != nil {
		return "", err
	}
	return blockId, nil
}
