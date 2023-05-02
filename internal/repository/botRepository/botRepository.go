package botRepository

import (
	"advocate-back/internal/repository"
	"context"
	"github.com/go-redis/redis/v8"
)

type repo struct {
	rdb *redis.Client
}

func NewRepo(r *redis.Client) *repo {
	return &repo{rdb: r}
}

func (r repo) AddBotToBotsList(botID string) error {
	err := r.rdb.SAdd(context.Background(), repository.AllBotsKey, botID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) DeleteBotFromList(botID string) error {
	err := r.rdb.SRem(context.Background(), repository.AllBotsKey, botID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GelAllBotsFromList() ([]string, error) {
	bots, err := r.rdb.SMembers(context.Background(), repository.AllBotsKey).Result()
	if err != nil {
		return nil, err
	}
	return bots, nil
}

func (r repo) CreateBotConfig(botID string, config []byte) error {
	err := r.rdb.Set(context.Background(), repository.BotConfigKey(botID), config, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GetBotConfigByID(botID string) (string, error) {
	conf, err := r.rdb.Get(context.Background(), repository.BotConfigKey(botID)).Result()
	if err != nil {
		return "", err
	}
	return conf, nil
}

func (r repo) CreateBotToken(botID string, token string) error {
	err := r.rdb.Set(context.Background(), repository.BotTokenKey(botID), token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GetBotTokenByID(botID string) (string, error) {
	conf, err := r.rdb.Get(context.Background(), repository.BotTokenKey(botID)).Result()
	if err != nil {
		return "", err
	}
	return conf, nil
}
