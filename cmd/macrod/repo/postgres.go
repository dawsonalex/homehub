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
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
)

var _ FoodRepo = &Postgres{}

// Postgres provides an implementation of the repository
// backed by a postgres database.
type Postgres struct {
	postgres.DB
}

func (p *Postgres) AddMeal(ctx context.Context, meal pkg.Meal) (pkg.Meal, error) {
	insertStmt := Meals.INSERT(
		Meals.Name,
	).MODEL(model.Meals{
		Name: meal.Name,
	}).RETURNING(Meals.AllColumns)

	var mealModel model.Meals
	err := insertStmt.Query(p.DB, &mealModel)
	if err != nil {
		return pkg.Meal{}, errors.Wrap(err, "error reading returned meal row")
	}

	_, err = p.insertFoodEntry(ctx, &meal.Id, meal.GetFood()...)
	if err != nil {
		return pkg.Meal{}, err
	}

	meal.Id = mealModel.ID
	return meal, nil
}

func (p *Postgres) insertFoodEntry(ctx context.Context, mealId *uuid.UUID, entries ...pkg.FoodEntry) ([]model.FoodEntries, error) {
	models := make([]model.FoodEntries, len(entries))
	for i, entry := range entries {
		models[i] = model.FoodEntries{
			FoodListing: entry.Id,
			ServingName: entry.SelectedServing().Size(),
			MealID:      mealId,
		}
	}

	insertStmt := FoodEntries.INSERT(
		FoodEntries.FoodListing,
		FoodEntries.ServingName,
		FoodEntries.MealID,
	).MODELS(models).RETURNING(FoodEntries.AllColumns)

	resultModels := make([]model.FoodEntries, len(entries))
	err := insertStmt.QueryContext(ctx, p.DB, &resultModels)
	if err != nil {
		return []model.FoodEntries{}, errors.Wrap(err, "can't insert serving")
	}
	return resultModels, nil
}

func NewPostgres(db postgres.DB) *Postgres {
	return &Postgres{db}
}

func (p *Postgres) AddFoodListing(ctx context.Context, listing pkg.FoodListing) (pkg.FoodListing, error) {
	// TODO: Make sure this is run a single transaction to avoid partially committing a food listing + servings.
	insertStmt := FoodListings.INSERT(
		FoodListings.Name,
	).MODEL(model.FoodListings{
		Name: listing.Name,
	}).RETURNING(FoodListings.AllColumns)

	var foodListingModel model.FoodListings
	err := insertStmt.Query(p.DB, &foodListingModel)
	if err != nil {
		return pkg.FoodListing{}, errors.Wrap(err, "error reading returned food listing row")
	}

	modelServings := []model.Servings{}
	if servings := listing.Servings(); len(servings) > 0 {
		modelServings, err = p.AddServing(ctx, &foodListingModel.ID, servings)
		if err != nil {
			return pkg.FoodListing{}, err
		}
	}

	newListing := pkg.FoodListing{
		Id:   foodListingModel.ID,
		Name: foodListingModel.Name,
	}
	for _, serving := range modelServings {
		newListing.AddServing(pkg.NewServing(serving.Name, pkg.NewMacros(serving.Carbs, serving.Fats, serving.Protein)))
	}

	return newListing, nil
}

func (p *Postgres) AddServing(ctx context.Context, foodListingId *uuid.UUID, servings []pkg.Serving) ([]model.Servings, error) {
	log.Printf("adding servings %d: %v", len(servings), servings)
	models := make([]model.Servings, len(servings))
	for i, serving := range servings {
		carbs, fats, proteins := serving.Macros()

		models[i] = model.Servings{
			Name:        serving.Size(),
			FoodListing: foodListingId,
			Carbs:       carbs,
			Fats:        fats,
			Protein:     proteins,
		}
	}
	log.Printf("adding models %d: %v", len(models), models)

	insertStmt := Servings.INSERT(
		Servings.Name,
		Servings.FoodListing,
		Servings.Carbs,
		Servings.Fats,
		Servings.Protein,
	).MODELS(models).RETURNING(Servings.AllColumns)

	resultModels := make([]model.Servings, len(servings))
	// TODO: error reporting of statement, rows affected, and metrics to ctx.
	query, args := insertStmt.Sql()
	fmt.Printf("query: %s, args: %+v\n", query, args)
	err := insertStmt.QueryContext(ctx, p.DB, &resultModels)
	if err != nil {
		return []model.Servings{}, errors.Wrap(err, "can't insert serving")
	}

	//servings := make([]model.Servings, len(models))
	//// TODO: Update repo stuff so that it returns pkg level components. The model stuff should be contained to this layer.
	//for i, serving := range models {
	//	servings[i] = pkgjhfl
	//}

	return resultModels, nil
}
