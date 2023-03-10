package config

import (
	"flag"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"os"
	"strings"
)

func getConfigFromCli() *Config {
	conf := &Config{}
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of ./server: server [options] in-memory|sql \n")
		fmt.Fprintf(os.Stderr, "options:\n")
		fmt.Fprintf(os.Stderr, "		-a 'host to listen on'.\n")
		fmt.Fprintf(os.Stderr, "		-f 'file Storage Path'.\n")
		fmt.Fprintf(os.Stderr, "		-c 'config File Path'.\n")
		fmt.Fprintf(os.Stderr, "		-tv 'token validity time'.\n")
		fmt.Fprintf(os.Stderr, "		-tr 'token renewal time'.\n")
		fmt.Fprintf(os.Stderr, "		-fs 'file stale time'.\n")
		fmt.Fprintf(os.Stderr, "		-e 'enable secret encryption'.\n")
		fmt.Fprintf(os.Stderr, "		-cert 'certificate file path'.\n")
		fmt.Fprintf(os.Stderr, "		-cert_key 'certificate key file  path'.\n")
	}

	flag.StringVar(&conf.Address, "a", "", "host to listen on")
	flag.StringVar(&conf.FileStorage, "f", "", "File Storage Path")
	flag.StringVar(&conf.ConfigFile, "c", "", "Config File Path")
	flag.Int64Var(&conf.TokenValidityTimeMinutes, "tv", -1, "token")
	flag.Int64Var(&conf.TokenRenewalTimeMinutes, "tr", -1, "File Storage Path")
	flag.Int64Var(&conf.FileStaleTimeMinutes, "fs", -1, "Config File Path")
	flag.BoolVar(&conf.SecretEncryptionEnabled, "e", false, "Enable encryption")
	flag.StringVar(&conf.TlsCertPath, "cert", "", "Tls Cert File Path")
	flag.StringVar(&conf.TlsKeyPath, "cert_key", "", "Tls Cert Key Path")

	var logLevel string
	flag.StringVar(&logLevel, "log-level", "", "Loggin level")

	flag.Parse()
	conf.LogLevel = logging.GetLogLevelFromString(logLevel)
	parsedRepoType := flag.Arg(0)
	//if parsedRepoType != "in-memory" && parsedRepoType != "sql" {
	//	return conf
	//}

	repoType := strings.ReplaceAll(strings.ToUpper(parsedRepoType), "-", "_")
	conf.RepositoryType = getRepoValueFromName(&repoType)
	if conf.RepositoryType == RepositoryType_SQL {
		if len(flag.Args()) == 2 {
			conf.DatabaseDsn = flag.Arg(1)
		} else if len(flag.Args()) == 1 {
			fmt.Fprintf(
				os.Stderr,
				"Incomplete command: ./server %s. Use ./server sql {datebase_dsn} \n",
				strings.Join(flag.Args(), " "))
		} else {
			fmt.Fprintf(
				os.Stderr,
				"Excessive command arguments: %s in ./server %s. Use ./server sql {datebase_dsn} \n",
				strings.Join(flag.Args()[2:], " "),
				strings.Join(flag.Args(), " "))
		}
	}
	//fmt.Println(conf)
	return conf
}
