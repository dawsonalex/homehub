package main

import (
	"context"
	"github.com/dawsonalex/homehub/cmd/macrod/schema"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := schema.NewMacrodClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	versionResponse, err := c.Version(ctx, &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	log.Printf("talking to server version: %v\n", versionResponse)

	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	r, err := c.CreateFoodListing(ctx, &schema.FoodListing{
		Name: "Test Food",
		MacrosPerServing: map[string]*schema.Macros{
			"100g": {
				Carbs:    10,
				Fats:     5,
				Proteins: 40,
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Created food listing: %v\n", r)

	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.CreateMeal(ctx, &schema.Meal{
	//	Name:     "Test Meal",
	//	Food:     make([]*schema.FoodEntry),
	//	Calories: 0,
	//	Macros:   nil,
	//})
}
