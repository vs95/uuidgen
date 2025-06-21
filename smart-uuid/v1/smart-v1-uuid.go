package v1smartid

import (
	"github.com/google/uuid"
)

func V1UUIDSmart() string {
	return "v1-smart-" + uuid.NewString()
}
