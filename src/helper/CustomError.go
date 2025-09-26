package helper

import (
	"fmt"
)

type CustomErrors struct {
	Message string
	Field  string
}

func (ce *CustomErrors) CreateUserError() error {
	return fmt.Errorf("❌ User with email %s %s\n", ce.Field, ce.Message)
}