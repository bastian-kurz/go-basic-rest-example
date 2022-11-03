package user

import (
	"context"
	"errors"
	"github.com/ggwhite/go-masker"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserId int

type UserInterface interface {
	Create(ctx context.Context, entity Entity) error
}

type Response struct {
	Data []Items `json:"data"`
}

type Items struct {
	*Entity
}

type Entity struct {
	UserName  string `json:"userName" validate:"required,alphanum"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,alpha"`
	LastName  string `json:"lastName" validate:"required,alpha"`
	Password  string `json:"password" validate:"required"`
}

func (u *Entity) Bind(r *http.Request) error {
	// u.User is nil if no User fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if u == nil {
		return errors.New("missing required User fields")
	}

	if err := u.validateStruct(); err != nil {
		return err
	}

	return nil
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	for _, u := range res.Data {
		// Pre-processing before a response is marshalled and sent across the wire
		u.Password = masker.Password(u.Password)
		u.Email = masker.Email(u.Email)
		u.FirstName = masker.Name(u.FirstName)
		u.LastName = masker.Name(u.LastName)
	}

	return nil
}

func (u *Entity) validateStruct() error {
	validate := validator.New()

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(u)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}
