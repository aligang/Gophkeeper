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
	"os"
	"testing"
	"time"
)

type handlerTestInput struct {
	storage     *inmemory.Repository
	fileStorage *fs.FileRepository
}

type handlerTestExpected struct {
	storage            *inmemory.Repository
	fileStorage        *fs.FileRepository
	expectedResponse   *secret.SecretDescription
	expectedErrorMsg   error
	expectedSecretType secret.SecretType
}

func TestCreditCardSecretWritableOperation(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                string
		handlerTestInput    handlerTestInput
		run                 func(r *GrpcHandler) (*secret.SecretDescription, error)
		handlerTestExpected handlerTestExpected
	}{
		{
			name: "Create Text Secret",
			handlerTestInput: handlerTestInput{
				storage: func() *inmemory.Repository {
					r := inmemory.New()
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance1)
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance2)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance1)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance11)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance2)
					return r
				}(),
				fileStorage: nil,
			},
			handlerTestExpected: handlerTestExpected{
				storage: func() *inmemory.Repository {
					r := inmemory.New()
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance1)
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance2)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance1)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance11)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance2)
					_ = r.AddTextSecret(ctx, fixtures.ReferenceTextSecretInstance1)
					return r
				}(),
				fileStorage:        nil,
				expectedResponse:   &secret.SecretDescription{},
				expectedErrorMsg:   nil,
				expectedSecretType: secret.SecretType_TEXT,
			},
			run: func(h *GrpcHandler) (*secret.SecretDescription, error) {
				ctx = metadata.NewIncomingContext(ctx,
					metadata.New(
						map[string]string{"account_id": fixtures.ReferenceAccountId1}),
				)
				return h.Create(ctx, &secret.CreateSecretRequest{
					Secret: &secret.CreateSecretRequest_Text{
						Text: &secret.PlainText{
							Data: fixtures.ReferenceTextSecretValue1,
						},
					},
				})
			},
		},
		{
			name: "Create File Secret",
			handlerTestInput: handlerTestInput{
				storage: func() *inmemory.Repository {
					r := inmemory.New()
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance1)
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance2)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance1)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance11)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance2)
					return r
				}(),
				fileStorage: fs.New(&config.ServerConfig{FileStorage: fixtures.PathFixture()}),
			},
			handlerTestExpected: handlerTestExpected{
				storage: func() *inmemory.Repository {
					r := inmemory.New()
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance1)
					_ = r.Register(ctx, fixtures.ReferenceAccountInstance2)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance1)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance11)
					_ = r.AddToken(ctx, fixtures.ReferenceTokenInstance2)
					_ = r.AddFileSecret(ctx, fixtures.ReferenceFileSecretInstance1)
					return r
				}(),
				fileStorage: func() *fs.FileRepository {
					repo := fs.New(&config.ServerConfig{FileStorage: fixtures.PathFixture()})
					repo.Save(ctx, fixtures.ReferenceObjectId1, fixtures.ReferenceFileContent1)
					return repo
				}(),
				expectedResponse:   &secret.SecretDescription{},
				expectedErrorMsg:   nil,
				expectedSecretType: secret.SecretType_FILE,
			},
			run: func(h *GrpcHandler) (*secret.SecretDescription, error) {
				ctx = metadata.NewIncomingContext(ctx,
					metadata.New(
						map[string]string{"account_id": fixtures.ReferenceAccountId1}),
				)
				return h.Create(ctx, &secret.CreateSecretRequest{
					Secret: &secret.CreateSecretRequest_File{
						File: &secret.File{
							Data: fixtures.ReferenceFileContent1,
						},
					},
				})
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := New(test.handlerTestInput.storage, test.handlerTestInput.fileStorage)
			resp, err := test.run(handler)
			assert.Equal(t, test.handlerTestExpected.expectedErrorMsg, err)
			assert.NotEqual(t, "", resp.Id)
			assert.Equal(t, test.handlerTestExpected.expectedSecretType, resp.SecretType)
			assert.Equal(t, time.Now().Format(time.RFC3339), resp.CreatedAt)
			//assert.Equal(t, true, handler.storage.(*inmemory.Repository).Equals(test.handlerTestExpected.storage))
			if test.handlerTestExpected.fileStorage != nil {
				assert.Equal(t, true, handler.fileStorage.Equals(test.handlerTestExpected.fileStorage))
				os.RemoveAll(test.handlerTestExpected.fileStorage.GetRootDirectory())
				os.RemoveAll(handler.fileStorage.GetRootDirectory())
			}
		})
	}
}
