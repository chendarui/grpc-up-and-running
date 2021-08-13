package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "productinfo/client/ecommerce"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "Apple IPhone 11"
	description := `Meet Apple IPhone 11. All沙坑里睡觉的力气`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	productID, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("Could not add product:%v", err)
	}
	log.Printf("Product ID:%s added successfully", productID.Value)

	getProduct, err := c.GetProduct(ctx, &pb.ProductID{Value: productID.Value})
	if err != nil {
		log.Fatalf("Could not get product:%v", err)
	}
	log.Printf("Product: %v", getProduct.String())

}



