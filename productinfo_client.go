package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thanhftu/go-client/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connected %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
	name := "Apple 11"
	price := float32(1000)
	description := "Meet Apple 11"
	ctx, canncel := context.WithTimeout(context.Background(), time.Second)
	defer canncel()
	r, err := c.AddProduct(ctx,
		&pb.Product{
			Name:        name,
			Description: description,
			Price:       price,
		})
	if err != nil {
		log.Fatalf("Could not add product %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s ", product.String())

}
