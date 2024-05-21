-- +goose Up
-- +goose StatementBegin

-- food_listings stores an item of food, e.g. an Apple
create table food_listings (
    id uuid primary key default gen_random_uuid(),
    name text not null
);

-- servings stores nutrition info about a food_listing, e.g. an 100g of Apple has 10g carbs etc.
create table servings (
    id uuid primary key default gen_random_uuid(),
    food_listing uuid references food_listings(id) on delete cascade,
    name text not null,
    carbs float not null, -- per 100g
    fats float not null, -- per 100g
    protein float not null, -- per 100g

    -- unique constraint to prevent inserting duplicate servings for one food
    constraint food_listing_serving_name_constraint unique (food_listing, name)
);

-- meals stores data about a collection of servings of food_listings
create table meals (
    id uuid primary key default gen_random_uuid(),
    name text not null
);

-- food_entries stores the serving information for food in a meal
create table food_entries (
    id uuid primary key default gen_random_uuid(),
    food_listing uuid not null,
    serving_name text not null,
    meal_id uuid references meals(id) on delete cascade,
    foreign key (food_listing, serving_name) references servings(food_listing, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table food_entries;
drop table meals;
drop table servings;
drop table food_listings;
-- +goose StatementEnd