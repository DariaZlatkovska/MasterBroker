package storage

import (
	"MasterBroker/config"
	"context"
)

func SaveMessageToRedis(id, content string) error {
	err := config.Rdb.Set(context.Background(), "message"+id, content, 0).Err()
	return err
}
func GetMessageFromRedis() ([]string, error) {
	keys, err := config.Rdb.Keys(context.Background(), "message*").Result()
	if err != nil {
		return nil, err
	}

	var message []string
	for _, key := range keys {
		msg, err := config.Rdb.Get(context.Background(), key).Result()
		if err != nil {
			message = append(message, msg)
		}
	}
	return message, nil
}
