package token

import "fmt"

func (t *Token) ToStdout() {
	fmt.Println(t.Id)
	fmt.Println(t.TokenValue)
	fmt.Println(t.IssuedAt)
}
