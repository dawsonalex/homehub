// Package repo provides an abstraction for service business logic to
// access the data layer for the service. The type of data layer
// used is inconsequential to calling code and all should follow
// the same interface.
package repo

import (
	"context"
	"fmt"
	"github.com/dawsonalex/homehub/cmd/macrod/pkg"
	"github.com/dawsonalex/homehub/pkg/db/postgres"
	"github.com/dawsonalex/homehub/pkg/db/postgres/gen/macrod/public/model"
	. "github.com/dawsonalex/homehub/pkg/db/postgres/gen/macrod/public/table"
	"github.com/dawsonalex/homehub/pkg/id"
	"github.com/pkg/errors"
)

func init() {
	if err := id.RegisterType("FDLIS", "FOOD_LISTING"); err != nil {
		panic(fmt.Sprintf("already a registered type for %s", "FDLIS"))
	}
}

var _ FoodRepo = &Postgres{}

// Postgres provides an implementation of the repository
// backed by a postgres database.
type Postgres struct {
	postgres.DB
}

func NewPostgres(db postgres.DB) *Postgres {
	return &Postgres{db}
}

func (p *Postgres) AddFoodListing(ctx context.Context, listing pkg.FoodListing) (pkg.FoodListing, error) {
	servings := listing.Servings()
	models := make([]model.FoodListings, len(servings))

	for _, serving := range servings {
		carbs, fats, proteins := serving.Macros()

		models = append(models, model.FoodListings{
			ID:      "", // TODO: Add sequence generation and fetching
			Name:    listing.Name(),
			Carbs:   carbs,
			Fats:    fats,
			Protein: proteins,
		})
	}

	insertStmt := FoodListings.INSERT(
		FoodListings.ID,
		FoodListings.Name,
		FoodListings.Carbs,
		FoodListings.Fats,
		FoodListings.Protein,
	).MODELS(models)

	qry, args := insertStmt.Sql()

	// TODO: error reporting of statement, rows affected, and metrics to ctx.
	_, err := p.ExecContext(ctx, qry, args)
	if err != nil {
		return pkg.FoodListing{}, errors.Wrap(err, "can't insert foodlisting")
	}

	return listing, nil
}

// TODO: Use the created sequences, stub out funcs for other types,
// Use this sequence in the call to AddFoodListing
func nextFoodListingSequence() (int64, error) {
	return 0, errors.New("nextFoodListingSequence: not implemented")
}
