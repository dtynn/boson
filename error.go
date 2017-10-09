package boson

import (
	"fmt"
)

type RequestError struct {
	Code    int
	Message string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("boson.RequestError: %d-%q", r.Code, r.Message)
}

type MessageError struct {
	Err string `json:"error"`
}

func (m *MessageError) Error() string {
	return fmt.Sprintf("boson.MessageError: %q", m.Err)
}
