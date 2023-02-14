package handler

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/fixtures"
	"github.com/aligang/Gophkeeper/internal/repository/fs"
	"github.com/aligang/Gophkeeper/internal/repository/inmemory"
	"github.com/aligang/Gophkeeper/internal/secret"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

type operationalDelta string

const (
	operationalDeltaINCREASED operationalDelta = "ADD"
	operationalDeltaDECREASED operationalDelta = "DELETE"
	operationalDeltaUNCHANGED operationalDelta = "UNCHAGED"
)

type handlerTestInput struct {
	accountId string ``
}

type handlerTestExpected struct {
	expectedRunResponse   *secret.SecretDescription
	expectedRunErrorMsg   error
	expectedSecretType    secret.SecretType
	expectedCheckErrorMsg error
	delta                 operationalDelta
}

func TestWritableOperation(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		run      func(ctx context.Context, r *GrpcHandler) (*secret.SecretDescription, error)
		check    func(ctx context.Context, r *GrpcHandler, s *secret.SecretDescription) (*secret.Secret, error)
		Expected handlerTestExpected
		Input    handlerTestInput
	}{
		{
			name: "Create First Text Secret",
			Input: handlerTestInput{
				accountId: fixtures.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret.SecretType_TEXT,
				delta:                 operationalDeltaINCREASED,
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret.SecretDescription, error) {
				return h.Create(ctx, &secret.CreateSecretRequest{
					Secret: &secret.CreateSecretRequest_Text{
						Text: &secret.PlainText{
							Data: fixtures.ReferenceTextSecretValue1,
						},
					},
				})
			},
			check: func(ctx context.Context, h *GrpcHandler, s *secret.SecretDescription) (*secret.Secret, error) {
				return h.Get(ctx, &secret.GetSecretRequest{
					Id:         s.Id,
					SecretType: s.SecretType,
				})
			},
		},
		{
			name: "Create First File Secret",
			Input: handlerTestInput{
				accountId: fixtures.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg: nil,
				expectedSecretType:  secret.SecretType_FILE,
				delta:               operationalDeltaINCREASED,
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret.SecretDescription, error) {
				return h.Create(ctx, &secret.CreateSecretRequest{
					Secret: &secret.CreateSecretRequest_File{
						File: &secret.File{
							Data: fixtures.ReferenceFileContent1,
						},
					},
				})
			},
			check: func(ctx context.Context, h *GrpcHandler, s *secret.SecretDescription) (*secret.Secret, error) {
				return h.Get(ctx, &secret.GetSecretRequest{
					Id:         s.Id,
					SecretType: s.SecretType,
				})
			},
		},
	}

	storage := func() *inmemory.Repository {
		r := inmemory.New()
		_ = r.Register(ctx, fixtures.ReferenceAccountInstance1)
		_ = r.Register(ctx, fixtures.ReferenceAccountInstance2)
		_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance1)
		_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance11)
		_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance2)
		return r
	}()

	fileStorage := fs.New(&config.ServerConfig{FileStorage: "/tmp/" + fixtures.PathFixture()})
	handler := New(storage, fileStorage)
	expectedAccountSecrets := map[string]map[string]*secret.SecretDescription{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx = metadata.NewIncomingContext(ctx,
				metadata.New(
					map[string]string{"account_id": test.Input.accountId}),
			)

			runResp, err := test.run(ctx, handler)
			assert.Equal(t, test.Expected.expectedRunErrorMsg, err)
			if err == nil {
				switch test.Expected.delta {
				case operationalDeltaINCREASED:
					if _, ok := expectedAccountSecrets[test.Input.accountId]; !ok {
						expectedAccountSecrets[test.Input.accountId] = map[string]*secret.SecretDescription{}
					}
					expectedAccountSecrets[test.Input.accountId][runResp.Id] = runResp
				case operationalDeltaDECREASED:
					if _, ok := expectedAccountSecrets[test.Input.accountId][runResp.Id]; ok {
						delete(expectedAccountSecrets[test.Input.accountId], runResp.Id)
					}
					if len(expectedAccountSecrets[test.Input.accountId]) == 0 {
						delete(expectedAccountSecrets, test.Input.accountId)
					}
				}

				assert.NotEqual(t, "", runResp.Id)
				assert.Equal(t, test.Expected.expectedSecretType, runResp.SecretType)
				assert.Equal(t, time.Now().Format(time.RFC3339), runResp.CreatedAt)
				checkResp, err := test.check(ctx, handler, runResp)
				assert.Equal(t, test.Expected.expectedCheckErrorMsg, err)
				if err == nil {
					assert.Equal(t, runResp.Id, checkResp.Id)
					assert.Equal(t, time.Now().Format(time.RFC3339), runResp.CreatedAt)
					switch test.Expected.expectedSecretType {
					case secret.SecretType_TEXT:
						assert.Equal(t, fixtures.ReferenceTextSecretValue1, checkResp.GetPlainText().GetData())
					case secret.SecretType_LOGIN_PASSWORD:
						assert.Equal(t, fixtures.ReferenceLoginPasswordSecretLogin1, checkResp.GetLoginPassword().GetLogin())
						assert.Equal(t, fixtures.ReferenceLoginPasswordSecretPassword1, checkResp.GetLoginPassword().GetPassword())
					case secret.SecretType_CREDIT_CARD:
						assert.Equal(t, fixtures.ReferenceCreditCardSecretCardNumber1, checkResp.GetCreditCard().GetNumber())
						assert.Equal(t, fixtures.ReferenceCreditCardSecretCardHolder1, checkResp.GetCreditCard().GetCardholderName())
						assert.Equal(t, fixtures.ReferenceCreditCardSecretValidTill1, checkResp.GetCreditCard().GetValidTill())
						assert.Equal(t, fixtures.ReferenceCreditCardSecretCvc11, checkResp.GetCreditCard().GetCvc())
					case secret.SecretType_FILE:
						assert.Equal(t, fixtures.ReferenceFileContent1, checkResp.GetFile().GetData())
					}

					listedAccountSecrets := map[string]*secret.SecretDescription{}
					listResp, err := handler.List(ctx, &secret.ListSecretRequest{})
					assert.Equal(t, nil, err)
					if err == nil {
						for _, desc := range listResp.Secrets {
							listedAccountSecrets[desc.Id] = desc
						}
						assert.Equal(t, expectedAccountSecrets[test.Input.accountId], listedAccountSecrets)
					}
				}
			}
		})
	}
	handler.fileStorage.CleanOut()
}
