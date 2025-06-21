package v2id

import (
	"github.com/google/uuid"
)

func V2UUIDString() string {
	return "v2-" + uuid.NewString()
}
