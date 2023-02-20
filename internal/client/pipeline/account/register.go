package account

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	account2 "github.com/aligang/Gophkeeper/internal/common/account"
	"os"
)

func Register(client account2.AccountServiceClient, cli *pipeline.PipelineInitTree) {
	register := cli.Account.Register
	req := &account2.RegisterRequest{Login: register.Login, Password: register.Password}
	response, err := client.Register(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("account %s created\n", response.Account.Id)
	fmt.Printf("creation date: %s\n", response.Account.CreatedAt)
}
