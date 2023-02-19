package encryption

import "crypto/rand"

const KeySize int = 32

func NewKey() (string, error) {
	key := make([]byte, KeySize)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return string(base64Random(key)), nil
}

const LetterAlphabet = 26
const DigitAlphabet = 10

func base64Random(rnd []byte) []byte {
	generator := []func(int) byte{
		func(i int) byte { return 'A' + rnd[i]%LetterAlphabet },
		func(i int) byte { return 'a' + rnd[i]%LetterAlphabet },
		func(i int) byte { return '0' + rnd[i]%DigitAlphabet },
	}

	for i := range rnd {
		rnd[i] = generator[rnd[i]%3](i)
	}

	return rnd
}
