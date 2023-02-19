package handler

import (
	"github.com/aligang/Gophkeeper/pkg/common/secret"
)

func CheckOwnership(a, b string) error {
	if a == b {
		return nil
	}
	return secret.ErrAccessProhibited
}
