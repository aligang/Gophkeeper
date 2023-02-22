package transaction

import "github.com/jmoiron/sqlx"

type DBTransaction struct {
	Sql *sqlx.Tx
	//may be extended to support another transaction mechanisms
}
