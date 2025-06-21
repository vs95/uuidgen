package alpha_beta

import (
	googleuuid "github.com/google/uuid"
)

func AlphaUUIDString() string {
	return "alpha-" + googleuuid.NewString()
}
