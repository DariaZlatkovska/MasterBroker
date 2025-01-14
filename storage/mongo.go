package storage

import (
	"MasterBroker/config"
	"MasterBroker/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveMessageToMongo(message models.Message) error {
	collection := config.MongoDB.Database("broker").Collection("messages")
	_, err := collection.InsertOne(context.Background(), bson.M{
		"id":        message.ID,
		"content":   message.Content,
		"timestamp": message.Timestamp,
	})
	return err
}

func GetMessageFromMongo() ([]models.Message, error) {
	collection := config.MongoDB.Database("broker").Collection("messages")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var messages []models.Message
	err = cursor.All(context.Background(), &messages)
	return messages, nil
}
