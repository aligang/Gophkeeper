package tokengetter

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/account"
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/logging"
	"google.golang.org/grpc"
	"log"
)

type TokenGetter struct {
	StaticToken        string
	AuthServiceAddress string
	Login              string
	Password           string
	GetToken           func() string
	Client             account.AccountServiceClient
	logger             *logging.InternalLogger
}

func (t *TokenGetter) getStaticToken() string {
	logger := t.logger.GetSubLogger("Static", "Token")
	logger.Debug("Using predefined token")
	return t.StaticToken
}

func (t *TokenGetter) getDynamicToken() string {
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
	return resp.Token.Value
}

func New(conn grpc.ClientConnInterface, cfg *config.ClientConfig) *TokenGetter {
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
