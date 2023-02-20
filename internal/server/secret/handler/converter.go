package handler

import (
	"github.com/aligang/Gophkeeper/internal/common/secret"
	secretInstance "github.com/aligang/Gophkeeper/internal/common/secret/instance"
	"time"
)

func convertTextSecretInstance(instance *secretInstance.TextSecret) *secret.Secret {
	s := &secret.Secret{
		Id:        instance.Id,
		CreatedAt: instance.CreatedAt.Format(time.RFC3339),
		AccountID: instance.AccountId,
		Secret: &secret.Secret_PlainText{
			PlainText: &secret.PlainText{
				Data: instance.Text,
			},
		},
	}
	if !instance.ModifiedAt.IsZero() {
		s.ModifiedAt = instance.ModifiedAt.Format(time.RFC3339)
	}
	return s
}

func convertLoginPasswordSecretInstance(instance *secretInstance.LoginPasswordSecret) *secret.Secret {
	s := &secret.Secret{
		Id:        instance.Id,
		CreatedAt: instance.CreatedAt.Format(time.RFC3339),
		AccountID: instance.AccountId,
		Secret: &secret.Secret_LoginPassword{
			LoginPassword: &secret.LoginPassword{
				Login:    instance.Login,
				Password: instance.Password,
			},
		},
	}
	if !instance.ModifiedAt.IsZero() {
		s.ModifiedAt = instance.ModifiedAt.Format(time.RFC3339)
	}
	return s
}

func convertCreditCardSecretInstance(instance *secretInstance.CreditCardSecret) *secret.Secret {
	s := &secret.Secret{
		Id:        instance.Id,
		CreatedAt: instance.CreatedAt.Format(time.RFC3339),
		AccountID: instance.AccountId,
		Secret: &secret.Secret_CreditCard{
			CreditCard: &secret.CreditCard{
				Number:         instance.CardNumber,
				CardholderName: instance.CardHolder,
				ValidTill:      instance.ValidTill,
				Cvc:            instance.Cvc,
			},
		},
	}
	if !instance.ModifiedAt.IsZero() {
		s.ModifiedAt = instance.ModifiedAt.Format(time.RFC3339)
	}
	return s
}

func convertFileSecretInstance(instance *secretInstance.FileSecret) *secret.Secret {
	s := &secret.Secret{
		Id:        instance.Id,
		CreatedAt: instance.CreatedAt.Format(time.RFC3339),
		AccountID: instance.AccountId,
		Secret: &secret.Secret_File{
			File: &secret.File{},
		},
	}
	if !instance.ModifiedAt.IsZero() {
		s.ModifiedAt = instance.ModifiedAt.Format(time.RFC3339)
	}
	return s
}

func convertTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)

}
