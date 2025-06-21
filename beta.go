package alpha_beta

import (
	googleuuid "github.com/google/uuid"
)

func BetaUUIDString() string {
	return "beta-" + googleuuid.NewString()
}
