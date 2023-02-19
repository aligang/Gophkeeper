package encryption

import (
	secretInstance "github.com/aligang/Gophkeeper/pkg/common/secret/instance"
)

func EncryptTextSecret(plainSecret *secretInstance.TextSecret, encryptionKey string) (*secretInstance.TextSecret, error) {
	var err error
	encryptedSecret := &secretInstance.TextSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         plainSecret.Id,
			AccountId:  plainSecret.AccountId,
			CreatedAt:  plainSecret.CreatedAt,
			ModifiedAt: plainSecret.ModifiedAt,
		},
	}
	encryptedSecret.Text, err = encrypt(plainSecret.Text, encryptionKey)
	if err != nil {
		return nil, err
	}
	return encryptedSecret, nil
}

func DecryptTextSecret(encryptedSecret *secretInstance.TextSecret, encryptionKey string) (*secretInstance.TextSecret, error) {
	var err error
	decryptedSecret := &secretInstance.TextSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         encryptedSecret.Id,
			AccountId:  encryptedSecret.AccountId,
			CreatedAt:  encryptedSecret.CreatedAt,
			ModifiedAt: encryptedSecret.ModifiedAt,
		},
	}
	decryptedSecret.Text, err = decrypt(encryptedSecret.Text, encryptionKey)
	if err != nil {
		return nil, err
	}
	return decryptedSecret, nil
}
