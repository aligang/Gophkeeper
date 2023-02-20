package instance

import (
	"github.com/aligang/Gophkeeper/internal/common/token"
	"time"
)

func ConvertTokenInstance(instance *Token) *token.Token {
	return &token.Token{
		Id:         instance.Id,
		IssuedAt:   instance.IssuedAt.Format(time.RFC3339),
		TokenValue: instance.TokenValue,
	}
}
