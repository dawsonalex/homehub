package repo

import (
	"context"
	"github.com/dawsonalex/homehub/cmd/macrod/pkg"
)

type FoodRepo interface {
	AddFoodListing(ctx context.Context, food pkg.FoodListing) (pkg.FoodListing, error)

	AddMeal(ctx context.Context, meal pkg.Meal) (pkg.Meal, error)
}
