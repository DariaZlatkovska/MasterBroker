package broker

import (
	"MasterBroker/internal/dispatcher"
	"MasterBroker/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisBroker struct {
	client     *redis.Client
	ctx        context.Context
	dispatcher *dispatcher.Dispatcher
}

func NewRedisBroker(addr string, dispatcher *dispatcher.Dispatcher) *RedisBroker {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisBroker{
		client:     rdb,
		ctx:        context.Background(),
		dispatcher: dispatcher,
	}
}

func (rb *RedisBroker) Subscribe(channel string) {
	sub := rb.client.Subscribe(rb.ctx, channel)
	ch := sub.Channel()

	go func() {
		for msg := range ch {
			fmt.Println("Received message:", msg.Payload)

			var base model.BaseMessage
			if err := json.Unmarshal([]byte(msg.Payload), &base); err != nil {
				fmt.Println("error unmarshalling message:", err)
				continue
			}

			if err := rb.dispatcher.Dispatch(base); err != nil {
				fmt.Println("error dispatching message:", err)
			}
		}
	}()
}
