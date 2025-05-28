package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

var Validate = validator.New()