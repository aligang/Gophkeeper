package handler

import (
	"github.com/aligang/Gophkeeper/internal/repository"
	"github.com/aligang/Gophkeeper/internal/repository/fs"
	"github.com/aligang/Gophkeeper/internal/secret"
)

type GrpcHandler struct {
	secret.UnimplementedSecretServiceServer
	storage     repository.Storage
	fileStorage *fs.FileRepository
}

func New(storage repository.Storage, fileStorage *fs.FileRepository) *GrpcHandler {
	return &GrpcHandler{
		storage:     storage,
		fileStorage: fileStorage,
	}
}
