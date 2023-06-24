package main

import (
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	entity "github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/entities"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)
	co
}
