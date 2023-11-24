package consumer

import (
	"fmt"
	"github.com/IBM/sarama"
	"homework-6/internal/infrastructure/kafka"
)

type HandleFunc func(message *sarama.ConsumerMessage)

type KafkaReceiver struct {
	consumer *kafka.Consumer
}

func NewReceiver(consumer *kafka.Consumer) *KafkaReceiver {
	return &KafkaReceiver{
		consumer: consumer,
	}
}

func (r *KafkaReceiver) Subscribe(topic string) error {

	partitionList, err := r.consumer.SingleConsumer.Partitions(topic)

	if err != nil {
		return err
	}
	initialOffset := sarama.OffsetNewest

	for _, partition := range partitionList {
		pc, err := r.consumer.SingleConsumer.ConsumePartition(topic, partition, initialOffset)

		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer, partition int32) {
			for message := range pc.Messages() {
				fmt.Println("From Kafka Received Key: ", string(message.Key), " Value: ", string(message.Value))
			}
		}(pc, partition)
	}

	return nil
}
