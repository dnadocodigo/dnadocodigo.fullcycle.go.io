package main

import (
	"encoding/json"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/infra/kafka"
	"github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/dto"
	entity "github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/entities"
	"github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/transformers"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}
	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})

	go kafka.Consume(kafkaMsgChan) // T2

	// revebe do canal do kafka, joga no input, processa joga no output e depois publica
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // T3

	go func(){
		for msg := range kafkaMsgChan{
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil{
				panic(err)
			}
			order := transformers.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()
	for res := range ordersOut{
		output := transformers.TransformOutput(res)
		outPutJson, err := json.MarshalIndent(output, "", " ")
		
	}
}
