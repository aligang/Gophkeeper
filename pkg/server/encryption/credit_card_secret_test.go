package encryption

import (
	secretInstance "github.com/aligang/Gophkeeper/pkg/secret/instance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditCardSecretEncryption(t *testing.T) {
	copyContainer := make([]*secretInstance.CreditCardSecret, 1)
	copy(copyContainer, []*secretInstance.CreditCardSecret{referenceCreditCardSecret})
	plainSecret := copyContainer[0]

	encryptedSecret, err := EncryptCreditCardSecret(plainSecret, referenceEncryptionKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, referenceCreditCardSecret.Id, encryptedSecret.Id)
	assert.Equal(t, referenceCreditCardSecret.AccountId, encryptedSecret.AccountId)
	assert.Equal(t, referenceCreditCardSecret.CreatedAt, encryptedSecret.CreatedAt)
	assert.Equal(t, referenceCreditCardSecret.ModifiedAt, encryptedSecret.ModifiedAt)

	decryptedCardNumber, err := decrypt(encryptedSecret.Number, referenceEncryptionKey)
	assert.Equal(t, referenceCardNumber, decryptedCardNumber)
	decryptedCardHolder, err := decrypt(encryptedSecret.CardHolder, referenceEncryptionKey)
	assert.Equal(t, referenceCardHolder, decryptedCardHolder)
	decryptedValidTill, err := decrypt(encryptedSecret.ValidTill, referenceEncryptionKey)
	assert.Equal(t, referenceValidTill, decryptedValidTill)
	decryptedCvc, err := decrypt(encryptedSecret.Cvc, referenceEncryptionKey)
	assert.Equal(t, referenceCvc, decryptedCvc)
}

func TestCreditCardSecretDecryption(t *testing.T) {
	encryptedCreditCardSecret := &secretInstance.CreditCardSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         referenceSecretId,
			AccountId:  referenceAccountId,
			CreatedAt:  referenceCreationTime,
			ModifiedAt: referenceModificationTime,
		},
		Number:     "dpGMcbO6apxrCRumHG1VcBtsUkbLBmVNuHDbS9DG+YjmdzFbeeWHZD49fOQ=",
		CardHolder: "JDDUrpS+g7QtpGUrzJbK1ZT8PNj7RqAUR+bfUNcUPr+UuWKB",
		ValidTill:  "c9SFXhG7J7TaKn+OUszGAO7bxzbpdEElYWBeWR9L8CFf",
		Cvc:        "M4bYwfqjD8+9lAKUl/wSm/wpQ+J/7bvaUhE6HhgOEg==",
	}

	decryptedSecret, err := DecryptCreditCardSecret(encryptedCreditCardSecret, referenceEncryptionKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, referenceCreditCardSecret, decryptedSecret)

}
