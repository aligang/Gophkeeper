package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/fixtures"
	secretInstance "github.com/aligang/Gophkeeper/pkg/secret/instance"
	"github.com/aligang/Gophkeeper/pkg/server/repository/repositoryerrors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFileSecretWritableOperation(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name              string
		inputDBContent    *databaseContent
		expectedDBContent *databaseContent
		run               func(r *Repository) error
	}{
		{
			name: "add first secret",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance1)
			},
		},
		{
			name: "add second secret",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId11,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance11)
			},
		},
		{
			name: "add secret for second account",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId11,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId11,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance2)
			},
		},
		{
			name: "update second secret",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId11,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11.Add(time.Hour),
						fileRecord: fileRecord{
							objectId: "Updated object Id",
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {

				return r.UpdateFileSecret(ctx, &secretInstance.FileSecret{
					BaseSecret: secretInstance.BaseSecret{
						Id:         fixtures.ReferenceSecretId11,
						AccountId:  fixtures.ReferenceAccountId1,
						CreatedAt:  fixtures.ReferenceSecretCreationTime11,
						ModifiedAt: fixtures.ReferenceSecretModificationTime11.Add(time.Hour),
					},
					ObjectId: "Updated object Id",
				})
			},
		},
		{
			name: "Delete non-single secret",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId11,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteFileSecret(ctx, fixtures.ReferenceSecretId11)
			},
		},
		{
			name: "Delete last secret",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId1,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteFileSecret(ctx, fixtures.ReferenceSecretId1)
			},
		},
		{
			name: "Move file secret to deletion queue",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{
					fixtures.ReferenceObjectId1: fixtures.ReferenceSecretDeletionTime1,
				},
			},
			run: func(r *Repository) error {
				return r.MoveFileSecretToDeletionQueue(ctx, fixtures.ReferenceObjectId1, fixtures.ReferenceSecretDeletionTime1)
			},
		},
		{
			name: "Move file secret from deletion queue",
			inputDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{
					fixtures.ReferenceObjectId1: fixtures.ReferenceSecretDeletionTime1,
				},
			},
			expectedDBContent: &databaseContent{
				accounts: AccountIdToAccountMapping{
					fixtures.ReferenceAccountId1: &accountRecord{
						id:            fixtures.ReferenceAccountId1,
						login:         fixtures.ReferenceLogin1,
						password:      fixtures.ReferencePassword1,
						encryptionKey: fixtures.ReferenceEncryptionKey1,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
					fixtures.ReferenceAccountId2: &accountRecord{
						id:            fixtures.ReferenceAccountId2,
						login:         fixtures.ReferenceLogin2,
						password:      fixtures.ReferencePassword2,
						encryptionKey: fixtures.ReferenceEncryptionKey2,
						createdAt:     fixtures.ReferenceAccountCreationTime,
					},
				},
				loginToIdMapping: LoginToAccountIdMapping{
					fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
					fixtures.ReferenceLogin2: fixtures.ReferenceAccountId2,
				},
				tokens: TokenValueToTokenMapping{
					fixtures.ReferenceTokenValue1: &tokenRecord{
						id:       fixtures.ReferenceTokenId1,
						value:    fixtures.ReferenceTokenValue1,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime1,
					},
					fixtures.ReferenceTokenValue11: &tokenRecord{
						id:       fixtures.ReferenceTokenId11,
						value:    fixtures.ReferenceTokenValue11,
						owner:    fixtures.ReferenceAccountId1,
						issuedAt: fixtures.ReferenceTokenCreationTime11,
					},
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceTokenValue2: nil,
					},
				},

				textSecrets:        SecretIdToSecretMapping{},
				accountTextSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						fileRecord: fileRecord{
							objectId: fixtures.ReferenceObjectId2,
						},
					},
				},
				accountFileSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},
				fileDeletionQueue: DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteFileSecretFromDeletionQueue(ctx, fixtures.ReferenceObjectId1)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := Init(
				test.inputDBContent.accounts,
				test.inputDBContent.loginToIdMapping,
				test.inputDBContent.tokens,
				test.inputDBContent.accountTokens,
				test.inputDBContent.textSecrets,
				test.inputDBContent.accountTextSecrets,
				test.inputDBContent.loginPasswordSecrets,
				test.inputDBContent.accountLoginPasswordSecrets,
				test.inputDBContent.creditCardSecrets,
				test.inputDBContent.accountCreditCardSecrets,
				test.inputDBContent.fileSecrets,
				test.inputDBContent.accountFileSecrets,
				test.inputDBContent.fileDeletionQueue,
			)
			err := test.run(repo)

			assert.Equal(t, nil, err)
			assert.Equal(t, test.expectedDBContent, repo.dump())
		})
	}
}

