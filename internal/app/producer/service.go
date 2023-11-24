package producer

import (
	"fmt"
	"time"
)

type Sender interface {
	sendMessage(message Request) error
}

type Service struct {
	sender Sender
}

func NewService(sender Sender) *Service {
	return &Service{
		sender: sender,
	}
}

func (s *Service) Send(header, body string) {

	err := s.sender.sendMessage(
		Request{
			TimeRequest: time.Now(),
			Header:      header,
			Body:        body,
		},
	)

	if err != nil {
		fmt.Println("Send sync message error: ", err)
	}

	return
}
