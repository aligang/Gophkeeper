package main

import (
	"context"
	"github.com/aligang/Gophkeeper/internal/common/account"
	"github.com/aligang/Gophkeeper/internal/common/logging"
	"github.com/aligang/Gophkeeper/internal/common/secret"
	accountHandler "github.com/aligang/Gophkeeper/internal/server/account/handler"
	"github.com/aligang/Gophkeeper/internal/server/config"
	"github.com/aligang/Gophkeeper/internal/server/gb/fsgc"
	"github.com/aligang/Gophkeeper/internal/server/gb/tokengc"
	"github.com/aligang/Gophkeeper/internal/server/repository"
	"github.com/aligang/Gophkeeper/internal/server/repository/fs"
	secretHandler "github.com/aligang/Gophkeeper/internal/server/secret/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	logging.Init(os.Stdout)
	logging.SetLogLevel(cfg.LogLevel)
	exitSignal := make(chan os.Signal, 1)
	stopServer := make(chan any, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	wg := sync.WaitGroup{}

	storage := repository.New(cfg)
	fileStorage := fs.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		<-exitSignal
		stopServer <- struct{}{}
		close(exitSignal)
		cancel()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logging.Info("starting FileSystemGC...")
		fsGarbageCollection := fsgc.New(cfg, storage, fileStorage)
		fsGarbageCollection.Run(ctx)
		logging.Info("FileSystemGC stopped")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logging.Info("starting TokenGC...")
		tokenGarbageCollector := tokengc.New(cfg, storage)
		tokenGarbageCollector.Run(ctx)
		logging.Info("TokenGC stopped")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		listen, err := net.Listen("tcp", cfg.Address)
		if err != nil {
			logging.Fatal(err.Error())
		}
		accountHandler := accountHandler.New(storage, cfg)
		secretHandler := secretHandler.New(storage, fileStorage, cfg)

		serverOpts := []grpc.ServerOption{grpc.ChainUnaryInterceptor(accountHandler.AuthInterceptor)}
		if cfg.TlsCertPath != "" && cfg.TlsKeyPath != "" {
			creds, err := credentials.NewServerTLSFromFile(cfg.TlsCertPath, cfg.TlsKeyPath)
			if err != nil {
				logging.Fatal("Could not read Certificate file: %s", err.Error())
			}
			serverOpts = append(serverOpts, grpc.Creds(creds))
		}

		s := grpc.NewServer(serverOpts...)

		account.RegisterAccountServiceServer(s, accountHandler)
		secret.RegisterSecretServiceServer(s, secretHandler)
		logging.Info("GRPC Server Starts...")

		wg.Add(1)
		go func() {
			<-stopServer
			logging.Info("GRPC Server Stops...")
			s.GracefulStop()
			close(stopServer)
			wg.Done()
		}()

		if err := s.Serve(listen); err != nil {
			logging.Fatal(err.Error())
		}
		wg.Done()
	}()

	wg.Wait()
}
