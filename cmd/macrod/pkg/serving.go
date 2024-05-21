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
	return int((s.macros.Carbs * CaloriesPer1gCarbohydrate) *
		(s.macros.Fats * CaloriesPer1gFat) *
		(s.macros.Proteins * CaloriesPer1gProtein))
}

func (s Serving) Macros() (carbs, fats, proteins float64) {
	return s.macros.Carbs, s.macros.Fats, s.macros.Proteins

}

func (s Serving) Size() string {
	return s.size
}

func NewServing(size string, macros Macros) Serving {
	return Serving{
		size:   size,
		macros: macros,
	}
}
