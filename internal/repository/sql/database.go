package sql

import (
	"fmt"
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/jmoiron/sqlx"
)
import _ "github.com/jackc/pgx/v4/stdlib"

type Repository struct {
	DB  *sqlx.DB
	log *logging.InternalLogger
}

func New(conf *config.ServerConfig) *Repository {
	logging.Debug("Initialisating SQL Repository")
	db, err := sqlx.Open("pgx", conf.DatabaseDsn)
	if err != nil {
		panic(err)
	}
	s := &Repository{
		DB:  db,
		log: logging.Logger.GetSubLogger("filerepository", "sql"),
	}
	_, err = s.DB.Exec(
		"create table if not exists accounts(Login text NOT NULL UNIQUE, Password text NOT NULL, Current double precision, Withdraw double precision)",
	)
	if err != nil {
		msg, _ := fmt.Printf("Failure during initialisation of accounts table: %s\n", err.Error())
		panic(msg)
	}
	_, err = s.DB.Exec(
		"create table if not exists orders(Number bigint NOT NULL UNIQUE, Status text NULL UNIQUE, Accural double precision, UploadedAt TIMESTAMP WITH TIME ZONE, Owner text)",
	)
	if err != nil {
		msg, _ := fmt.Printf("Failure during initialisation of orders table: %s\n", err.Error())
		panic(msg)
	}
	_, err = s.DB.Exec(
		"create table if not exists pending_orders(order_id text NOT NULL UNIQUE)",
	)
	if err != nil {
		msg, _ := fmt.Printf("Failure during initialisation of pending orders table: %s\n", err.Error())
		panic(msg)
	}
	_, err = s.DB.Exec(
		"create table if not exists withdrawns(OrderID bigint NOT NULL UNIQUE, Sum double precision, ProcessedAt TIMESTAMP WITH TIME ZONE, owner text)",
	)
	if err != nil {
		msg, _ := fmt.Printf("Failure during initialisation of withdrawns table: %s\n", err.Error())
		panic(msg)
	}
	logging.Debug(" SQL Repository initialisation successeeded")
	return s
}
