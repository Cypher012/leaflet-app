package utils

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid id")

func ValidateUUID(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidID
	}
	return nil
}
