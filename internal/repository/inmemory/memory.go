package inmemory

import (
	"fmt"
	"github.com/aligang/Gophkeeper/internal/logging"
	"reflect"
	"sync"
	"time"
)

type AccountIdToAccountMapping map[string]*accountRecord
type LoginToAccountIdMapping map[string]string

type TokenValueToTokenMapping map[string]*tokenRecord
type AccountIdToTokenValueMapping map[string]map[string]any

type SecretIdToSecretMapping map[string]*SecretRecord
type AccountIdToSecretIdMapping map[string]map[string]any

type DeletionQueueMapping map[string]time.Time

type SecretRecord struct {
	id                  string
	accountId           string
	createdAt           time.Time
	modifiedAt          time.Time
	text                string
	loginPasswordRecord loginPasswordRecord
	creditCardRecord    creditCardRecord
	fileRecord          fileRecord
}

type databaseContent struct {
	accounts         AccountIdToAccountMapping
	loginToIdMapping LoginToAccountIdMapping

	tokens        TokenValueToTokenMapping
	accountTokens AccountIdToTokenValueMapping

	textSecrets        SecretIdToSecretMapping
	accountTextSecrets AccountIdToSecretIdMapping

	loginPasswordSecrets        SecretIdToSecretMapping
	accountLoginPasswordSecrets AccountIdToSecretIdMapping

	creditCardSecrets        SecretIdToSecretMapping
	accountCreditCardSecrets AccountIdToSecretIdMapping

	fileSecrets        SecretIdToSecretMapping
	accountFileSecrets AccountIdToSecretIdMapping
	fileDeletionQueue  DeletionQueueMapping
}

type Repository struct {
	log  *logging.InternalLogger
	Lock sync.Mutex

	databaseContent
}

func New() *Repository {
	logging.Debug("Initialization In-MemoryStorage Backend")
	m := &Repository{
		log: logging.Logger.GetSubLogger("repository", "IN-Memory"),
		databaseContent: databaseContent{
			accounts:         AccountIdToAccountMapping{},
			loginToIdMapping: LoginToAccountIdMapping{},

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
	}
	logging.Debug("Initialization IN-Memory Storage Backend is Finished")
	return m
}

func Init(
	accounts AccountIdToAccountMapping, loginToIdMapping LoginToAccountIdMapping,
	tokens TokenValueToTokenMapping, accountTokens AccountIdToTokenValueMapping,
	textSecrets SecretIdToSecretMapping, accountTextSecrets AccountIdToSecretIdMapping,
	loginPasswordSecrets SecretIdToSecretMapping, accountLoginPasswordSecrets AccountIdToSecretIdMapping,
	creditCardSecrets SecretIdToSecretMapping, accountCreditCardsSecrets AccountIdToSecretIdMapping,
	fileSecrets SecretIdToSecretMapping, accountFileSecrets AccountIdToSecretIdMapping,
	fileDeletionQueue DeletionQueueMapping) *Repository {

	m := &Repository{
		log: logging.Logger.GetSubLogger("database", "IN-Memory"),
		databaseContent: databaseContent{
			accounts:         accounts,
			loginToIdMapping: loginToIdMapping,

			tokens:        tokens,
			accountTokens: accountTokens,

			textSecrets:        textSecrets,
			accountTextSecrets: accountTextSecrets,

			loginPasswordSecrets:        loginPasswordSecrets,
			accountLoginPasswordSecrets: accountLoginPasswordSecrets,

			creditCardSecrets:        creditCardSecrets,
			accountCreditCardSecrets: accountCreditCardsSecrets,

			fileSecrets:        fileSecrets,
			accountFileSecrets: accountFileSecrets,
			fileDeletionQueue:  fileDeletionQueue,
		},
	}
	return m
}

func (r *Repository) dump() *databaseContent {
	return &databaseContent{
		accounts:         r.accounts,
		loginToIdMapping: r.loginToIdMapping,

		tokens:        r.tokens,
		accountTokens: r.accountTokens,

		textSecrets:        r.textSecrets,
		accountTextSecrets: r.accountTextSecrets,

		loginPasswordSecrets:        r.loginPasswordSecrets,
		accountLoginPasswordSecrets: r.accountLoginPasswordSecrets,

		creditCardSecrets:        r.creditCardSecrets,
		accountCreditCardSecrets: r.accountCreditCardSecrets,

		fileSecrets:        r.fileSecrets,
		accountFileSecrets: r.accountFileSecrets,
		fileDeletionQueue:  r.fileDeletionQueue,
	}
}

func (r *Repository) Equals(other *Repository) bool {
	if !reflect.DeepEqual(r.accounts, other.accounts) {
		fmt.Println("accounts maps are not equal")
		return false
	}
	if !reflect.DeepEqual(r.loginToIdMapping, other.loginToIdMapping) {
		fmt.Println("logging to Id  maps are not equal")
		return false
	}
	if !reflect.DeepEqual(r.tokens, other.tokens) {
		fmt.Println("token maps are not equal")
		return false
	}
	if !reflect.DeepEqual(r.accountTokens, other.accountTokens) {
		fmt.Println("account token maps are not equal")
		return false
	}

	if !reflect.DeepEqual(r.textSecrets, other.textSecrets) {
		fmt.Println("Text secrets are not equal")
		fmt.Println(r.textSecrets)
		fmt.Println(other.textSecrets)
		return false
	}

	if !reflect.DeepEqual(r.accountTextSecrets, other.accountTextSecrets) {
		fmt.Println("Account text secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.loginPasswordSecrets, other.loginPasswordSecrets) {
		fmt.Println("Login password secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.accountLoginPasswordSecrets, other.accountLoginPasswordSecrets) {
		fmt.Println("Account text secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.creditCardSecrets, other.creditCardSecrets) {
		fmt.Println("Credit card secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.accountCreditCardSecrets, other.accountCreditCardSecrets) {
		fmt.Println("Account credit card secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.fileSecrets, other.fileSecrets) {
		fmt.Println("File secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.accountFileSecrets, other.accountFileSecrets) {
		fmt.Println("Account credit card secrets are not equal")
		return false
	}

	if !reflect.DeepEqual(r.fileDeletionQueue, other.fileDeletionQueue) {
		fmt.Println("File deletion queues are not equal")
		return false
	}

	return true
}
