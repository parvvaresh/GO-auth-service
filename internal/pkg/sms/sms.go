package sms

import "log"

type Sender interface {
	Send(phone string, message string) error
}

type MockSender struct{}

func NewMockSender() *MockSender {
	return &MockSender{}
}

func (s *MockSender) Send(phone string, message string) error {
	log.Printf("Sending SMS to %s: %s", phone, message)
	return nil
}
