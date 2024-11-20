package main

import (
	"context"
	"service/internal/repository"
	"service/internal/service"

	"service/internal/shared/config"
	"service/internal/shared/kafka"
	"service/internal/shared/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"log"
	"net"
	"os"

	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc/credentials/insecure"

	transport "service/internal/transport/grpc"

	desc "service/pkg/grpc/apartment_v1"
)

const grpcAddress = ":8080"

func main() {
	postgresCfg, err := config.GetPostgres()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	kafkaCfg := kafka.NewConfig([]string{os.Getenv("KAFKA_BROKER_ADDRESS")})

	kafkaProducer, err := kafka.NewProducer(kafkaCfg)

	dbPool, err := postgres.InitPostgres(postgresCfg, 5)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	defer dbPool.Close()

	pgxWrapper := postgres.NewWrapper(dbPool)

	appRepository := repository.NewRepository(pgxWrapper)

	appService := service.NewService(appRepository, kafkaProducer)

	appServer := transport.NewServer(appService)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := startGrpcServer(ctx, appServer); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gRPC server...")

	cancel()
	wg.Wait()
	log.Println("Server gracefully stopped.")
}

func startGrpcServer(ctx context.Context, appServer *transport.Server) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	desc.RegisterApartmentServiceServer(grpcServer, appServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	<-ctx.Done()

	grpcServer.GracefulStop()
	return nil
}
