package tokengc

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/config"
	"github.com/aligang/Gophkeeper/pkg/logging"
	"github.com/aligang/Gophkeeper/pkg/server/repository"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
	"time"
)

type TokenGC struct {
	conf    *config.ServerConfig
	storage repository.Storage
	logger  *logging.InternalLogger
}

func New(conf *config.ServerConfig, storage repository.Storage) *TokenGC {
	logger := logging.Logger.GetSubLogger("GarbageCollector", "Token")
	logging.Debug("Instantiating Token Garbage Collector")
	t := &TokenGC{
		conf:    conf,
		storage: storage,
		logger:  logger,
	}
	logging.Debug("Token Garbage Collector successfully Instantiated")
	return t
}

func (gc *TokenGC) CleanStale(ctx context.Context) {
	gc.logger.Debug("starting token clean sequence")
	deletedTokenCounter := 0
	err := gc.storage.WithinTransaction(
		ctx, func(tCtx context.Context, tx *transaction.DBTransaction) error {
			gc.logger.Debug("Listing current tokens...")
			tokens, terr := gc.storage.ListTokens(tCtx, tx)
			if terr != nil {
				return terr
			}
			gc.logger.Debug("Token records found: %d", len(tokens))
			for _, token := range tokens {
				logging.Debug("Checking token issued at: %s", token.IssuedAt.Format(time.RFC3339))
				if token.IssuedAt.Add(time.Minute * time.Duration(gc.conf.TokenValidityTimeMinutes)).Before(time.Now()) {
					terr = gc.storage.DeleteToken(tCtx, token, tx)
					if terr != nil {
						gc.logger.Debug("Could not delete token %s: %s", token.TokenValue, terr.Error())
					}
					deletedTokenCounter = deletedTokenCounter + 1
				}
			}
			return nil
		},
	)
	if err != nil {
		gc.logger.Debug("Error during deletion: %s", err.Error())
	}
	if deletedTokenCounter != 0 {
		gc.logger.Debug("Token were deleted: %d", deletedTokenCounter)
	} else {
		gc.logger.Debug("no tokens were deleted")
	}
	gc.logger.Debug("token clean sequence is finished")
}

func (gc *TokenGC) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * time.Duration(gc.conf.TokenValidityTimeMinutes))
	for {
		select {
		case <-ticker.C:
			gc.CleanStale(ctx)
		case <-ctx.Done():
			gc.logger.Debug("Received stop signal")
			return
		}
	}
}
