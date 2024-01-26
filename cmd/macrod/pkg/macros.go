package pkg

// Macros represents macro values for a food source.
// They are useful only when used relative to a serving or quantity.
type Macros struct {
	carbs    float64
	fats     float64
	proteins float64
}

func NewMacros(carbs, fats, proteins float64) *Macros {
	return &Macros{
		carbs:    carbs,
		fats:     fats,
		proteins: proteins,
	}
}
