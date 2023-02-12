package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {
	inMemory := New()

	t.Run("REGISTER New Account test", func(t *testing.T) {
		expectedContent := &databaseContent{
			accounts: AccountIdToAccountMapping{
				fixtures.ReferenceAccountId1: &accountRecord{
					id:            fixtures.ReferenceAccountId1,
					login:         fixtures.ReferenceLogin1,
					password:      fixtures.ReferencePassword1,
					encryptionKey: fixtures.ReferenceEncryptionKey1,
					createdAt:     fixtures.ReferenceAccountCreationTime,
				},
			},
			loginToIdMapping: LoginToAccountIdMapping{
				fixtures.ReferenceLogin1: fixtures.ReferenceAccountId1,
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
		}

		err := inMemory.Register(
			context.Background(),
			fixtures.ReferenceAccountInstance1,
		)
		assert.Equal(t, nil, err)
		assert.Equal(t, expectedContent, inMemory.dump())
	})

	t.Run("GET BY LOGIN Account test", func(t *testing.T) {
		accountInstance, err := inMemory.GetAccountByLogin(
			context.Background(),
			fixtures.ReferenceLogin1,
		)
		assert.Equal(t, nil, err)
		assert.Equal(t, fixtures.ReferenceAccountInstance1, accountInstance)
	})

	t.Run("GET BY ID Account test", func(t *testing.T) {
		accountInstance, err := inMemory.GetAccountById(
			context.Background(),
			fixtures.ReferenceAccountId1,
		)
		assert.Equal(t, nil, err)
		assert.Equal(t, fixtures.ReferenceAccountInstance1, accountInstance)
	})
}
