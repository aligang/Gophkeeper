package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/fixtures"
	"github.com/aligang/Gophkeeper/internal/repository/repositoryerrors"
	secretInstance "github.com/aligang/Gophkeeper/internal/secret/instance"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoginPasswordSecretWritableOperation(t *testing.T) {
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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
				},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddLoginPasswordSecret(ctx, fixtures.ReferenceLoginPasswordSecretInstance1)
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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
				},

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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin11,
							password: fixtures.ReferenceLoginPasswordSecretPassword11,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
				},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddLoginPasswordSecret(ctx, fixtures.ReferenceLoginPasswordSecretInstance11)
			},
		},
		{
			name: "add secret for another account",
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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin11,
							password: fixtures.ReferenceLoginPasswordSecretPassword11,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
				},

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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin11,
							password: fixtures.ReferenceLoginPasswordSecretPassword11,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddLoginPasswordSecret(ctx, fixtures.ReferenceLoginPasswordSecretInstance2)
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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin11,
							password: fixtures.ReferenceLoginPasswordSecretPassword11,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11.Add(time.Hour),
						loginPasswordRecord: loginPasswordRecord{
							login:    "updated secret-login",
							password: "updated secret-password",
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {

				return r.UpdateLoginPasswordSecret(ctx, &secretInstance.LoginPasswordSecret{
					BaseSecret: secretInstance.BaseSecret{
						Id:         fixtures.ReferenceSecretId11,
						AccountId:  fixtures.ReferenceAccountId1,
						CreatedAt:  fixtures.ReferenceSecretCreationTime11,
						ModifiedAt: fixtures.ReferenceSecretModificationTime11.Add(time.Hour),
					},
					Login:    "updated secret-login",
					Password: "updated secret-password",
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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId11: &SecretRecord{
						id:         fixtures.ReferenceSecretId11,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime11,
						modifiedAt: fixtures.ReferenceSecretModificationTime11,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin11,
							password: fixtures.ReferenceLoginPasswordSecretPassword11,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1:  nil,
						fixtures.ReferenceSecretId11: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceSecretId1: nil,
					},
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteLoginPasswordSecret(ctx, fixtures.ReferenceSecretId11)
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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId1: &SecretRecord{
						id:         fixtures.ReferenceSecretId1,
						accountId:  fixtures.ReferenceAccountId1,
						createdAt:  fixtures.ReferenceSecretCreationTime1,
						modifiedAt: fixtures.ReferenceSecretModificationTime1,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin1,
							password: fixtures.ReferenceLoginPasswordSecretPassword1,
						},
					},
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

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

				loginPasswordSecrets: SecretIdToSecretMapping{
					fixtures.ReferenceSecretId2: &SecretRecord{
						id:         fixtures.ReferenceSecretId2,
						accountId:  fixtures.ReferenceAccountId2,
						createdAt:  fixtures.ReferenceSecretCreationTime2,
						modifiedAt: fixtures.ReferenceSecretModificationTime2,
						loginPasswordRecord: loginPasswordRecord{
							login:    fixtures.ReferenceLoginPasswordSecretLogin2,
							password: fixtures.ReferenceLoginPasswordSecretPassword2,
						},
					},
				},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{
					fixtures.ReferenceAccountId2: {
						fixtures.ReferenceSecretId2: nil,
					},
				},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.DeleteLoginPasswordSecret(ctx, fixtures.ReferenceSecretId1)
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

func TestLoginPasswordSecretReadOperations(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)
	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)

	t.Run("empty login-password secret list", func(t *testing.T) {
		_, err := repo.ListLoginPasswordSecrets(ctx, fixtures.ReferenceAccountId1)
		assert.ErrorIs(t, nil, err)
	})

	_ = repo.AddLoginPasswordSecret(ctx, fixtures.ReferenceLoginPasswordSecretInstance1)
	_ = repo.AddLoginPasswordSecret(ctx, fixtures.ReferenceLoginPasswordSecretInstance11)
	_ = repo.AddLoginPasswordSecret(ctx, fixtures.ReferenceLoginPasswordSecretInstance2)

	t.Run("get existing instance", func(t *testing.T) {
		instance, err := repo.GetLoginPasswordSecret(ctx, fixtures.ReferenceSecretId1)
		assert.Equal(t, nil, err)
		assert.Equal(t, fixtures.ReferenceLoginPasswordSecretInstance1, instance)
	})

	t.Run("get non-existing instance", func(t *testing.T) {
		_, err := repo.GetLoginPasswordSecret(ctx, "this text secret does not exist")
		assert.ErrorIs(t, repositoryerrors.ErrRecordNotFound, err)
	})
	t.Run("list all secrets", func(t *testing.T) {
		list, err := repo.ListLoginPasswordSecrets(ctx, fixtures.ReferenceAccountId1)
		assert.Equal(t, nil, err)
		resultMap := map[string]*secretInstance.LoginPasswordSecret{}
		for _, e := range list {
			resultMap[e.Id] = e
		}
		assert.Equal(
			t,
			map[string]*secretInstance.LoginPasswordSecret{
				fixtures.ReferenceSecretId1:  fixtures.ReferenceLoginPasswordSecretInstance1,
				fixtures.ReferenceSecretId11: fixtures.ReferenceLoginPasswordSecretInstance11,
			},
			resultMap)
	})
}
