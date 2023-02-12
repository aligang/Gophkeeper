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
	Id       string
	Value    string
	Owner    string
	IssuedAt time.Time
}

func New(owner string) *Token {
	tokenValue, _ := genTokenValue(tokenLength)
	return &Token{
		Id:       uuid.New().String(),
		Value:    tokenValue,
		Owner:    owner,
		IssuedAt: time.Now(),
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

//const LetterAlphabet = 26
//const DigitAlphabet = 10

//
//func base64Random(rnd []byte) []byte {
//	generator := []func(int) byte{
//		func(i int) byte { return 'A' + rnd[i]%LetterAlphabet },
//		func(i int) byte { return 'a' + rnd[i]%LetterAlphabet },
//		func(i int) byte { return '0' + rnd[i]%DigitAlphabet },
//	}
//	rand.Prime()
//	for i := range rnd {
//		rnd[i] = generator[rnd[i]%3](i)
//	}
//}
