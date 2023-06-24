package kafka

import ckafka "github.com/confluentic/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
}