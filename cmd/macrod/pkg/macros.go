package pkg

// Macros represents macro values for a food source.
// They are useful only when used relative to a serving or quantity.
type Macros struct {
	Carbs    float64
	Fats     float64
	Proteins float64
}

func NewMacros(carbs, fats, proteins float64) Macros {
	return Macros{
		Carbs:    carbs,
		Fats:     fats,
		Proteins: proteins,
	}
}
