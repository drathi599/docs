package kafka

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	pb "github.com/cnnrznn/docs/editor"
)

// Need a better way to init a producer once only and then we can directly
// priduce
func produce(message *pb.Op) (ok bool) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()
	topic := "myTopic"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(word),
	}, nil)
	p.Flush(15 * 1000)
	return true
}
