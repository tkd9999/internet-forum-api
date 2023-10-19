package validator

import (
	"github.com/junshintakeda/internet-forum/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IThreadValidator interface {
	ThreadValidate(thread models.Thread) error
}

type threadValidator struct{}

func NewThreadValidator() IThreadValidator {
	return &threadValidator{}
}

func (tv *threadValidator) ThreadValidate(thread models.Thread) error {
	return validation.ValidateStruct(&thread,
		validation.Field(
			&thread.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
	)
}
