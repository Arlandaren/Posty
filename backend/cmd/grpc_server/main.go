package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	httpHandler "service/internal/transport/http"

	"net/http"
	"os"
	"os/signal"
	"service/internal/repository"
	"service/internal/service"
	"service/internal/shared/storage/postgres"
	transport "service/internal/transport/grpc"
	pb "service/pkg/grpc/posty_v1"
	"sync"
	"syscall"
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

	// Создаем HTTP mux для обработки как grpc-gateway запросов, так и кастомных HTTP обработчиков
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	httpMux.HandleFunc("/v1/images", httpHandler.UploadImageHandler)
	httpMux.HandleFunc("/v1/images/", httpHandler.GetImageHandler) // для получения изображений

	fmt.Println("http обработчики")

	// Оборачиваем с помощью CORS, если требуется
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"}, // ваш фронтенд адрес
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(httpMux)

	server := &http.Server{
		Addr:    httpAddress,
		Handler: handler,
	}

	go func() {
		log.Printf("HTTP server listening at %v\n", httpAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to serve HTTP server: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("Shutting down HTTP server...")
	return server.Shutdown(context.Background())
}
