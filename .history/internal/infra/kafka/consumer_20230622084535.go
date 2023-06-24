package kafka

import ckafka "github.com/confluentic/confuentic-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
}