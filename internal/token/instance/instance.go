package instance

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const tokenLength = 20

type Token struct {
	Id         string
	TokenValue string
	Owner      string
	IssuedAt   time.Time
}

func New(owner string) *Token {
	tokenValue, _ := genTokenValue(tokenLength)
	return &Token{
		Id:         uuid.New().String(),
		TokenValue: tokenValue,
		Owner:      owner,
		IssuedAt:   time.Now(),
	}
}

func genTokenValue(n int) (string, error) {
	rnd := make([]byte, n)
	nrnd, err := rand.Read(rnd)
	if err != nil {
		return "", err
	} else if nrnd != n {
		return "", fmt.Errorf(`nrnd %d != n %d`, nrnd, n)
	}

	return base64.StdEncoding.EncodeToString(rnd), nil
}
