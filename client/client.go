package main

import (
	"context"
	"fmt"
	pb "github.com/graywolfff/test_grpc/coffeeshop_protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStrem, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	var items []*pb.Item

	go func() {
		for {
			resp, err := menuStrem.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			items = resp.Items
			log.Printf("Got items: %v", resp.Items)
		}
	}()
	<-done
	fmt.Println(items)
}
