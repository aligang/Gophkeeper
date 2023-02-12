package handler

import (
	"github.com/aligang/Gophkeeper/internal/secret"
)

func CheckOwnership(a, b string) error {
	if a == b {
		return nil
	}
	return secret.ErrAccessProhibited
}
