package pkg

type FoodListing struct {
	name     string
	servings map[string]Serving
}

func (f FoodListing) Name() string {
	return f.name
}

func (f FoodListing) Servings() []Serving {
	// Note: This might need making safe for concurrent use at some point
	// TODO: cache this list
	servings := make([]Serving, len(f.servings))
	i := 0
	for _, serving := range f.servings {
		servings[i] = serving
		i++
	}
	return servings
}
