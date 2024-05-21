package pkg

import "github.com/google/uuid"

type FoodListing struct {
	Id       uuid.UUID
	Name     string
	servings map[string]Serving
}

func NewFoodListing(name string) FoodListing {
	return FoodListing{
		Name:     name,
		servings: map[string]Serving{},
	}
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

func (f FoodListing) AddServing(size string, macros Macros) {
	if f.servings == nil {
		f.servings = make(map[string]Serving)
	}
	f.servings[size] = Serving{
		size:   size,
		macros: macros,
	}
}
