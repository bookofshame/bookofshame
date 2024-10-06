package sms

import "fmt"

type Payload struct {
	Number  string
	Message string
}

type Sms interface {
	// Send The body should be a valid HTML string
	Send(payload Payload) error
}

func (s *Payload) Validate() error {
	if s.Number == "" {
		return fmt.Errorf("number cannot be empty")
	}

	if s.Message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	return nil
}
