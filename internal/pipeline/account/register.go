package account

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/account"
	"github.com/aligang/Gophkeeper/internal/pipeline"
	"os"
)

func Register(client account.AccountServiceClient, cli *pipeline.PipelineInitTree) {
	register := cli.Account.Register
	req := &account.RegisterRequest{Login: register.Login, Password: register.Password}
	response, err := client.Register(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s \n", response.Account.Id)
	fmt.Printf("%s \n", response.Account.Login)
	fmt.Printf("%s \n", response.Account.CreatedAt)
}
