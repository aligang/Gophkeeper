package main

import (
	"context"
	"fmt"
	"github.com/aligang/Gophkeeper/internal/account"
	accountHandler "github.com/aligang/Gophkeeper/internal/account/handler"
	"github.com/aligang/Gophkeeper/internal/config"
	"github.com/aligang/Gophkeeper/internal/gb/tokengc"
	"github.com/aligang/Gophkeeper/internal/logging"
	"github.com/aligang/Gophkeeper/internal/repository"
	"github.com/aligang/Gophkeeper/internal/repository/fs"
	"github.com/aligang/Gophkeeper/internal/secret"
	secretHandler "github.com/aligang/Gophkeeper/internal/secret/handler"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	logging.Configure(os.Stdout, zerolog.DebugLevel)
	exitSignal := make(chan os.Signal, 1)
	stopServer := make(chan any, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	wg := sync.WaitGroup{}

	cfg := config.GetServerConfig()
	fmt.Println(cfg)
	storage := repository.New(cfg)
	fileStorage := fs.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		<-exitSignal
		stopServer <- struct{}{}
		close(stopServer)
		cancel()
		wg.Done()
	}()

	//wg.Add(1)
	//go func() {
	//	logging.Debug("starting FileSystemGC...")
	//	fsGarbageCollection := fsgc.New(cfg, storage, fileStorage)
	//	fsGarbageCollection.Run(ctx)
	//	logging.Debug("FileSystemGC stopped")
	//	wg.Done()
	//}()

	wg.Add(1)
	go func() {
		logging.Debug("starting TokenGC...")
		tokenGarbageCollector := tokengc.New(cfg, storage)
		tokenGarbageCollector.Run(ctx)
		logging.Debug("TokenGC stopped")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		listen, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal(err)
		}
		accountHandler := accountHandler.New(storage, cfg)
		secretHandler := secretHandler.New(storage, fileStorage)
		s := grpc.NewServer(grpc.ChainUnaryInterceptor(accountHandler.AuthInterceptor))

		account.RegisterAccountServiceServer(s, accountHandler)
		secret.RegisterSecretServiceServer(s, secretHandler)
		logging.Debug("GRPC Server Starts...")

		wg.Add(1)
		go func() {
			<-stopServer
			logging.Debug("GRPC Server Stops...")
			s.GracefulStop()
			close(stopServer)
			wg.Done()
		}()

		if err := s.Serve(listen); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}