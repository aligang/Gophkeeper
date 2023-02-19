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

func TestTextSecretWritableOperation(t *testing.T) {
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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
				},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance1)
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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
				},

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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
				},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance11)
			},
		},
		{
			name: "add  secret for another account",
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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
				},

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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance2)
			},
		},
		{
			name: "update secret",
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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1.Add(time.Hour),
						text:       "updated text secret",
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {

				return r.UpdateTextSecret(ctx, &secretInstance.TextSecret{
					BaseSecret: secretInstance.BaseSecret{
						Id:         fixtures.ReferenceSecretId1,
						AccountId:  fixtures.ReferenceAccountId1,
						CreatedAt:  fixtures.ReferenceSecretCreationTime1,
						ModifiedAt: fixtures.ReferenceSecretModificationTime1.Add(time.Hour),
					},
					Text: "updated text secret",
				})
			},
		},
		{
			name: "Delete non-last secret",
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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						text:       fixtures.ReferenceTextSecretValue1,
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteTextSecret(ctx, fixtures.ReferenceSecretId1)
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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						text:       fixtures.ReferenceTextSecretValue11,
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

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

				textSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						text:       fixtures.ReferenceTextSecretValue2,
					},
				},
				accountTextSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteTextSecret(ctx, fixtures.ReferenceSecretId11)
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

func TestTextSecretReadOperations(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)
	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)

	t.Run("empty text secret list", func(t *testing.T) {
		_, err := repo.ListTextSecrets(ctx, fixtures.ReferenceAccountId1)
		assert.ErrorIs(t, nil, err)
	})

	_ = repo.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance1)
	_ = repo.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance11)
	_ = repo.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance2)

	t.Run("get existing instance", func(t *testing.T) {
		instance, err := repo.GetTextSecret(ctx, fixtures.ReferenceSecretId1)
		assert.Equal(t, nil, err)
		assert.Equal(t, fixtures.ReferenceTextSecretInstance1, instance)
	})

	t.Run("get non-existing instance", func(t *testing.T) {
		_, err := repo.GetTextSecret(ctx, "this text secret does not exist")
		assert.ErrorIs(t, repositoryerrors.ErrRecordNotFound, err)
	})
	t.Run("list all secrets", func(t *testing.T) {
		list, err := repo.ListTextSecrets(ctx, fixtures.ReferenceAccountId1)
		assert.Equal(t, nil, err)
		resultMap := map[string]*secretInstance.TextSecret{}
		for _, e := range list {
			resultMap[e.Id] = e
		}
		assert.Equal(
			t,
			map[string]*secretInstance.TextSecret{
				fixtures.ReferenceSecretId1:  fixtures.ReferenceTextSecretInstance1,
				fixtures.ReferenceSecretId11: fixtures.ReferenceTextSecretInstance11,
			},
			resultMap)
	})

}
