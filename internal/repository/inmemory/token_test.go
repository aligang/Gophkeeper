package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/fixtures"
	"github.com/aligang/Gophkeeper/internal/repository/repositoryerrors"
	tokenInstance "github.com/aligang/Gophkeeper/internal/token/instance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenWritableOperation(t *testing.T) {

	tests := []struct {
		name              string
		inputDBContent    *databaseContent
		expectedDBContent *databaseContent
		run               func(r *Repository) error
	}{
		{
			name: "add first token",
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
				tokens:        TokenValueToTokenMapping{},
				accountTokens: AccountIdToTokenValueMapping{},

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
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1: nil,
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
			run: func(r *Repository) error {
				return r.AddToken(context.Background(), fixtures.ReferenceTokenInstance1)
			},
		},
		{
			name: "add second token",
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
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1: nil,
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
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
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
			run: func(r *Repository) error {
				return r.AddToken(context.Background(), fixtures.ReferenceTokenInstance11)
			},
		},
		{
			name: "add second token",
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
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1:  nil,
						fixtures.ReferenceTokenValue11: nil,
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

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},
				fileDeletionQueue:  DeletionQueueMapping{},
			},
			run: func(r *Repository) error {
				return r.AddToken(context.Background(), fixtures.ReferenceTokenInstance2)
			},
		},
		{
			name: "Delete non-single token",
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

				fileSecrets:        SecretIdToSecretMapping{},
				accountFileSecrets: AccountIdToSecretIdMapping{},

				loginPasswordSecrets:        SecretIdToSecretMapping{},
				accountLoginPasswordSecrets: AccountIdToSecretIdMapping{},

				creditCardSecrets:        SecretIdToSecretMapping{},
				accountCreditCardSecrets: AccountIdToSecretIdMapping{},

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
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1: nil,
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
			run: func(r *Repository) error {
				return r.DeleteToken(context.Background(), fixtures.ReferenceTokenInstance11)
			},
		},
		{
			name: "Delete last token",
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
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
					fixtures.ReferenceAccountId1: {
						fixtures.ReferenceTokenValue1: nil,
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
					fixtures.ReferenceTokenValue2: &tokenRecord{
						id:       fixtures.ReferenceTokenId2,
						value:    fixtures.ReferenceTokenValue2,
						owner:    fixtures.ReferenceAccountId2,
						issuedAt: fixtures.ReferenceTokenCreationTime2,
					},
				},
				accountTokens: AccountIdToTokenValueMapping{
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
			run: func(r *Repository) error {
				return r.DeleteToken(context.Background(), fixtures.ReferenceTokenInstance1)
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

func TestTokenReadOperations(t *testing.T) {
	repo := New()
	ctx := context.Background()

	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)
	_ = repo.Register(ctx, fixtures.ReferenceAccountInstance1)

	t.Run("list account tokens who has not tokens", func(t *testing.T) {
		_, err := repo.ListAccountTokens(ctx, fixtures.ReferenceAccountId1)
		assert.ErrorIs(t, nil, err)
	})

	_ = repo.AddToken(ctx, fixtures.ReferenceTokenInstance1)
	_ = repo.AddToken(ctx, fixtures.ReferenceTokenInstance11)
	_ = repo.AddToken(ctx, fixtures.ReferenceTokenInstance2)

	t.Run("get existing instance", func(t *testing.T) {
		instance, err := repo.GetToken(ctx, fixtures.ReferenceTokenValue1)
		assert.Equal(t, nil, err)
		assert.Equal(t, fixtures.ReferenceTokenInstance1, instance)
	})

	t.Run("get non-existing instance", func(t *testing.T) {
		_, err := repo.GetToken(ctx, "this token does not exist")
		assert.ErrorIs(t, repositoryerrors.ErrRecordNotFound, err)
	})
	t.Run("list all tokens", func(t *testing.T) {
		list, err := repo.ListTokens(ctx)
		assert.Equal(t, nil, err)
		resultMap := map[string]*tokenInstance.Token{}
		for _, e := range list {
			resultMap[e.Id] = e
		}
		assert.Equal(
			t,
			map[string]*tokenInstance.Token{
				fixtures.ReferenceTokenId1:  fixtures.ReferenceTokenInstance1,
				fixtures.ReferenceTokenId11: fixtures.ReferenceTokenInstance11,
				fixtures.ReferenceTokenId2:  fixtures.ReferenceTokenInstance2,
			},
			resultMap)
	})
	t.Run("list account tokens", func(t *testing.T) {
		list, err := repo.ListAccountTokens(ctx, fixtures.ReferenceAccountId1)
		assert.Equal(t, nil, err)
		resultMap := map[string]*tokenInstance.Token{}
		for _, e := range list {
			resultMap[e.Id] = e
		}
		assert.Equal(
			t,
			map[string]*tokenInstance.Token{
				fixtures.ReferenceTokenId1:  fixtures.ReferenceTokenInstance1,
				fixtures.ReferenceTokenId11: fixtures.ReferenceTokenInstance11,
			},
			resultMap)
	})
}
