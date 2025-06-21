package deltapkg

import (
	"github.com/google/uuid"
)

func UUIDString() string {
	return "delta-" + uuid.NewString()
}
