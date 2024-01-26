package main

import (
	context "context"
	"github.com/dawsonalex/homehub/cmd/macrod/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct {
	schema.UnimplementedMacrodServer
}

func (s *server) CreateFoodListing(ctx context.Context, listing *schema.FoodListing) (*schema.FoodListing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) GetFoodListing(ctx context.Context, value *wrapperspb.StringValue) (*schema.FoodListing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateFoodListing(ctx context.Context, listing *schema.FoodListing) (*schema.FoodListing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) mustEmbedUnimplementedMacrodServer() {
	//TODO implement me
	panic("implement me")
}
