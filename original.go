package idoriginal

import (
	googleuuid "github.com/google/uuid"
)

func V1UUIDString() string {
	return "idoriginal-" + googleuuid.NewString()
}
