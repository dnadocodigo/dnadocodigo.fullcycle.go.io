package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Producer struct{
	ConfigMap *ckafka.ConfigMap
}
func NewKafkaProducer(configMap *ckafka.ConfigMap)*Producer{
	return &Producer{
		ConfigMap: *ckafka.ConfigMap,
	}
}
func (p *Producer) Publish(msg interface{}, key )