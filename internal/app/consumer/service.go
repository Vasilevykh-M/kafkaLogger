package consumer

import (
	"context"
	"fmt"
)

type Receiver interface {
	Subscribe(topic string) error
}

type Service struct {
	receiver Receiver
}

func NewService(receiver Receiver) *Service {
	return &Service{
		receiver: receiver,
	}
}

func (s *Service) StartConsume(topic string) error {
	err := s.receiver.Subscribe(topic)

	if err != nil {
		fmt.Println("Subscribe error ", err)
		return err
	}
	return nil
}

func (s Service) ConsumerRun() error {
	err := s.StartConsume("URL")
	<-context.TODO().Done()
	return err
}
