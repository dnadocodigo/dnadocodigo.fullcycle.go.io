package kafka

ConfigMap *ckafka.ConfigMap
	Topics    []string

type Producer struct{
	ConfigMap *ckafka.ConfigMap
}