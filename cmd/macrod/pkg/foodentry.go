package pkg

type FoodEntry struct {
	FoodListing
	selectedServingName string
	quantity            float64
}

func (f FoodEntry) selectedServing() Serving {
	return f.servings[f.selectedServingName]
}

// Calories returns the number of calories for 100g of Food to nearest whole number.
func (f FoodEntry) Calories() int {
	return int(float64(f.selectedServing().Calories()) * f.quantity)
}
