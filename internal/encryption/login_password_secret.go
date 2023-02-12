package encryption

import secretInstance "github.com/aligang/Gophkeeper/internal/secret/instance"

func EncryptLoginPasswordSecret(plainSecret *secretInstance.LoginPasswordSecret,
	encryptionKey string) (*secretInstance.LoginPasswordSecret, error) {
	var err error
	encryptedSecret := &secretInstance.LoginPasswordSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         plainSecret.Id,
			AccountId:  plainSecret.AccountId,
			CreatedAt:  plainSecret.CreatedAt,
			ModifiedAt: plainSecret.ModifiedAt,
		},
	}

	encryptedSecret.Login, err = encrypt(plainSecret.Login, encryptionKey)
	if err != nil {
		return nil, err
	}
	encryptedSecret.Password, err = encrypt(plainSecret.Password, encryptionKey)
	if err != nil {
		return nil, err
	}
	return encryptedSecret, nil
}

func DecryptLoginPasswordSecret(encryptedSecret *secretInstance.LoginPasswordSecret,
	encryptionKey string) (*secretInstance.LoginPasswordSecret, error) {
	var err error
	plainSecret := &secretInstance.LoginPasswordSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         encryptedSecret.Id,
			AccountId:  encryptedSecret.AccountId,
			CreatedAt:  encryptedSecret.CreatedAt,
			ModifiedAt: encryptedSecret.ModifiedAt,
		},
	}

	plainSecret.Login, err = decrypt(encryptedSecret.Login, encryptionKey)
	if err != nil {
		return nil, err
	}
	plainSecret.Password, err = decrypt(encryptedSecret.Password, encryptionKey)
	if err != nil {
		return nil, err
	}
	return plainSecret, nil
}
