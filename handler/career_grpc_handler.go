package handler

import (
	"net"
	"sims-daas/usecase"

	"google.golang.org/grpc"
)

// GrpcServer adalah struktur pembungkus server gRPC kita
type GrpcServer struct {
	Usecase usecase.CareerUsecaseInterface
}

func NewGrpcServer(u usecase.CareerUsecaseInterface) *GrpcServer {
	return &GrpcServer{Usecase: u}
}

// Run berfungsi untuk menyalakan server gRPC di port terpisah (:50051)
func (s *GrpcServer) Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	println("📡 Server gRPC Enterprise aktif dan mendengarkan di port :50051")

	return server.Serve(lis)
}
