package repo

import (
	"context"
	"github.com/dawsonalex/homehub/cmd/macrod/pkg"
	"github.com/google/uuid"
)

var _ FoodRepo = &InMemory{}

// InMemory provides an in-memory implementation of the repository.
type InMemory struct {
	foodListings map[uuid.UUID]pkg.FoodListing
	meals        map[uuid.UUID]pkg.Meal
}

func NewInMemory() *InMemory {
	return &InMemory{
		foodListings: make(map[uuid.UUID]pkg.FoodListing),
		meals:        make(map[uuid.UUID]pkg.Meal),
	}
}

func (i InMemory) AddFoodListing(ctx context.Context, food pkg.FoodListing) (pkg.FoodListing, error) {
	newUUID := uuid.New()
	food.Id = newUUID
	i.foodListings[newUUID] = food
	return food, nil
}

func (i InMemory) AddMeal(ctx context.Context, meal pkg.Meal) (pkg.Meal, error) {
	newUUID := uuid.New()
	meal.Id = newUUID
	i.meals[newUUID] = meal
	return meal, nil
}
