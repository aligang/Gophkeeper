package pipeline

import (
	"flag"
	"fmt"
	"os"
)

func GetPipeline(target *PipelineInitTree, inline func()) *PipelineInitTree {
	flag.Usage = func() {
		inline()
		fmt.Fprintf(os.Stderr, "sub-command:\n")
		fmt.Fprintf(os.Stderr, "account|secret|version|token'.\n")
	}

	flag.Parse()
	args := flag.Args()

	//target := &PipelineInitTree{}
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	subcommand := args[0]
	args = args[1:]
	switch subcommand {
	case "version":
		target.BuildInfo = &PipelineInitTree_BuildInfo{}
		subFlags := flag.NewFlagSet("client version", flag.ExitOnError)
		subFlags.Parse(args)
	case "token":
		target.Token = &PipelineInitTree_Token{}
		subFlags := flag.NewFlagSet("client token", flag.ExitOnError)
		subFlags.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage: ./client token \n")
			fmt.Fprintf(os.Stderr, "sub-command:\n")
			fmt.Fprintf(os.Stderr, "    get\n")
		}
		if len(args) < 1 {
			flag.Usage()
			os.Exit(1)
		}
		subcommand = args[0]
		args = args[1:]
		switch subcommand {
		case "get":
			target.Token.Get = &PipelineInitTree_Token_Get{}
			subFlags = flag.NewFlagSet("client token get", flag.ExitOnError)
			subFlags.Parse(args)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported command: %s\n ", subcommand)
			subFlags.Usage()
			os.Exit(1)
		}
		subFlags.Parse(args)
	case "account":
		target.Account = &PipelineInitTree_Account{}
		subFlags := flag.NewFlagSet("\account", flag.ExitOnError)
		subFlags.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage: ./client account \n")
			fmt.Fprintf(os.Stderr, "sub-command:\n")
			fmt.Fprintf(os.Stderr, "    register\n")
		}
		if len(args) < 1 {
			flag.Usage()
			os.Exit(1)
		}
		subcommand = args[0]
		args = args[1:]
		switch subcommand {
		case "register":
			target.Account.Register = &PipelineInitTree_Account_Register{}
			subFlags = flag.NewFlagSet("client account register", flag.ExitOnError)
			subFlags.StringVar(&target.Account.Register.Login, "login", "", "login for new account")
			subFlags.StringVar(
				&target.Account.Register.Password, "password", "", "password for new account")
			if len(args) < 2 {
				subFlags.Usage()
				os.Exit(1)
			}
			subFlags.Parse(args)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported command: %s\n ", subcommand)
			subFlags.Usage()
			os.Exit(1)
		}
		subFlags.Parse(args)
	case "secret":
		target.Secret = &PipelineInitTree_Secret{}
		subFlags := flag.NewFlagSet("secret", flag.ExitOnError)
		subFlags.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage: ./client secret \n")
			fmt.Fprintf(os.Stderr, "sub-command:\n")
			fmt.Fprintf(os.Stderr, "    text\n")
			fmt.Fprintf(os.Stderr, "    login-password\n")
			fmt.Fprintf(os.Stderr, "    credit-card\n")
			fmt.Fprintf(os.Stderr, "    file\n")
		}
		subFlags.Parse(args)
		if len(args) < 1 {
			subFlags.Usage()
			os.Exit(1)
		}
		subcommand = args[0]
		args = args[1:]
		switch subcommand {
		case "text":
			target.Secret.Text = &PipelineInitTree_Secret_Text{}
			subFlags := flag.NewFlagSet("text", flag.ExitOnError)
			subFlags.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: ./client secret text\n")
				fmt.Fprintf(os.Stderr, "sub-command:\n")
				fmt.Fprintf(os.Stderr, "    get\n")
				fmt.Fprintf(os.Stderr, "    create\n")
				fmt.Fprintf(os.Stderr, "    delete\n")
				fmt.Fprintf(os.Stderr, "    update\n")
			}
			subFlags.Parse(args)
			if len(args) < 1 {
				subFlags.Usage()
				os.Exit(1)
			}
			subcommand = args[0]
			args = args[1:]
			switch subcommand {
			case "get":
				target.Secret.Text.Get = &PipelineInitTree_Secret_Text_Get{}
				subFlags = flag.NewFlagSet("client secret text get", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.Text.Get.Id, "id", "", "text secret id")
				subFlags.Parse(args)
				if target.Secret.Text.Get.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "create":
				target.Secret.Text.Create = &PipelineInitTree_Secret_Text_Create{}
				subFlags = flag.NewFlagSet("client secret text create", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.Text.Create.Data, "data", "", "text secret data")
				subFlags.Parse(args)
				if target.Secret.Text.Create.Data == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "delete":
				target.Secret.Text.Delete = &PipelineInitTree_Secret_Text_Delete{}
				subFlags = flag.NewFlagSet("client secret text delete", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.Text.Delete.Id, "id", "", "text secret id")
				subFlags.Parse(args)
				if target.Secret.Text.Delete.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}

			case "update":
				target.Secret.Text.Update = &PipelineInitTree_Secret_Text_Update{}
				subFlags = flag.NewFlagSet("client secret text update", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.Text.Update.Id, "id", "", "text secret id")
				subFlags.StringVar(&target.Secret.Text.Update.Data, "data", "", "text secret data")
				subFlags.Parse(args)
				if target.Secret.Text.Update.Id == "" || target.Secret.Text.Update.Data == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			default:
				subFlags.Usage()
				os.Exit(1)
			}
		case "login-password":
			target.Secret.LoginPassword = &PipelineInitTree_Secret_LoginPassword{}
			subFlags := flag.NewFlagSet("login-password", flag.ExitOnError)
			subFlags.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: ./client secret text\n")
				fmt.Fprintf(os.Stderr, "sub-command:\n")
				fmt.Fprintf(os.Stderr, "    get\n")
				fmt.Fprintf(os.Stderr, "    create\n")
				fmt.Fprintf(os.Stderr, "    delete\n")
				fmt.Fprintf(os.Stderr, "    update\n")
			}
			subFlags.Parse(args)
			if len(args) < 1 {
				subFlags.Usage()
				os.Exit(1)
			}
			subcommand = args[0]
			args = args[1:]
			switch subcommand {
			case "get":
				target.Secret.LoginPassword.Get = &PipelineInitTree_Secret_LoginPassword_Get{}
				subFlags = flag.NewFlagSet("client secret login-password get", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.LoginPassword.Get.Id, "id", "", "login-password secret id")
				subFlags.Parse(args)
				if target.Secret.LoginPassword.Get.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "create":
				target.Secret.LoginPassword.Create = &PipelineInitTree_Secret_LoginPassword_Create{}
				subFlags = flag.NewFlagSet("client secret login password create", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.LoginPassword.Create.Login, "login", "", "login-password secret login")
				subFlags.StringVar(&target.Secret.LoginPassword.Create.Password, "password", "", "login-password secret password")
				subFlags.Parse(args)
				if target.Secret.LoginPassword.Create.Login == "" || target.Secret.LoginPassword.Create.Password == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "delete":
				target.Secret.LoginPassword.Delete = &PipelineInitTree_Secret_LoginPassword_Delete{}
				subFlags = flag.NewFlagSet("client secret login password delete", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.LoginPassword.Delete.Id, "id", "", "text secret id")
				subFlags.Parse(args)
				if target.Secret.LoginPassword.Delete.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}

			case "update":
				target.Secret.LoginPassword.Update = &PipelineInitTree_Secret_LoginPassword_Update{}
				subFlags = flag.NewFlagSet("client secret login password update", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.LoginPassword.Update.Id, "id", "", "login password secret id")
				subFlags.StringVar(&target.Secret.LoginPassword.Update.Login, "login", "", "login secret data")
				subFlags.StringVar(&target.Secret.LoginPassword.Update.Password, "password", "", "login secret data")
				subFlags.Parse(args)
				if target.Secret.LoginPassword.Update.Id == "" ||
					target.Secret.LoginPassword.Update.Login == "" ||
					target.Secret.LoginPassword.Update.Password == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			default:
				subFlags.Usage()
				os.Exit(1)
			}
		case "credit-card":
			target.Secret.CreditCard = &PipelineInitTree_Secret_CreditCard{}
			subFlags := flag.NewFlagSet("credit-card", flag.ExitOnError)
			subFlags.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: ./client secret credit-card\n")
				fmt.Fprintf(os.Stderr, "sub-command:\n")
				fmt.Fprintf(os.Stderr, "    get\n")
				fmt.Fprintf(os.Stderr, "    create\n")
				fmt.Fprintf(os.Stderr, "    delete\n")
				fmt.Fprintf(os.Stderr, "    update\n")
			}
			subFlags.Parse(args)
			if len(args) < 1 {
				subFlags.Usage()
				os.Exit(1)
			}
			subcommand = args[0]
			args = args[1:]
			switch subcommand {
			case "get":
				target.Secret.CreditCard.Get = &PipelineInitTree_Secret_CreditCard_Get{}
				subFlags = flag.NewFlagSet("client secret credit-card get", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.CreditCard.Get.Id, "id", "", "credit card secret id")
				subFlags.Parse(args)
				if target.Secret.CreditCard.Get.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "create":
				target.Secret.CreditCard.Create = &PipelineInitTree_Secret_CreditCard_Create{}
				subFlags = flag.NewFlagSet("client secret credit card create", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.CreditCard.Create.CardNumber, "number", "", "credit card number")
				subFlags.StringVar(&target.Secret.CreditCard.Create.CardHolder, "owner", "", "credit card owner")
				subFlags.StringVar(&target.Secret.CreditCard.Create.ValidTill, "valid-till", "", "credit card expiration date")
				subFlags.StringVar(&target.Secret.CreditCard.Create.Cvc, "cvc", "", "credit card Cvc")
				subFlags.Parse(args)
				if target.Secret.CreditCard.Create.CardNumber == "" ||
					target.Secret.CreditCard.Create.CardHolder == "" ||
					target.Secret.CreditCard.Create.ValidTill == "" ||
					target.Secret.CreditCard.Create.Cvc == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "delete":
				target.Secret.CreditCard.Delete = &PipelineInitTree_Secret_CreditCard_Delete{}
				subFlags = flag.NewFlagSet("client secret credit card delete", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.CreditCard.Delete.Id, "id", "", "credit card secret id")
				subFlags.Parse(args)
				if target.Secret.CreditCard.Delete.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}

			case "update":
				target.Secret.CreditCard.Update = &PipelineInitTree_Secret_CreditCard_Update{}
				subFlags = flag.NewFlagSet("client secret credit card update", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.CreditCard.Update.Id,
					"id", "", "credit card secret id")
				subFlags.StringVar(&target.Secret.CreditCard.Update.CardNumber,
					"number", "", "credit card secret card number")
				subFlags.StringVar(&target.Secret.CreditCard.Update.CardHolder,
					"owner", "", "credit card secret cardholder")
				subFlags.StringVar(&target.Secret.CreditCard.Update.ValidTill,
					"valid-till", "", "credit card secret valid-till")
				subFlags.StringVar(&target.Secret.CreditCard.Update.Cvc,
					"cvc", "", "credit card secret cvc")
				subFlags.Parse(args)
				if target.Secret.CreditCard.Update.Id == "" || target.Secret.CreditCard.Update.CardNumber == "" ||
					target.Secret.CreditCard.Update.CardHolder == "" ||
					target.Secret.CreditCard.Update.ValidTill == "" ||
					target.Secret.CreditCard.Update.Cvc == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			default:
				subFlags.Usage()
				os.Exit(1)
			}
		case "file":
			target.Secret.File = &PipelineInitTree_Secret_File{}
			subFlags = flag.NewFlagSet("file", flag.ExitOnError)
			subFlags.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: ./client secret file\n")
				fmt.Fprintf(os.Stderr, "sub-command:\n")
				fmt.Fprintf(os.Stderr, "    download\n")
				fmt.Fprintf(os.Stderr, "    upload\n")
				fmt.Fprintf(os.Stderr, "    delete\n")
				fmt.Fprintf(os.Stderr, "    update\n")
			}
			subFlags.Parse(args)
			if len(args) < 1 {
				subFlags.Usage()
				os.Exit(1)
			}
			subcommand = args[0]
			args = args[1:]
			switch subcommand {
			case "download":
				target.Secret.File.Download = &PipelineInitTree_Secret_File_Download{}
				subFlags = flag.NewFlagSet("client secret file download", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.File.Download.Id, "id", "", "file secret id")
				subFlags.StringVar(&target.Secret.File.Download.FilePath, "path", "", "file secret id")
				subFlags.Parse(args)
				if target.Secret.File.Download.Id == "" || target.Secret.File.Download.FilePath == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "upload":
				target.Secret.File.Upload = &PipelineInitTree_Secret_File_Upload{}
				subFlags = flag.NewFlagSet("client secret file upload", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.File.Upload.FilePath, "path", "", "file secret id")
				subFlags.Parse(args)
				if target.Secret.File.Upload.FilePath == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "delete":
				target.Secret.File.Delete = &PipelineInitTree_Secret_File_Delete{}
				subFlags = flag.NewFlagSet("client secret file delete", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.File.Delete.Id, "id", "", "file secret id")
				subFlags.Parse(args)
				if target.Secret.File.Delete.Id == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			case "update":
				target.Secret.File.Update = &PipelineInitTree_Secret_File_Update{}
				subFlags = flag.NewFlagSet("client secret file update", flag.ExitOnError)
				subFlags.StringVar(&target.Secret.File.Update.Id, "id", "", "file secret id")
				subFlags.StringVar(&target.Secret.File.Update.FilePath, "path", "", "file secret id")
				subFlags.Parse(args)
				if target.Secret.File.Update.Id == "" || target.Secret.File.Update.FilePath == "" {
					subFlags.Usage()
					os.Exit(1)
				}
			default:
				subFlags.Usage()
				os.Exit(1)
			}
		case "list":
			target.Secret.List = &PipelineInitTree_Secret_List{}
			subFlags = flag.NewFlagSet("file", flag.ExitOnError)
			subFlags.Usage = func() {
				fmt.Fprintf(os.Stderr, "Usage: ./client secret list\n")
			}
			subFlags.Parse(args)
			if len(args) != 0 {
				subFlags.Usage()
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "Unsupported command\n %s", subcommand)
			subFlags.Usage()
			os.Exit(1)
		}

	default:
		flag.Usage()
		os.Exit(1)
	}

	return target
}
