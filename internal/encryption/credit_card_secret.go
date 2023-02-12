package encryption

import secretInstance "github.com/aligang/Gophkeeper/internal/secret/instance"

func EncryptCreditCardSecret(plainSecret *secretInstance.CreditCardSecret,
	encryptionKey string) (*secretInstance.CreditCardSecret, error) {
	var err error
	encryptedSecret := &secretInstance.CreditCardSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         plainSecret.Id,
			AccountId:  plainSecret.AccountId,
			CreatedAt:  plainSecret.CreatedAt,
			ModifiedAt: plainSecret.ModifiedAt,
		},
	}

	encryptedSecret.Number, err = encrypt(plainSecret.Number, encryptionKey)
	if err != nil {
		return nil, err
	}
	encryptedSecret.CardHolder, err = encrypt(plainSecret.CardHolder, encryptionKey)
	if err != nil {
		return nil, err
	}
	encryptedSecret.ValidTill, err = encrypt(plainSecret.ValidTill, encryptionKey)
	if err != nil {
		return nil, err
	}
	encryptedSecret.Cvc, err = encrypt(plainSecret.Cvc, encryptionKey)
	if err != nil {
		return nil, err
	}

	return encryptedSecret, nil

}

func DecryptCreditCardSecret(encryptedSecret *secretInstance.CreditCardSecret,
	encryptionKey string) (*secretInstance.CreditCardSecret, error) {
	var err error
	plainSecret := &secretInstance.CreditCardSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         encryptedSecret.Id,
			AccountId:  encryptedSecret.AccountId,
			CreatedAt:  encryptedSecret.CreatedAt,
			ModifiedAt: encryptedSecret.ModifiedAt,
		},
	}

	plainSecret.Number, err = decrypt(encryptedSecret.Number, encryptionKey)
	if err != nil {
		return nil, err
	}
	plainSecret.CardHolder, err = decrypt(encryptedSecret.CardHolder, encryptionKey)
	if err != nil {
		return nil, err
	}
	plainSecret.ValidTill, err = decrypt(encryptedSecret.ValidTill, encryptionKey)
	if err != nil {
		return nil, err
	}
	plainSecret.Cvc, err = decrypt(encryptedSecret.Cvc, encryptionKey)
	if err != nil {
		return nil, err
	}
	return plainSecret, nil
}
