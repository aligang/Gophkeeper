package inmemory

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/server/repository/transaction"
)

func (r *Repository) WithinTransaction(
	ctx context.Context,
	fn func(context.Context, *transaction.DBTransaction) error) error {

	r.log.Debug("Starting transaction")
	r.Lock.Lock()
	err := fn(ctx, nil)
	if err != nil {
		r.log.Debug("Transaction Failed")
		r.Lock.Unlock()
		return err
	}
	r.log.Debug("Transaction Succeeded")
	r.Lock.Unlock()
	return nil
}
