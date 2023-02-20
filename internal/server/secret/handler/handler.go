package handler

import (
	"github.com/aligang/Gophkeeper/internal/common/secret"
	"github.com/aligang/Gophkeeper/internal/server/config"
	"github.com/aligang/Gophkeeper/internal/server/repository"
	"github.com/aligang/Gophkeeper/internal/server/repository/fs"
)

type GrpcHandler struct {
	secret.UnimplementedSecretServiceServer
	storage     repository.Storage
	fileStorage *fs.FileRepository
	cfg         *config.Config
}

func New(storage repository.Storage, fileStorage *fs.FileRepository, cfg *config.Config) *GrpcHandler {
	return &GrpcHandler{
		storage:     storage,
		fileStorage: fileStorage,
		cfg:         cfg,
	}
}

func (h *GrpcHandler) isSecretEncryptionEnabled() bool {
	return h.cfg.SecretEncryptionEnabled
}
