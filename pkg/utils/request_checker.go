package utils

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jexroid/gopi/api"
)

func SignupCheker(request api.User) error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Phone, validation.Required, validation.Min(989000000000), validation.Max(989999999999)),
		validation.Field(&request.Firstname, validation.Required, validation.Length(3, 65)),
		validation.Field(&request.Lastname, validation.Required, validation.Length(3, 65)),
		validation.Field(&request.Password, validation.Required, validation.Length(4, 40)),
	)
}

func SigninCheker(request api.LoginRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Phone, validation.Required, validation.Min(989000000000), validation.Max(989999999999)),
		validation.Field(&request.Password, validation.Required, validation.Length(4, 40)),
	)
}

func ValidateChecker(request api.ValidateRequest) error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Token, validation.Required, validation.Length(110, 290)),
	)
}