func TestFileSecretReadOperations(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)
	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)

	t.Run("empty file secret list", func(t *testing.T) {
		_, err := repo.ListFileSecrets(ctx, fixtures.ReferenceAccountId1)
		assert.ErrorIs(t, nil, err)
	})

	_ = repo.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance1)
	_ = repo.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance11)
	_ = repo.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance2)

	t.Run("get existing instance", func(t *testing.T) {
		instance, err := repo.GetFileSecret(ctx, fixtures.ReferenceSecretId1)
		assert.Equal(t, nil, err)
		assert.Equal(t, fixtures.ReferenceFileSecretInstance1, instance)
	})

	t.Run("get non-existing instance", func(t *testing.T) {
		_, err := repo.GetFileSecret(ctx, "this text secret does not exist")
		assert.ErrorIs(t, repositoryerrors.ErrRecordNotFound, err)
	})
	t.Run("list all secrets", func(t *testing.T) {
		list, err := repo.ListFileSecrets(ctx, fixtures.ReferenceAccountId1)
		assert.Equal(t, nil, err)
		resultMap := map[string]*secretInstance.FileSecret{}
		for _, e := range list {
			resultMap[e.Id] = e
		}
		assert.Equal(
			t,
			map[string]*secretInstance.FileSecret{
				fixtures.ReferenceSecretId1:  fixtures.ReferenceFileSecretInstance1,
				fixtures.ReferenceSecretId11: fixtures.ReferenceFileSecretInstance11,
			},
			resultMap)
	})

	_ = repo.DeleteFileSecret(ctx, fixtures.ReferenceSecretId1)
	_ = repo.MoveFileSecretToDeletionQueue(ctx, fixtures.ReferenceObjectId1, fixtures.ReferenceSecretDeletionTime1)
	_ = repo.DeleteFileSecret(ctx, fixtures.ReferenceSecretId11)
	_ = repo.MoveFileSecretToDeletionQueue(ctx, fixtures.ReferenceObjectId11, fixtures.ReferenceSecretDeletionTime11)
	_ = repo.DeleteFileSecret(ctx, fixtures.ReferenceSecretId2)
	_ = repo.MoveFileSecretToDeletionQueue(ctx, fixtures.ReferenceObjectId2, fixtures.ReferenceSecretDeletionTime2)

	t.Run("list all objects in deletion queue", func(t *testing.T) {
		list, err := repo.ListFileDeletionQueue(ctx)
		assert.Equal(t, nil, err)
		resultMap := map[string]*secretInstance.DeletionQueueElement{}
		for _, e := range list {
			resultMap[e.Id] = e
		}
		assert.Equal(
			t,
			map[string]*secretInstance.DeletionQueueElement{
				fixtures.ReferenceObjectId1: &secretInstance.DeletionQueueElement{
					Id: fixtures.ReferenceObjectId1,
					Ts: fixtures.ReferenceSecretDeletionTime1,
				},
				fixtures.ReferenceObjectId11: &secretInstance.DeletionQueueElement{
					Id: fixtures.ReferenceObjectId11,
					Ts: fixtures.ReferenceSecretDeletionTime11,
				},
				fixtures.ReferenceObjectId2: &secretInstance.DeletionQueueElement{
					Id: fixtures.ReferenceObjectId2,
					Ts: fixtures.ReferenceSecretDeletionTime2,
				},
			},
			resultMap)
	})
}
