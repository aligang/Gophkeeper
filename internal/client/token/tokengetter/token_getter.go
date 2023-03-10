package tokengetter

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/client/config"
	"github.com/aligang/Gophkeeper/internal/common/account"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"github.com/aligang/Gophkeeper/internal/common/token"
	"google.golang.org/grpc"
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
	logger.Info("Using predefined token")
	return &token.Token{
		Id:         "Statically Defined Token",
		TokenValue: t.StaticToken,
		IssuedAt:   "N/A",
	}

}

func (t *TokenGetter) getDynamicToken() *token.Token {
	logger := t.logger.GetSubLogger("Dynamic", "Token")
	logger.Info("Fetching token from AuthServer")
	resp, err := t.Client.Authenticate(
		context.Background(),
		&account.AuthenticationRequest{Login: t.Login, Password: t.Password})
	if err != nil {
		logger.Fatal("Could not fetch token from AuthServer: %s", err.Error())
	}
	logger.Info("Using token received from Auth Service")
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
