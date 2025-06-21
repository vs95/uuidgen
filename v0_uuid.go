package v0id

import (
	googleuuid "github.com/google/uuid"
)

func V1UUIDString() string {
	return "v0-" + googleuuid.NewString()
}
