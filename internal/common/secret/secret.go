package secret

import (
	"fmt"
	"time"
)

func (s *Secret) ToStdout() {
	fmt.Println("  {")
	fmt.Printf("    id          : %s\n", s.Id)
	fmt.Printf("    created at  : %s\n", s.CreatedAt)
	if t, err := time.Parse(time.RFC3339, s.ModifiedAt); err == nil && !t.IsZero() {
		fmt.Printf("    modified at : %s\n", s.ModifiedAt)
	}
	switch s.Secret.(type) {
	case *Secret_PlainText:
		data := s.Secret.(*Secret_PlainText).PlainText.Data
		fmt.Printf("    text        : %s\n", data)
	case *Secret_LoginPassword:
		data := s.Secret.(*Secret_LoginPassword).LoginPassword
		fmt.Printf("    login       : %s\n", data.Login)
		fmt.Printf("    password    : %s\n", data.Password)
	case *Secret_CreditCard:
		data := s.Secret.(*Secret_CreditCard).CreditCard
		fmt.Printf("    number      : %s\n", data.Number)
		fmt.Printf("    cardholder  : %s\n", data.CardholderName)
		fmt.Printf("    valid till  : %s\n", data.ValidTill)
		fmt.Printf("    cvc         : %s\n", data.Cvc)
	default:
		fmt.Println("Secret is not considered as valid target for stdout")
	}
	fmt.Println("  }")
}
