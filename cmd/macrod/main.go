package main

import (
	"context"
	"fmt"
	"github.com/dawsonalex/homehub/cmd/macrod/pkg"
	"github.com/dawsonalex/homehub/cmd/macrod/repo"
	"github.com/dawsonalex/homehub/cmd/macrod/schema"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//db, err := postgres.OpenConnection(postgres.Config{
	//	User:   "postgres",
	//	Dbname: "macrod",
	//})
	//if err != nil {
	//	panic(err)
	//}

	r := repo.NewInMemory()

	foodListing := pkg.FoodListing{
		Name: "Apple",
	}

	foodListing.AddServing(pkg.NewServing("100g", pkg.NewMacros(10, 0, 0)))

	foodListing, err = r.AddFoodListing(context.TODO(), foodListing)
	panicOnErr(err)

	meal := pkg.Meal{
		Name: "Lunch",
	}

	meal.AddFood(pkg.FoodEntry{
		FoodListing:       foodListing,
		SelectedServingId: foodListing.Servings()[0].Id,
		Quantity:          1,
	})

	schema.RegisterMacrodServer(s, &server{
		repo: r,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
