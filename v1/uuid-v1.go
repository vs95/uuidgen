package v1id

import (
	googleuuid "github.com/google/uuid"
)

func V1UUIDString() string {
	return "v1-" + googleuuid.NewString()
}
