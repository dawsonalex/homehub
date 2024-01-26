package pkg

type Meal struct {
	name    string
	entries []FoodEntry
}

// Calories return the total calories for a meal
func (m Meal) Calories() int {
	totalCalories := 0
	for _, entry := range m.entries {
		totalCalories += entry.Calories()
	}

	return totalCalories
}

// Nutrition returns the total Macros and calories for a meal in one call.
func (m Meal) Nutrition() (Macros, int) {
	macros := Macros{}
	totalCalories := 0
	for _, entry := range m.entries {
		macros.carbs += entry.selectedServing().macros.carbs
		macros.fats += entry.selectedServing().macros.fats
		macros.proteins += entry.selectedServing().macros.proteins
		totalCalories += entry.Calories()
	}

	return macros, totalCalories
}
