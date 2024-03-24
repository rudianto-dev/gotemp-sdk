package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID to generate random uuid.
func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
