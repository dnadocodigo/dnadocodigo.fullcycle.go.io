package kafka

import ckafka "github.com/confluentic/confluentic-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
}