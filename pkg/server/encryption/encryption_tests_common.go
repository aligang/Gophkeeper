package encryption

import (
	secretInstance "github.com/aligang/Gophkeeper/pkg/common/secret/instance"
	"time"
)

const (
	referenceSecretId   string = "this is test secret"
	referenceAccountId  string = "this is test account ID"
	referenceTextData   string = "test data"
	referenceLogin      string = "test login"
	referencePassword   string = "test password"
	referenceCardNumber string = "1111111111111111"
	referenceCardHolder string = "john doe"
	referenceValidTill  string = "01/01"
	referenceCvc        string = "000"
)

var (
	referenceEncryptionKey = func() string {
		key := make([]byte, KeySize)
		for idx, _ := range key {
			key[idx] = '0'
		}
		return string(key)
	}()

	referenceCreationTime, _     = time.Parse(time.RFC3339, "2023-02-07T14:00:02+03:00")
	referenceModificationTime, _ = time.Parse(time.RFC3339, "2023-02-08T14:00:02+03:00")
	referenceTextSecret          = &secretInstance.TextSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         referenceSecretId,
			AccountId:  referenceAccountId,
			CreatedAt:  referenceCreationTime,
			ModifiedAt: referenceModificationTime,
		},
		Text: referenceTextData,
	}
	referenceLoginPasswordSecret = &secretInstance.LoginPasswordSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         referenceSecretId,
			AccountId:  referenceAccountId,
			CreatedAt:  referenceCreationTime,
			ModifiedAt: referenceModificationTime,
		},
		Login:    referenceLogin,
		Password: referencePassword,
	}
	referenceCreditCardSecret = &secretInstance.CreditCardSecret{
		BaseSecret: secretInstance.BaseSecret{
			Id:         referenceSecretId,
			AccountId:  referenceAccountId,
			CreatedAt:  referenceCreationTime,
			ModifiedAt: referenceModificationTime,
		},
		CardNumber: referenceCardNumber,
		CardHolder: referenceCardHolder,
		ValidTill:  referenceValidTill,
		Cvc:        referenceCvc,
	}
)
