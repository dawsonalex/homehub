package pkg

const (
	CaloriesPer1gProtein      = 4
	CaloriesPer1gCarbohydrate = 4
	CaloriesPer1gFat          = 9
)

type Serving struct {
	size   string
	macros Macros
}

func (s Serving) Calories() int {
	return int((s.macros.carbs * CaloriesPer1gCarbohydrate) *
		(s.macros.fats * CaloriesPer1gFat) *
		(s.macros.proteins * CaloriesPer1gProtein))
}

func (s Serving) Macros() (carbs, fats, proteins float64) {
	return s.macros.carbs, s.macros.fats, s.macros.proteins

}

func NewServing(size string, macros Macros) Serving {
	return Serving{
		size:   size,
		macros: macros,
	}
}
