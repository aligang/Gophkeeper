package encryption

import (
	secretInstance "github.com/aligang/Gophkeeper/internal/secret/instance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextSecretEncryption(t *testing.T) {
	encryptedSecret, err := EncryptTextSecret(referenceTextSecret, referenceEncryptionKey)
	assert.Equal(t, nil, err)

	assert.Equal(t, referenceTextSecret.Id, encryptedSecret.Id)
	assert.Equal(t, referenceTextSecret.AccountId, encryptedSecret.AccountId)
	assert.Equal(t, referenceTextSecret.CreatedAt, encryptedSecret.CreatedAt)
	assert.Equal(t, referenceTextSecret.ModifiedAt, encryptedSecret.ModifiedAt)

	decryptedText, err := decrypt(encryptedSecret.Text, referenceEncryptionKey)
	assert.Equal(t, referenceTextData, decryptedText)
}

func TestTextSecretDecryption(t *testing.T) {

	encryptedTextSecret := &secretInstance.TextSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         referenceSecretId,
			AccountId:  referenceAccountId,
			CreatedAt:  referenceCreationTime,
			ModifiedAt: referenceModificationTime,
		},
		Text: "HiZmu0qIyS77IDHIzCP5WSLvncEGbSO+U1Z4elMNOC7m4U3KPg==",
	}

	decryptedSecret, err := DecryptTextSecret(encryptedTextSecret, referenceEncryptionKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, referenceTextSecret, decryptedSecret)

}
