package main

import (
	context "context"
	"github.com/dawsonalex/homehub/cmd/macrod/pkg"
	"github.com/dawsonalex/homehub/cmd/macrod/repo"
	"github.com/dawsonalex/homehub/cmd/macrod/schema"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"strings"
)

const version = "0.1.0.d8294dfae"

type versionDetails struct {
	major string
	minor string
	point string
	build string
}

type server struct {
	schema.UnimplementedMacrodServer
	repo repo.FoodRepo
}

func (s *server) CreateFoodListing(ctx context.Context, listing *schema.FoodListing) (*schema.FoodListing, error) {
	newListing := pkg.NewFoodListing(listing.Name)
	for size, serving := range listing.MacrosPerServing {
		// TODO: sort real num sizing between DB and proto
		newListing.AddServing(pkg.NewServing(size, pkg.NewMacros(float64(serving.Carbs), float64(serving.Fats), float64(serving.Proteins))))
	}

	newListing, err := s.repo.AddFoodListing(ctx, newListing)
	if err != nil {
		return nil, errors.Wrap(err, "error adding food listing")
	}

	listing.Id = newListing.Id.String()
	return listing, nil
}

func (s *server) GetFoodListing(ctx context.Context, value *wrapperspb.StringValue) (*schema.FoodListing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateFoodListing(ctx context.Context, listing *schema.FoodListing) (*schema.FoodListing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) CreateMeal(ctx context.Context, meal *schema.Meal) (*schema.Meal, error) {
	newMeal := pkg.Meal{
		Name: meal.Name,
	}

	newMeal, err := s.repo.AddMeal(ctx, newMeal)
	if err != nil {
		return nil, errors.Wrap(err, "error adding meal")
	}

	// set up food entries
	food := newMeal.GetFood()
	schemaFoodEntries := make([]*schema.FoodEntry, len(food))
	for i, foodEntry := range food {
		entry := pkgToSchemaFoodEntry(foodEntry)
		schemaFoodEntries[i] = &entry
	}

	macros, calories := newMeal.Nutrition()
	return &schema.Meal{
		Id:       newMeal.Id.String(),
		Name:     newMeal.Name,
		Food:     schemaFoodEntries,
		Calories: uint32(calories),
		Macros: &schema.Macros{
			// TODO: converting between float64 and float32 is silly. Pick one.
			Carbs:    float32(macros.Carbs),
			Fats:     float32(macros.Fats),
			Proteins: float32(macros.Proteins),
		},
	}, nil
}

func (s *server) AddFoodToMeal(context.Context, *schema.UpdateMealRequest) (*schema.Meal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFoodToMeal not implemented")
}

func (s *server) GetMeal(context.Context, *wrapperspb.StringValue) (*schema.Meal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeal not implemented")
}

func (s *server) mustEmbedUnimplementedMacrodServer() {
	//TODO implement me
	panic("implement me")
}

func pkgToSchemaFoodEntry(entry pkg.FoodEntry) schema.FoodEntry {
	return schema.FoodEntry{
		//Id:                entry.Id,
		ListingId:         entry.FoodListing.Id.String(),
		SelectServingSize: entry.SelectedServing().Size(),
		ServingQuantity:   float32(entry.Quantity),
	}
}

func (s *server) Version(context.Context, *emptypb.Empty) (*schema.VersionResponse, error) {
	splitVersion := strings.Split(version, ".")
	return &schema.VersionResponse{
		Major: splitVersion[0],
		Minor: splitVersion[1],
		Point: splitVersion[2],
		Build: splitVersion[3],
	}, nil
}
