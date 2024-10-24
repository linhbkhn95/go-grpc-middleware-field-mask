package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/linhbkhn95/go-grpc-middleware-field-mask/pb/go/discovery/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	address := "localhost:10080"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewDiscoveryServiceClient(conn)

	// for normal response that return all fields
	normalRequest := &pb.ListProductsRequest{
		Id: "1"}

	resp, err := client.ListProducts(context.Background(), normalRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resp for normal request", resp)

	// request that use field mask
	fieldMaskReq := &pb.ListProductsRequest{
		Id: "1", FieldMask: &fieldmaskpb.FieldMask{
			Paths: []string{"result.products.id", "result.products.price"},
		}}

	resp, err = client.ListProducts(context.Background(), fieldMaskReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resp with field mask", resp)
}
