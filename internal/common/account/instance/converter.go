package instance

import (
	"github.com/aligang/Gophkeeper/internal/common/account"
	"time"
)

func ConvertAccountInstance(instance *Account) *account.Account {
	return &account.Account{
		Id:        instance.Id,
		CreatedAt: instance.CreatedAt.Format(time.RFC3339),
	}
}
