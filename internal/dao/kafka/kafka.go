package kafka

var GlobalProducer = map[string]*KafkaConn{}

type KafkaConn struct {
	//produer *sarama.Producer
}
