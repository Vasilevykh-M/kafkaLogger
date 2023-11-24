package producer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"homework-6/internal/infrastructure/kafka"
	"time"
)

type Request struct {
	AnswerID    int
	TimeRequest time.Time
	Header      string
	Body        string
}

type KafkaSender struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaSender(producer *kafka.Producer, topic string) *KafkaSender {
	return &KafkaSender{
		producer,
		topic,
	}
}

func (s *KafkaSender) sendMessage(message Request) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return err
	}

	_, _, err = s.producer.SendSyncMessage(kafkaMsg)

	if err != nil {
		fmt.Println("Send message connector error", err)
		return err
	}

	return nil
}

func (s *KafkaSender) buildMessage(message Request) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Send message marshal error", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprint(message.AnswerID)),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("test-header"),
				Value: []byte("test-value"),
			},
		},
	}, nil
}
