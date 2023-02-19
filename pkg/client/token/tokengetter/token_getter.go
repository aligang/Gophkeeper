package tokengetter

import (
	"context"
	"github.com/aligang/Gophkeeper/pkg/client/config"
	"github.com/aligang/Gophkeeper/pkg/common/account"
	"github.com/aligang/Gophkeeper/pkg/common/logging"
	"github.com/aligang/Gophkeeper/pkg/common/token"
	"google.golang.org/grpc"
	"log"
)

type TokenGetter struct {
	StaticToken        string
	AuthServiceAddress string
	Login              string
	Password           string
	GetToken           func() *token.Token
	Client             account.AccountServiceClient
	logger             *logging.InternalLogger
}

func (t *TokenGetter) getStaticToken() *token.Token {
	logger := t.logger.GetSubLogger("Static", "Token")
	logger.Debug("Using predefined token")
	return &token.Token{
		Id:         "Statically Defined Token",
		TokenValue: t.StaticToken,
		IssuedAt:   "N/A",
	}

}

func (t *TokenGetter) getDynamicToken() *token.Token {
	logger := t.logger.GetSubLogger("Dynamic", "Token")
	logger.Debug("Fetching token from AuthServer")
	resp, err := t.Client.Authenticate(
		context.Background(),
		&account.AuthenticationRequest{Login: t.Login, Password: t.Password})
	if err != nil {
		logger.Debug("Could not fetch token from AuthServer")
		log.Fatal(err.Error())
	}
	logger.Debug("Using token received from Auth Service")
	logger.Debug("token id %s, issued at %s", resp.Token.Id, resp.Token.IssuedAt)
	return resp.Token
}

func New(conn grpc.ClientConnInterface, cfg *config.Config) *TokenGetter {
	var t *TokenGetter
	if cfg.StaticToken != "" {
		t = &TokenGetter{
			StaticToken: cfg.StaticToken,
			logger:      logging.Logger.GetSubLogger("Static", "Token"),
		}
		t.GetToken = t.getStaticToken
		return t
	}

	t = &TokenGetter{
		AuthServiceAddress: cfg.ServerAddress,
		Login:              cfg.Login,
		Password:           cfg.Password,
		Client:             account.NewAccountServiceClient(conn),
		logger:             logging.Logger.GetSubLogger("Dynamic", "Token"),
	}
	t.GetToken = t.getDynamicToken
	return t
}
