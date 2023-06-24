package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics: topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) error{
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if(err != nil){
		panic(err)
	}
	err = consumer.Subscribe
}