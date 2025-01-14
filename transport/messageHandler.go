package transport

import (
	"MasterBroker/models"
	"MasterBroker/storage"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message models.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "this format is not correct", http.StatusBadRequest)
		return
	}

	message.ID = uuid.New().String()
	message.Timestamp = time.Now()

	err = storage.SaveMessageToRedis(message.ID, message.Content)
	if err != nil {
		http.Error(w, "error when save message to redis", http.StatusInternalServerError)
		return
	}
	err = storage.SaveMessageToMongo(message)
	if err != nil {
		http.Error(w, "error when save message to mongo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := storage.GetMessageFromMongo()
	if err != nil {
		http.Error(w, "error when get message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}
