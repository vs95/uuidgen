package gammapkg

import (
	googleuuid "github.com/google/uuid"
)

func UUIDString() string {
	return "gamma-" + googleuuid.NewString()
}
