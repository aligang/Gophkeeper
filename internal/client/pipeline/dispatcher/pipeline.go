package dispatcher

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/client/config"
	"github.com/aligang/Gophkeeper/internal/client/pipeline"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/account"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/secret"
	creditCard "github.com/aligang/Gophkeeper/internal/client/pipeline/secret/creditcard"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/secret/file"
	loginPassword "github.com/aligang/Gophkeeper/internal/client/pipeline/secret/loginpassword"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/secret/text"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/token"
	"github.com/aligang/Gophkeeper/internal/client/pipeline/version"
	tokenGetter "github.com/aligang/Gophkeeper/internal/client/token/tokengetter"
	accountProto "github.com/aligang/Gophkeeper/internal/common/account"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	secretProto "github.com/aligang/Gophkeeper/internal/common/secret"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func RunPipeline(cfg *config.Config, pipelineTree *pipeline.PipelineInitTree) {
	if pipelineTree.BuildInfo != nil {
		version.Print()
		os.Exit(0)
	}

	dialOpts := []grpc.DialOption{}
	switch {
	case cfg.EnableTlsEncryption && cfg.CaCertPath != "":
		b, err := os.ReadFile(cfg.CaCertPath)
		if err != nil {
			logging.Fatal("Could not read  CA certificate file: %s", err.Error())
		}
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			logging.Fatal("Could not add CA certificate to certificate pool")
		}

		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            cp,
		})))
	case cfg.EnableTlsEncryption:
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})))
	default:
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	logging.Debug("Connecting to %s", cfg.ServerAddress)
	conn, err := grpc.Dial(cfg.ServerAddress, dialOpts...)
	if err != nil {
		logging.Crit(err.Error())
	}
	defer conn.Close()
	logging.Debug("Connection succeeded")

	tg := tokenGetter.New(conn, cfg)
	secretClient := secretProto.NewSecretServiceClient(conn)
	accountClient := accountProto.NewAccountServiceClient(conn)

	switch {
	case pipelineTree.BuildInfo != nil:
		version.Print()
	case pipelineTree.Token != nil:
		tokenPipelineTree := pipelineTree.Token
		switch {
		case tokenPipelineTree.Get != nil:
			token.Get(tg, pipelineTree)
		default:
			fmt.Fprintf(os.Stderr, "Token pipeline initialization error:\n")
		}
	case pipelineTree.Account != nil:
		acc := pipelineTree.Account
		switch {
		case acc.Register != nil:
			account.Register(accountClient, pipelineTree)
		default:
			fmt.Fprintf(os.Stderr, "Account pipeline initialization error:\n")
			os.Exit(1)
		}
	case pipelineTree.Secret != nil:
		secretPipelineTree := pipelineTree.Secret
		switch {
		case secretPipelineTree.Text != nil:
			t := secretPipelineTree.Text
			switch {
			case t.Create != nil:
				text.Create(secretClient, tg, pipelineTree)
			case t.Update != nil:
				text.Update(secretClient, tg, pipelineTree)
			case t.Delete != nil:
				text.Delete(secretClient, tg, pipelineTree)
			case t.Get != nil:
				text.Get(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret Text pipeline initialization error:\n")
			}
		case secretPipelineTree.LoginPassword != nil:
			l := secretPipelineTree.LoginPassword
			switch {
			case l.Create != nil:
				loginPassword.Create(secretClient, tg, pipelineTree)
			case l.Update != nil:
				loginPassword.Update(secretClient, tg, pipelineTree)
			case l.Delete != nil:
				loginPassword.Delete(secretClient, tg, pipelineTree)
			case l.Get != nil:
				loginPassword.Get(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret Login-Password pipeline initialization error:\n")
			}
		case secretPipelineTree.CreditCard != nil:
			cc := secretPipelineTree.CreditCard
			switch {
			case cc.Create != nil:
				creditCard.Create(secretClient, tg, pipelineTree)
			case cc.Update != nil:
				creditCard.Update(secretClient, tg, pipelineTree)
			case cc.Delete != nil:
				creditCard.Delete(secretClient, tg, pipelineTree)
			case cc.Get != nil:
				creditCard.Get(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret CreditCard pipeline initialization error:\n")
			}
		case secretPipelineTree.File != nil:
			f := secretPipelineTree.File
			switch {
			case f.Upload != nil:
				file.Upload(secretClient, tg, pipelineTree)
			case f.Update != nil:
				file.Update(secretClient, tg, pipelineTree)
			case f.Delete != nil:
				file.Delete(secretClient, tg, pipelineTree)
			case f.Download != nil:
				file.Download(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret CreditCard pipeline initialization error:\n")
			}
		case secretPipelineTree.List != nil:
			secret.List(secretClient, tg, pipelineTree)

		default:
			fmt.Fprintf(os.Stderr, "Secret pipeline initialization error:\n")
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Pipeline initialization error:\n")
		os.Exit(1)
	}
}
