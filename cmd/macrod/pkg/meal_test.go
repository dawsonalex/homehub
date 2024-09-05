package pkg

import (
	"github.com/google/uuid"
	"testing"
)

func TestMeal(t *testing.T) {
	pastaListing := NewFoodListing("pasta")
	pastaListing.AddServing(Serving{
		Id:     uuid.New(),
		size:   "100g",
		macros: NewMacros(35, 0, 5),
	})

	chickenListing := NewFoodListing("chicken")
	chickenListing.AddServing(Serving{
		Id:     uuid.New(),
		size:   "100g",
		macros: NewMacros(0, 4, 40),
	})

	meal := Meal{
		Id:   uuid.New(),
		Name: "Lunch",
	}

	meal.AddFood(FoodEntry{
		FoodListing:       pastaListing,
		ID:                uuid.New(),
		SelectedServingId: pastaListing.Servings()[0].Id,
		Quantity:          1,
	})

	meal.AddFood(FoodEntry{
		FoodListing:       chickenListing,
		ID:                uuid.New(),
		SelectedServingId: chickenListing.Servings()[0].Id,
		Quantity:          1,
	})

	// Just hard coding the calorie count for this test.
	expectedCals := 4*CaloriesPer1gFat + 40*CaloriesPer1gProtein + 35*CaloriesPer1gCarbohydrate + 5*CaloriesPer1gProtein
	if _, cals := meal.Nutrition(); cals != expectedCals {
		t.Errorf("Expected Nutrition calories to be %d, got %d", expectedCals, cals)
	}
}
