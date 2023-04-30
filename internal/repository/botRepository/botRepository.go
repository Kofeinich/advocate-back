package botRepository

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func botConfigKey(id string) string {
	return "botConfig_" + id
}

func botTokenKey(id string) string {
	return "botToken_" + id
}

const allBotsKey = "allBots"

type repo struct {
	rdb *redis.Client
}

func NewRepo(r *redis.Client) *repo {
	return &repo{rdb: r}
}

func (r repo) AddBotToBotsList(botID string) error {
	err := r.rdb.SAdd(context.Background(), allBotsKey, botID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) DeleteBotFromList(botID string) error {
	err := r.rdb.SRem(context.Background(), allBotsKey, botID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GelAllBotsFromList() error {
	err := r.rdb.SMembers(context.Background(), allBotsKey).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) CreateBotConfig(botID string, config string) error {
	err := r.rdb.Set(context.Background(), botConfigKey(botID), config, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GetBotConfigByID(botID string) (string, error) {
	conf, err := r.rdb.Get(context.Background(), botConfigKey(botID)).Result()
	if err != nil {
		return "", err
	}
	return conf, nil
}

func (r repo) CreateBotToken(botID string, token string) error {
	err := r.rdb.Set(context.Background(), botTokenKey(botID), token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r repo) GetBotTokenByID(botID string) (string, error) {
	conf, err := r.rdb.Get(context.Background(), botTokenKey(botID)).Result()
	if err != nil {
		return "", err
	}
	return conf, nil
}
