package main

import (
	"context"
	pb "github.com/graywolfff/test_grpc/coffeeshop_protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(_ *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		{
			Id:   "1",
			Name: "Black Coffee",
		},
		{
			Id:   "2",
			Name: "Americano",
		},
	}
	for i := range items {
		err := srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (s *server) PlaceOder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}
func (s *server) GetOrderStatus(_ context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		Status:  "IN PROGRESS",
		OrderId: receipt.Id,
	}, nil
}

func main() {
	// Set up a listener on port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})
	log.Printf("server listening at :8080")
	log.Fatal(grpcServer.Serve(lis))
}
