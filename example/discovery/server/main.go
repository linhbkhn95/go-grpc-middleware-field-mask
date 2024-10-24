package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	fieldmaskpkg "github.com/linhbkhn95/go-grpc-middleware-field-mask"
	pb "github.com/linhbkhn95/go-grpc-middleware-field-mask/pb/go/discovery/v1"
)

var (
	port = 10080
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			fieldmaskpkg.UnaryServerInterceptor(fieldmaskpkg.DefaultFilterFunc),
		),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDiscoveryServiceServer(grpcServer, NewDiscoveryServer())
	grpcServer.Serve(lis)
}

type DiscoveryServer struct {
	pb.UnimplementedDiscoveryServiceServer
}

func NewDiscoveryServer() *DiscoveryServer {
	return &DiscoveryServer{}
}

func (s *DiscoveryServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products := []*pb.Product{
		{
			Id:    "1",
			Name:  "Product 1",
			Img:   "Image 1",
			Price: 1,
			Shop: &pb.Shop{
				Id:   "1",
				Name: "Shop 1",
			},
		},
		{
			Id:    "2",
			Name:  "Product 2",
			Img:   "Image 2",
			Price: 1,
			Shop: &pb.Shop{
				Id:   "2",
				Name: "Shop 2",
			},
		},
		{
			Id:    "3",
			Name:  "Product 3",
			Img:   "Image 3",
			Price: 1,
			Shop: &pb.Shop{
				Id:   "3",
				Name: "Shop 3",
			},
		},
		{
			Id:    "4",
			Name:  "Product 4",
			Img:   "Image 4",
			Price: 1,
			Shop: &pb.Shop{
				Id:   "4",
				Name: "Shop 4",
			},
		},
	}
	return &pb.ListProductsResponse{
		Result: &pb.ListProductsResult{
			Products: products,
		},
	}, nil
}
