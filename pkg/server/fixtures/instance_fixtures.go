package fixtures

import (
	accountInstance "github.com/aligang/Gophkeeper/pkg/account/instance"
	secretInstance "github.com/aligang/Gophkeeper/pkg/secret/instance"
	tokenInstance "github.com/aligang/Gophkeeper/pkg/token/instance"
	"time"
)

const ReferenceAccountId1 = "test-account1"
const ReferenceLogin1 = "test-login1"
const ReferencePassword1 = "test-password1"
const ReferenceEncryptionKey1 = "test-key1"

const ReferenceAccountId2 = "test-account2"
const ReferenceLogin2 = "test-login2"
const ReferencePassword2 = "test-password2"
const ReferenceEncryptionKey2 = "test-key2"

var ReferenceAccountCreationTime, _ = time.Parse(time.RFC3339, "2023-02-07T14:00:02+03:00")

var ReferenceAccountInstance1 = &accountInstance.Account{
	Id:            ReferenceAccountId1,
	Login:         ReferenceLogin1,
	Password:      ReferencePassword1,
	EncryptionKey: ReferenceEncryptionKey1,
	CreatedAt:     ReferenceAccountCreationTime,
}

var ReferenceAccountInstance2 = &accountInstance.Account{
	Id:            ReferenceAccountId2,
	Login:         ReferenceLogin2,
	Password:      ReferencePassword2,
	EncryptionKey: ReferenceEncryptionKey2,
	CreatedAt:     ReferenceAccountCreationTime,
}

const ReferenceTokenId1 = "test-token1"
const ReferenceTokenValue1 = "test-token-value1"
const ReferenceTokenId11 = "test-token11"
const ReferenceTokenValue11 = "test-token-value11"
const ReferenceTokenId2 = "test-token2"
const ReferenceTokenValue2 = "test-token-value2"

var ReferenceTokenCreationTime1, _ = time.Parse(time.RFC3339, "3023-02-07T14:00:02+03:00")
var ReferenceTokenCreationTime11, _ = time.Parse(time.RFC3339, "3023-02-07T15:00:02+03:00")
var ReferenceTokenCreationTime2, _ = time.Parse(time.RFC3339, "3023-02-07T14:00:02+03:00")

var ReferenceTokenInstance1 = &tokenInstance.Token{
	Id:       ReferenceTokenId1,
	Value:    ReferenceTokenValue1,
	Owner:    ReferenceAccountId1,
	IssuedAt: ReferenceTokenCreationTime1,
}

var ReferenceTokenInstance11 = &tokenInstance.Token{
	Id:       ReferenceTokenId11,
	Value:    ReferenceTokenValue11,
	Owner:    ReferenceAccountId1,
	IssuedAt: ReferenceTokenCreationTime11,
}

var ReferenceTokenInstance2 = &tokenInstance.Token{
	Id:       ReferenceTokenId2,
	Value:    ReferenceTokenValue2,
	Owner:    ReferenceAccountId2,
	IssuedAt: ReferenceTokenCreationTime2,
}

const ReferenceSecretId1 = "test-text-id1"
const ReferenceSecretId11 = "test-text-id11"
const ReferenceSecretId2 = "test-text-id2"

var ReferenceSecretCreationTime1, _ = time.Parse(time.RFC3339, "3023-02-07T14:00:02+03:00")
var ReferenceSecretCreationTime11, _ = time.Parse(time.RFC3339, "3023-02-07T15:00:02+03:00")
var ReferenceSecretCreationTime2, _ = time.Parse(time.RFC3339, "3023-02-07T14:00:02+03:00")
var ReferenceSecretModificationTime1, _ = time.Parse(time.RFC3339, "3023-03-07T14:00:02+03:00")
var ReferenceSecretModificationTime11, _ = time.Parse(time.RFC3339, "3023-03-07T15:00:02+03:00")
var ReferenceSecretModificationTime2, _ = time.Parse(time.RFC3339, "3023-03-07T14:00:02+03:00")

const ReferenceTextSecretValue1 = "test-text1"
const ReferenceTextSecretValue11 = "test-text11"
const ReferenceTextSecretValue2 = "test-text2"

var ReferenceTextSecretInstance1 = &secretInstance.TextSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId1,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime1,
		ModifiedAt: ReferenceSecretModificationTime1,
	},
	Text: ReferenceTextSecretValue1,
}

var ReferenceTextSecretInstance11 = &secretInstance.TextSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId11,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime11,
		ModifiedAt: ReferenceSecretModificationTime11,
	},
	Text: ReferenceTextSecretValue11,
}

var ReferenceTextSecretInstance2 = &secretInstance.TextSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId2,
		AccountId:  ReferenceAccountId2,
		CreatedAt:  ReferenceSecretCreationTime2,
		ModifiedAt: ReferenceSecretModificationTime2,
	},
	Text: ReferenceTextSecretValue2,
}

const ReferenceLoginPasswordSecretLogin1 = "test-login1"
const ReferenceLoginPasswordSecretLogin11 = "test-login1"
const ReferenceLoginPasswordSecretLogin2 = "test-login2"

