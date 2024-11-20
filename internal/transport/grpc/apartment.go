package grpc

import (
	"service/internal/service"
	desc "service/pkg/grpc/apartment_v1"
)

type Server struct {
	desc.UnimplementedApartmentServiceServer
	Service *service.Service
}

func NewServer(Service *service.Service) *Server {
	return &Server{
		Service: Service,
	}
}
