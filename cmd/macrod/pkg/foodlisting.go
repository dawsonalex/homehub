package pkg

import (
	"github.com/google/uuid"
)

type FoodListing struct {
	Id       uuid.UUID
	Name     string
	servings map[uuid.UUID]Serving
}

func NewFoodListing(name string) FoodListing {
	return FoodListing{
		Name:     name,
		servings: map[uuid.UUID]Serving{},
	}
}

func (f *FoodListing) Servings() []Serving {
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

func (f *FoodListing) AddServing(serving Serving) {
	if f.servings == nil {
		f.servings = make(map[uuid.UUID]Serving)
	}

	f.servings[serving.Id] = serving
}
