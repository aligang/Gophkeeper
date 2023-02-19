package handler

import (
	"context"
	secret2 "github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/config"
	fixtures2 "github.com/aligang/Gophkeeper/pkg/server/fixtures"
	"github.com/aligang/Gophkeeper/pkg/server/repository/fs"
	"github.com/aligang/Gophkeeper/pkg/server/repository/inmemory"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

var ReferenceTextSecretId1 string
var ReferenceTextSecretId11 string
var ReferenceTextSecretId2 string

var ReferenceLoginPasswordSecretId1 string
var ReferenceLoginPasswordSecretId11 string
var ReferenceLoginPasswordSecretId2 string

var ReferenceCreditCardSecretId1 string
var ReferenceCreditCardSecretId11 string
var ReferenceCreditCardSecretId2 string

var ReferenceFileSecretId1 string
var ReferenceFileSecretId11 string
var ReferenceFileSecretId2 string

type handlerTestInput struct {
	accountId string
}

type handlerTestExpected struct {
	expectedRunResponse   *secret2.SecretDescription
	expectedRunErrorMsg   error
	expectedSecretType    secret2.SecretType
	secret                secret2.Secret
	expectedCheckErrorMsg error
	creationTime          string
	modificationTime      string
}

func TestWritableOperation(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		run      func(ctx context.Context, r *GrpcHandler) (*secret2.SecretDescription, error)
		Expected handlerTestExpected
		Input    handlerTestInput
	}{
		{
			name: "Create First Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_TEXT,
				secret: secret2.Secret{
					Secret: &secret2.Secret_PlainText{
						PlainText: &secret2.PlainText{
							Data: fixtures2.ReferenceTextSecretValue1,
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: "",
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				desc, err := h.Create(ctx, &secret2.CreateSecretRequest{
					Secret: &secret2.CreateSecretRequest_Text{
						Text: &secret2.PlainText{
							Data: fixtures2.ReferenceTextSecretValue1,
						},
					},
				})
				ReferenceTextSecretId1 = desc.Id
				return desc, err
			},
		},
		{
			name: "Create Second Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_TEXT,
				secret: secret2.Secret{
					Secret: &secret2.Secret_PlainText{
						PlainText: &secret2.PlainText{
							Data: fixtures2.ReferenceTextSecretValue11,
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: "",
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				desc, err := h.Create(ctx, &secret2.CreateSecretRequest{
					Secret: &secret2.CreateSecretRequest_Text{
						Text: &secret2.PlainText{
							Data: fixtures2.ReferenceTextSecretValue11,
						},
					},
				})
				ReferenceTextSecretId11 = desc.Id
				return desc, err
			},
		},
		{
			name: "Create Text Secret For another account",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId2,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_TEXT,
				secret: secret2.Secret{
					Secret: &secret2.Secret_PlainText{
						PlainText: &secret2.PlainText{
							Data: fixtures2.ReferenceTextSecretValue2,
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: "",
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				desc, err := h.Create(ctx, &secret2.CreateSecretRequest{
					Secret: &secret2.CreateSecretRequest_Text{
						Text: &secret2.PlainText{
							Data: fixtures2.ReferenceTextSecretValue2,
						},
					},
				})
				ReferenceTextSecretId2 = desc.Id
				return desc, err
			},
		},
		{
			name: "Update Existing Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_TEXT,
				secret: secret2.Secret{
					Secret: &secret2.Secret_PlainText{
						PlainText: &secret2.PlainText{
							Data: "UpdatedTextSecret",
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: time.Now().Format(time.RFC3339),
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				return h.Update(ctx, &secret2.UpdateSecretRequest{
					Secret: &secret2.UpdateSecretRequest_Text{
						Text: &secret2.PlainText{
							Data: "UpdatedTextSecret",
						},
					},
					Id: ReferenceTextSecretId1,
				})
			},
		},
		{
			name: "Delete Existing Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: status.Errorf(codes.NotFound, "Secret not found"),
				expectedSecretType:    secret2.SecretType_TEXT,
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				_, err := h.Delete(ctx, &secret2.DeleteSecretRequest{
					Id:         ReferenceTextSecretId1,
					SecretType: secret2.SecretType_TEXT,
				})
				return &secret2.SecretDescription{Id: ReferenceTextSecretId1}, err
			},
		},
		{
			name: "Update NonExisting Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   status.Errorf(codes.Unavailable, "Could not update secret"),
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_TEXT,
				secret: secret2.Secret{
					Secret: &secret2.Secret_PlainText{
						PlainText: &secret2.PlainText{
							Data: "UpdatedTextSecret",
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: time.Now().Format(time.RFC3339),
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				return h.Update(ctx, &secret2.UpdateSecretRequest{
					Secret: &secret2.UpdateSecretRequest_Text{
						Text: &secret2.PlainText{
							Data: "UpdatedTextSecret",
						},
					},
					Id: ReferenceTextSecretId1,
				})
			},
		},
		{
			name: "Delete Non-Existing Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   status.Errorf(codes.NotFound, "Could not delete secret"),
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_TEXT,
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				_, err := h.Delete(ctx, &secret2.DeleteSecretRequest{
					Id:         ReferenceTextSecretId1,
					SecretType: secret2.SecretType_TEXT,
				})
				return &secret2.SecretDescription{Id: ReferenceTextSecretId1}, err
			},
		},
		{
			name: "Create First File Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_FILE,
				secret: secret2.Secret{
					Secret: &secret2.Secret_File{
						File: &secret2.File{
							Data: fixtures2.ReferenceFileContent1,
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: "",
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				desc, err := h.Create(ctx, &secret2.CreateSecretRequest{
					Secret: &secret2.CreateSecretRequest_File{
						File: &secret2.File{
							Data: fixtures2.ReferenceFileContent1,
						},
					},
				})
				ReferenceFileSecretId1 = desc.Id
				return desc, err
			},
		},
		{
			name: "Create Second File Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_FILE,
				secret: secret2.Secret{
					Secret: &secret2.Secret_File{
						File: &secret2.File{
							Data: fixtures2.ReferenceFileContent11,
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: "",
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				desc, err := h.Create(ctx, &secret2.CreateSecretRequest{
					Secret: &secret2.CreateSecretRequest_File{
						File: &secret2.File{
							Data: fixtures2.ReferenceFileContent11,
						},
					},
				})
				ReferenceFileSecretId11 = desc.Id
				return desc, err
			},
		},
		{
			name: "Create File Secret For Another Account",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId2,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_FILE,
				secret: secret2.Secret{
					Secret: &secret2.Secret_File{
						File: &secret2.File{
							Data: fixtures2.ReferenceFileContent2,
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: "",
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				desc, err := h.Create(ctx, &secret2.CreateSecretRequest{
					Secret: &secret2.CreateSecretRequest_File{
						File: &secret2.File{
							Data: fixtures2.ReferenceFileContent2,
						},
					},
				})
				ReferenceFileSecretId2 = desc.Id
				return desc, err
			},
		},
		{
			name: "Update  File Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_FILE,
				secret: secret2.Secret{
					Secret: &secret2.Secret_File{
						File: &secret2.File{
							Data: []byte("NewlyUpdatedFileSecret"),
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: time.Now().Format(time.RFC3339),
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				return h.Update(ctx, &secret2.UpdateSecretRequest{
					Secret: &secret2.UpdateSecretRequest_File{
						File: &secret2.File{
							Data: []byte("NewlyUpdatedFileSecret"),
						},
					},
					Id: ReferenceFileSecretId1,
				})
			},
		},

		{
			name: "Delete Existing File Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   nil,
				expectedCheckErrorMsg: status.Errorf(codes.NotFound, "Secret not found"),
				expectedSecretType:    secret2.SecretType_FILE,
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				_, err := h.Delete(ctx, &secret2.DeleteSecretRequest{
					Id:         ReferenceFileSecretId1,
					SecretType: secret2.SecretType_FILE,
				})
				return &secret2.SecretDescription{Id: ReferenceFileSecretId1}, err
			},
		},
		{
			name: "Update NonExisting File Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   status.Errorf(codes.Unavailable, "Could not update secret"),
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_FILE,
				secret: secret2.Secret{
					Secret: &secret2.Secret_File{
						File: &secret2.File{
							Data: []byte("UpdatedTextSecret"),
						},
					},
				},
				creationTime:     time.Now().Format(time.RFC3339),
				modificationTime: time.Now().Format(time.RFC3339),
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				return h.Update(ctx, &secret2.UpdateSecretRequest{
					Secret: &secret2.UpdateSecretRequest_Text{
						Text: &secret2.PlainText{
							Data: "UpdatedTextSecret",
						},
					},
					Id: ReferenceTextSecretId1,
				})
			},
		},
		{
			name: "Delete Non-Existing Text Secret",
			Input: handlerTestInput{
				accountId: fixtures2.ReferenceAccountId1,
			},
			Expected: handlerTestExpected{
				expectedRunErrorMsg:   status.Errorf(codes.NotFound, "Could not delete secret"),
				expectedCheckErrorMsg: nil,
				expectedSecretType:    secret2.SecretType_FILE,
			},
			run: func(ctx context.Context, h *GrpcHandler) (*secret2.SecretDescription, error) {
				_, err := h.Delete(ctx, &secret2.DeleteSecretRequest{
					Id:         ReferenceTextSecretId1,
					SecretType: secret2.SecretType_FILE,
				})
				return &secret2.SecretDescription{Id: ReferenceFileSecretId1}, err
			},
		},
	}

	storage := func() *inmemory.Repository {
		r := inmemory.New()
		_ = r.Register(ctx, fixtures2.ReferenceAccountInstance1)
		_ = r.Register(ctx, fixtures2.ReferenceAccountInstance2)
		_ = r.AddToken(ctx, fixtures2.ReferenceTokenInstance1)
		_ = r.AddToken(ctx, fixtures2.ReferenceTokenInstance11)
		_ = r.AddToken(ctx, fixtures2.ReferenceTokenInstance2)
		return r
	}()

	fileStorage := fs.New(&config.ServerConfig{FileStorage: "/tmp/" + fixtures2.PathFixture()})
	handler := New(storage, fileStorage, &config.ServerConfig{SecretEncryptionEnabled: false})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx = metadata.NewIncomingContext(ctx,
				metadata.New(
					map[string]string{"account_id": test.Input.accountId}),
			)

			runResp, err := test.run(ctx, handler)
			assert.Equal(t, test.Expected.expectedRunErrorMsg, err)
			if err == nil {
				checkResp, err := handler.Get(ctx, &secret2.GetSecretRequest{
					Id:         runResp.Id,
					SecretType: test.Expected.expectedSecretType,
				})
				assert.Equal(t, test.Expected.expectedCheckErrorMsg, err)
				if err == nil {
					assert.Equal(t, test.Expected.creationTime, runResp.CreatedAt)
					assert.Equal(t, test.Expected.modificationTime, runResp.ModifiedAt)
					assert.Equal(t, test.Expected.expectedSecretType, runResp.SecretType)
					assert.Equal(t, time.Now().Format(time.RFC3339), runResp.CreatedAt)
					assert.Equal(t, runResp.Id, checkResp.Id)
					assert.Equal(t, test.Expected.creationTime, checkResp.CreatedAt)
					assert.Equal(t, test.Expected.modificationTime, checkResp.ModifiedAt)
					switch test.Expected.expectedSecretType {
					case secret2.SecretType_TEXT:
						assert.Equal(t, test.Expected.secret.GetPlainText().GetData(), checkResp.GetPlainText().GetData())
					case secret2.SecretType_LOGIN_PASSWORD:
						assert.Equal(t, fixtures2.ReferenceLoginPasswordSecretLogin1, checkResp.GetLoginPassword().GetLogin())
						assert.Equal(t, fixtures2.ReferenceLoginPasswordSecretPassword1, checkResp.GetLoginPassword().GetPassword())
					case secret2.SecretType_CREDIT_CARD:
						assert.Equal(t, fixtures2.ReferenceCreditCardSecretCardNumber1, checkResp.GetCreditCard().GetNumber())
						assert.Equal(t, fixtures2.ReferenceCreditCardSecretCardHolder1, checkResp.GetCreditCard().GetCardholderName())
						assert.Equal(t, fixtures2.ReferenceCreditCardSecretValidTill1, checkResp.GetCreditCard().GetValidTill())
						assert.Equal(t, fixtures2.ReferenceCreditCardSecretCvc11, checkResp.GetCreditCard().GetCvc())
					case secret2.SecretType_FILE:
						assert.Equal(t, test.Expected.secret.GetFile().GetData(), checkResp.GetFile().GetData())
					}
				}
			}
		})
	}
	handler.fileStorage.CleanOut()
}
