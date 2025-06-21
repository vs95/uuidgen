package thetapkg

import (
	"github.com/google/uuid"
)

func UUIDSmart() string {
	return "theta-" + uuid.NewString()
}
