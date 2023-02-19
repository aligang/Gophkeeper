package fsgc

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/aligang/Gophkeeper/pkg/server/config"
	"github.com/aligang/Gophkeeper/pkg/server/repository"
	"github.com/aligang/Gophkeeper/pkg/server/repository/fs"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"time"
)

type FileSystemGB struct {
	conf        *config.Config
	storage     repository.Storage
	fileStorage *fs.FileRepository
	logger      *logging.InternalLogger
}

func New(
	conf *config.Config, storage repository.Storage,
	fileStorage *fs.FileRepository) *FileSystemGB {
	return &FileSystemGB{
		conf:        conf,
		storage:     storage,
		fileStorage: fileStorage,
		logger:      logging.Logger.GetSubLogger("GarbageCollector", "Filesystem"),
	}
}

func (gb *FileSystemGB) CleanStale(ctx context.Context) {
	gb.logger.Debug("starting filesystem clean sequence")
	err := gb.storage.WithinTransaction(
		ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			gb.logger.Debug("listing deletion queue")
			deletionQueue, terr := gb.storage.ListFileDeletionQueue(ctx, tx)
			if terr != nil {
				return terr
			}
			if len(deletionQueue) == 0 {
				gb.logger.Debug("Deletion queue is empty ")
				return nil
			}
			counter := 0
			for _, e := range deletionQueue {
				if e.DeletedAt.Add(time.Minute * time.Duration(gb.conf.GetFileStaleTimeMinutes())).Before(time.Now()) {
					gb.logger.Debug("Deleting object %s", e.ObjectId)
					terr = gb.fileStorage.Delete(ctx, e.ObjectId)
					if terr != nil {
						return terr
					}
					terr = gb.storage.DeleteFileSecretFromDeletionQueue(ctx, e.ObjectId, tx)
					if terr != nil {
						return terr
					}
					counter = counter + 1
				}
			}
			gb.logger.Debug("Elements from deletion  were deleted: %d", counter)
			return nil
		},
	)
	if err != nil {
		gb.logger.Debug("Error during deletion: %s", err.Error())
	}
	gb.logger.Debug("filesystem clean sequence is finished")
}

func (gb *FileSystemGB) Run(ctx context.Context) {
	gb.logger.Debug("Instantiating FileSystemGC")
	ticker := time.NewTicker(time.Minute * time.Duration(gb.conf.FileStaleTimeMinutes))
	for {
		select {
		case <-ticker.C:
			gb.CleanStale(ctx)
		case <-ctx.Done():
			gb.logger.Debug("Received stop signal")
			return
		}
	}
}
