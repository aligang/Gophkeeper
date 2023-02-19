package dispatcher

import (
	"fmt"
	acccountProto "github.com/aligang/Gophkeeper/pkg/account"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline/account"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline/secret"
	loginpassword2 "github.com/aligang/Gophkeeper/pkg/client/pipeline/secret/creditcard"
	file2 "github.com/aligang/Gophkeeper/pkg/client/pipeline/secret/file"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline/secret/loginpassword"
	text2 "github.com/aligang/Gophkeeper/pkg/client/pipeline/secret/text"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline/token"
	"github.com/aligang/Gophkeeper/pkg/client/pipeline/version"
	"github.com/aligang/Gophkeeper/pkg/config"
	"github.com/aligang/Gophkeeper/pkg/logging"
	secretProto "github.com/aligang/Gophkeeper/pkg/secret"
	tokenGetter "github.com/aligang/Gophkeeper/pkg/token/tokengetter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func Start(cfg *config.ClientConfig, pipelineTree *pipeline.PipelineInitTree) {
	if pipelineTree.BuildInfo != nil {
		version.Print()
		os.Exit(0)
	}

	logging.Debug("Connecting to %s", cfg.ServerAddress)
	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	logging.Debug("Connection succeeded")

	tg := tokenGetter.New(conn, cfg)
	secretClient := secretProto.NewSecretServiceClient(conn)
	accountClient := acccountProto.NewAccountServiceClient(conn)

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
				text2.Create(secretClient, tg, pipelineTree)
			case t.Update != nil:
				text2.Update(secretClient, tg, pipelineTree)
			case t.Delete != nil:
				text2.Delete(secretClient, tg, pipelineTree)
			case t.Get != nil:
				text2.Get(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret Text pipeline initialization error:\n")
			}
		case secretPipelineTree.LoginPassword != nil:
			l := secretPipelineTree.LoginPassword
			switch {
			case l.Create != nil:
				loginpassword.Create(secretClient, tg, pipelineTree)
			case l.Update != nil:
				loginpassword.Update(secretClient, tg, pipelineTree)
			case l.Delete != nil:
				loginpassword.Delete(secretClient, tg, pipelineTree)
			case l.Get != nil:
				loginpassword.Get(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret Login-Password pipeline initialization error:\n")
			}
		case secretPipelineTree.CreditCard != nil:
			cc := secretPipelineTree.CreditCard
			switch {
			case cc.Create != nil:
				loginpassword2.Create(secretClient, tg, pipelineTree)
			case cc.Update != nil:
				loginpassword2.Update(secretClient, tg, pipelineTree)
			case cc.Delete != nil:
				loginpassword2.Delete(secretClient, tg, pipelineTree)
			case cc.Get != nil:
				loginpassword2.Get(secretClient, tg, pipelineTree)
			default:
				fmt.Fprintf(os.Stderr, "Secret CreditCard pipeline initialization error:\n")
			}
		case secretPipelineTree.File != nil:
			f := secretPipelineTree.File
			switch {
			case f.Upload != nil:
				file2.Upload(secretClient, tg, pipelineTree)
			case f.Update != nil:
				file2.Update(secretClient, tg, pipelineTree)
			case f.Delete != nil:
				file2.Delete(secretClient, tg, pipelineTree)
			case f.Download != nil:
				file2.Download(secretClient, tg, pipelineTree)
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
