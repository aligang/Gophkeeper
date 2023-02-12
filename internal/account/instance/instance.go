package instance

import "time"

type Account struct {
	Id                string
	CreatedAt         time.Time
	Login             string
	Password          string
	EncryptionEnabled bool
	EncryptionKey     string
}
