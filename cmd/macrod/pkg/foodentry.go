package pkg

import "github.com/google/uuid"

type FoodEntry struct {
	FoodListing
	ID                uuid.UUID `json:"id"`
	SelectedServingId uuid.UUID
	Quantity          float64
}

func (f FoodEntry) SelectedServing() Serving {
	return f.servings[f.SelectedServingId]
}

// Calories returns the number of calories for 100g of Food to nearest whole number.
func (f FoodEntry) Calories() int {
	return int(float64(f.SelectedServing().Calories()) * f.Quantity)
}
