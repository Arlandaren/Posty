package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"service/internal/repository"
	"service/internal/service"
	"service/internal/shared/storage/postgres"
	pb "service/pkg/grpc/posty_v1"
	"sync"
	"syscall"
	"time"

	transport "service/internal/transport/grpc"
)

const (
	grpcAddress = ":50051"
	httpAddress = ":8086"
)

func main() {
	db, err := postgres.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Printf("DB connected")

	repo := repository.NewRepository(db)

	serv := service.NewService(repo)

	server := transport.NewServer(serv)

	wg := &sync.WaitGroup{}
	ctx := context.Background()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := runGrpcServer(ctx, server)
		if err != nil {
			log.Printf("failed to run grpc server: %v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := startHttpServer(ctx)
		if err != nil {
			log.Printf("failed to run http server: %v", err)
			return
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down servers...")

	wg.Wait()
	log.Println("Servers gracefully stopped.")
}

func runGrpcServer(ctx context.Context, server *transport.Server) error {

	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(service.AuthInterceptor),
	)

	reflection.Register(grpcServer)
	pb.RegisterPostyServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", grpcAddress, err)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Printf("gRPC server failed: %v", err)
		}
	}()

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	<-ctx.Done()

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	return nil
}

func startHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := pb.RegisterPostyServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("failed to register service handler: %w", err)
	}

	handler := allowCORS(mux)

	srv := &http.Server{
		Addr:    httpAddress,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server exited with error: %v", err)
		}
	}()

	log.Printf("HTTP server listening at %v\n", httpAddress)

	<-ctx.Done()

	log.Println("Shutting down HTTP server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("HTTP server Shutdown failed: %w", err)
	}

	return nil
}
