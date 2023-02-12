package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func encrypt(plainText string, key string) (string, error) {
	data := []byte(plainText)
	keySlice := []byte(key)

	c, err := aes.NewCipher(keySlice)
	if err != nil {
		return "", ErrCipherCreation
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err

	}
	gcmEnc := gcm.Seal(nonce, nonce, data, nil)
	res := base64.StdEncoding.EncodeToString(gcmEnc)
	return res, nil
}

func decrypt(data string, key string) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", ErrDataDecoding
	}
	keySlice := []byte(key)

	c, err := aes.NewCipher(keySlice)
	if err != nil {
		return "", ErrCipherCreation
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return "", err
	}

	nonce, encryptedData := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", ErrFailedDecryption
	}
	return string(plaintext), nil
}