const ReferenceLoginPasswordSecretPassword1 = "test-password1"
const ReferenceLoginPasswordSecretPassword11 = "test-password1"
const ReferenceLoginPasswordSecretPassword2 = "test-password2"

var ReferenceLoginPasswordSecretInstance1 = &secretInstance.LoginPasswordSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId1,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime1,
		ModifiedAt: ReferenceSecretModificationTime1,
	},
	Login:    ReferenceLoginPasswordSecretLogin1,
	Password: ReferenceLoginPasswordSecretPassword1,
}

var ReferenceLoginPasswordSecretInstance11 = &secretInstance.LoginPasswordSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId11,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime11,
		ModifiedAt: ReferenceSecretModificationTime11,
	},
	Login:    ReferenceLoginPasswordSecretLogin11,
	Password: ReferenceLoginPasswordSecretPassword11,
}

var ReferenceLoginPasswordSecretInstance2 = &secretInstance.LoginPasswordSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId2,
		AccountId:  ReferenceAccountId2,
		CreatedAt:  ReferenceSecretCreationTime2,
		ModifiedAt: ReferenceSecretModificationTime2,
	},
	Login:    ReferenceLoginPasswordSecretLogin2,
	Password: ReferenceLoginPasswordSecretPassword2,
}

const ReferenceCreditCardSecretCardNumber1 = "1000 1000 1000 1000"
const ReferenceCreditCardSecretCardNumber11 = "1100 1100 1100 1100"
const ReferenceCreditCardSecretCardNumber2 = "2000 2000 2000 2000 "
const ReferenceCreditCardSecretCardHolder1 = "test-card-holder1"
const ReferenceCreditCardSecretCardHolder11 = "test-card-holder11"
const ReferenceCreditCardSecretCardHolder2 = "test-card-holder2"
const ReferenceCreditCardSecretValidTill1 = "01/50"
const ReferenceCreditCardSecretValidTill11 = "11/50"
const ReferenceCreditCardSecretValidTill2 = "02/50"
const ReferenceCreditCardSecretCvc1 = "100"
const ReferenceCreditCardSecretCvc11 = "110"
const ReferenceCreditCardSecretCvc2 = "200"

var ReferenceCreditCardSecretInstance1 = &secretInstance.CreditCardSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId1,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime1,
		ModifiedAt: ReferenceSecretModificationTime1,
	},
	Number:     ReferenceCreditCardSecretCardNumber1,
	CardHolder: ReferenceCreditCardSecretCardHolder1,
	ValidTill:  ReferenceCreditCardSecretValidTill1,
	Cvc:        ReferenceCreditCardSecretCvc1,
}

var ReferenceCreditCardSecretInstance11 = &secretInstance.CreditCardSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId11,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime11,
		ModifiedAt: ReferenceSecretModificationTime11,
	},
	Number:     ReferenceCreditCardSecretCardNumber11,
	CardHolder: ReferenceCreditCardSecretCardHolder11,
	ValidTill:  ReferenceCreditCardSecretValidTill11,
	Cvc:        ReferenceCreditCardSecretCvc11,
}

var ReferenceCreditCardSecretInstance2 = &secretInstance.CreditCardSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId2,
		AccountId:  ReferenceAccountId2,
		CreatedAt:  ReferenceSecretCreationTime2,
		ModifiedAt: ReferenceSecretModificationTime2,
	},
	Number:     ReferenceCreditCardSecretCardNumber2,
	CardHolder: ReferenceCreditCardSecretCardHolder2,
	ValidTill:  ReferenceCreditCardSecretValidTill2,
	Cvc:        ReferenceCreditCardSecretCvc2,
}

const ReferenceObjectId1 = "test-object-1"
const ReferenceObjectId11 = "test-object-11"
const ReferenceObjectId2 = "test-object-2"

var ReferenceFileContent1 = []byte("test-object-1")
var ReferenceFileContent11 = []byte("test-object-11")
var ReferenceFileContent2 = []byte("test-object-2")

var ReferenceSecretDeletionTime1 = ReferenceSecretCreationTime1.Add(10 * time.Hour)
var ReferenceSecretDeletionTime11 = ReferenceSecretCreationTime11.Add(10 * time.Hour)
var ReferenceSecretDeletionTime2 = ReferenceSecretCreationTime2.Add(10 * time.Hour)

var ReferenceFileSecretInstance1 = &secretInstance.FileSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId1,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime1,
		ModifiedAt: ReferenceSecretModificationTime1,
	},
	ObjectId: ReferenceObjectId1,
}

var ReferenceFileSecretInstance11 = &secretInstance.FileSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId11,
		AccountId:  ReferenceAccountId1,
		CreatedAt:  ReferenceSecretCreationTime11,
		ModifiedAt: ReferenceSecretModificationTime11,
	},
	ObjectId: ReferenceObjectId11,
}

var ReferenceFileSecretInstance2 = &secretInstance.FileSecret{
	BaseSecret: secretInstance.BaseSecret{
		Id:         ReferenceSecretId2,
		AccountId:  ReferenceAccountId2,
		CreatedAt:  ReferenceSecretCreationTime2,
		ModifiedAt: ReferenceSecretModificationTime2,
	},
	ObjectId: ReferenceObjectId2,
}
