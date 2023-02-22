package instance

import "time"

type BaseSecret struct {
	Id         string
	AccountId  string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type TextSecret struct {
	BaseSecret
	Text string
}

type LoginPasswordSecret struct {
	BaseSecret
	Login    string
	Password string
}

type CreditCardSecret struct {
	BaseSecret
	CardNumber string
	CardHolder string
	ValidTill  string
	Cvc        string
}

type FileSecret struct {
	BaseSecret
	ObjectId string
}
