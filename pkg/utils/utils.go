package utils

import (
	"strings"

	"github.com/google/uuid"
)

// GenerateUUID to generate random uuid.
func GenerateUUID() string {
	id := uuid.New()
	return strings.Replace(id.String(), "-", "", -1)
}
