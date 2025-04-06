package app

import (
	"MasterBroker/internal/broker"
	"MasterBroker/internal/dispatcher"
	"fmt"
)

func Run() error {
	_dispatcher := dispatcher.NewDispatcher()
	_broker := broker.NewRedisBroker("localhost:6379", _dispatcher)

	_broker.Subscribe("events")
	fmt.Println("Broker is running")
	select {}
}
