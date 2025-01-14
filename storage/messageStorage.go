package storage

import (
	"MasterBroker/config"
	"MasterBroker/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var messageCollection *mongo.Collection

func InitMessageStorage() {
	messageCollection = config.MongoDB.Database("broker").Collection("messages")
}

func SaveMessages(message models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := messageCollection.InsertOne(ctx, message)
	if err != nil {
		log.Printf("error when save message: %v", err)
		return err
	}
	return nil
}

func GetMessages() ([]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := messageCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("error when get message: %v ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var message []models.Message
	if err := cursor.All(ctx, &message); err != nil {
		return nil, err
	}
	return message, nil
}
