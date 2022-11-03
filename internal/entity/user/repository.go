package user

import (
	"context"
)

type Repository struct {
	//For example add you database connection here
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(ctx context.Context, e *Entity) (*Entity, error) {
	//Do you database insert stuff right here
	return e, nil
}

func (r *Repository) Get(ctx context.Context) (*Entity, error) {
	//Do you database stuff right here
	//Don't forget you have the user id in you context right here ctx.Value("userId")

	e := &Entity{
		UserName:  "doe",
		Email:     "john.doe@johndoe.de",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "das5678",
	}

	return e, nil
}

func (r *Repository) List(ctx context.Context) ([]Items, error) {
	entity1 := &Entity{
		UserName:  "mustermann",
		Email:     "max.mustermann@maxmustermann.de",
		FirstName: "Max",
		LastName:  "Mustermann",
		Password:  "123456",
	}

	entity2 := &Entity{
		UserName:  "doe",
		Email:     "john.doe@johndoe.de",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "das5678",
	}

	items := []Items{
		{
			entity1,
		},
		{
			entity2,
		},
	}

	return items, nil
}
