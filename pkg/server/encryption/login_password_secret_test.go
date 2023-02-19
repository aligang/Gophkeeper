package encryption

import (
	secretInstance "github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginPasswordSecretEncryption(t *testing.T) {
	copyContainer := make([]*secretInstance.LoginPasswordSecret, 1)
	copy(copyContainer, []*secretInstance.LoginPasswordSecret{referenceLoginPasswordSecret})
	plainSecret := copyContainer[0]

	encryptedSecret, err := EncryptLoginPasswordSecret(plainSecret, referenceEncryptionKey)
	assert.Equal(t, nil, err)

	assert.Equal(t, referenceLoginPasswordSecret.Id, encryptedSecret.Id)
	assert.Equal(t, referenceLoginPasswordSecret.AccountId, encryptedSecret.AccountId)
	assert.Equal(t, referenceLoginPasswordSecret.CreatedAt, encryptedSecret.CreatedAt)
	assert.Equal(t, referenceLoginPasswordSecret.ModifiedAt, encryptedSecret.ModifiedAt)

	decryptedLogin, err := decrypt(encryptedSecret.Login, referenceEncryptionKey)
	assert.Equal(t, referenceLogin, decryptedLogin)
	decryptedPassword, err := decrypt(encryptedSecret.Password, referenceEncryptionKey)
	assert.Equal(t, referencePassword, decryptedPassword)
}

func TestLoginPasswordSecretDecryption(t *testing.T) {
	encryptedLoginPasswordSecret := &secretInstance.LoginPasswordSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         referenceSecretId,
			AccountId:  referenceAccountId,
			CreatedAt:  referenceCreationTime,
			ModifiedAt: referenceModificationTime,
		},
		Login:    "3ocbyn+nRAmQewq0gfihD7iCS+29yVlM6bSovcUkpdVzNdcV7YM=",
		Password: "chCvOE+JcDau/q2T3BVapy8ZL/A05wHe45wmTamJuqVWJynGEg/BD3w=",
	}

	decryptedSecret, err := DecryptLoginPasswordSecret(encryptedLoginPasswordSecret, referenceEncryptionKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, referenceLoginPasswordSecret, decryptedSecret)

}
