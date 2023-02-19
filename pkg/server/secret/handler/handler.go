package handler

import (
	"github.com/aligang/Gophkeeper/pkg/common/secret"
	"github.com/aligang/Gophkeeper/pkg/config"
	"github.com/aligang/Gophkeeper/pkg/server/repository"
	"github.com/aligang/Gophkeeper/pkg/server/repository/fs"
)

type GrpcHandler struct {
	secret.UnimplementedSecretServiceServer
	storage     repository.Storage
	fileStorage *fs.FileRepository
	cfg         *config.ServerConfig
}

func New(storage repository.Storage, fileStorage *fs.FileRepository, cfg *config.ServerConfig) *GrpcHandler {
	return &GrpcHandler{
		storage:     storage,
		fileStorage: fileStorage,
		cfg:         cfg,
	}
}
