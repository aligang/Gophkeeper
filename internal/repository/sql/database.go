package sql

import (
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
	logging.Debug("Initializing SQL Repository")
	db, err := sqlx.Open("pgx", conf.DatabaseDsn)
	if err != nil {
		panic(err)
	}
	s := &Repository{
		DB:  db,
		log: logging.Logger.GetSubLogger("Repository", "sql"),
	}
	_, err = s.DB.Exec(
		"create table if not exists Accounts(Id text NOT NULL UNIQUE, Login text NOT NULL UNIQUE, Password text NOT NULL, Encryption_Key text, Create_dAt TIMESTAMP WITH TIME ZONE)",
	)
	if err != nil {
		s.log.Crit("Failure during initialisation of accounts table: %s", err.Error())
		panic(err.Error())
	}
	_, err = s.DB.Exec(
		"create table if not exists Tokens(Id text NOT NULL UNIQUE, Token_Value text NOT NULL UNIQUE, Owner text NOT NULL, Issued_At TIMESTAMP WITH TIME ZONE NOT NULL)",
	)
	if err != nil {
		s.log.Crit("Failure during initialisation of tokens table: %s", err.Error())
		panic(err.Error())
	}
	_, err = s.DB.Exec(
		"create table if not exists Text_Secrets(Id text NOT NULL UNIQUE, Data text, Created_At TIMESTAMP WITH TIME ZONE NOT NULL, Modified_At TIMESTAMP WITH TIME ZONE)",
	)
	if err != nil {
		s.log.Crit("Failure during initialisation of text secrets table: %s", err.Error())
		panic(err.Error())
	}
	_, err = s.DB.Exec(
		"create table if not exists Login_Password_Secrets(Id text NOT NULL UNIQUE, Login text, Password text,  Created_At TIMESTAMP WITH TIME ZONE NOT NULL, Modified_At TIMESTAMP WITH TIME ZONE)",
	)
	if err != nil {
		s.log.Crit("Failure during initialisation of login password secrets table: %s", err.Error())
		panic(err.Error())
	}
	_, err = s.DB.Exec(
		"create table if not exists Credit_Card_Secrets(Id text NOT NULL UNIQUE, Card_Number text, Card_Holder text, Valid_Till text, Cvc text, Created_At TIMESTAMP WITH TIME ZONE NOT NULL, Modified_At TIMESTAMP WITH TIME ZONE)",
	)
	if err != nil {
		s.log.Debug("Failure during initialisation of credit card secrets table: %s", err.Error())
		panic(err.Error())
	}
	_, err = s.DB.Exec(
		"create table if not exists File_Secrets(Id text NOT NULL UNIQUE, Object_ID text NOT NULL UNIQUE, Created_At TIMESTAMP WITH TIME ZONE NOT NULL, Modified_At TIMESTAMP WITH TIME ZONE)",
	)
	if err != nil {
		s.log.Debug("Failure during initialisation of file secrets table: %s", err.Error())
		panic(err.Error())
	}
	_, err = s.DB.Exec(
		"create table if not exists File_Deletion_Queue(ObjectId text NOT NULL UNIQUE, Deleted_At TIMESTAMP WITH TIME ZONE NOT NULL)",
	)
	if err != nil {
		s.log.Debug("Failure during initialisation of deletion queue table: %s", err.Error())
		panic(err.Error())
	}

	s.log.Debug(" SQL Repository initialisation succeeded")
	return s
}
