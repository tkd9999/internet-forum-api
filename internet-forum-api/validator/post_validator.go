package validator

import (
	"github.com/junshintakeda/internet-forum/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPostValidator interface {
	PostValidate(post models.Post) error
}

type postValidator struct{}

func NewPostValidator() IPostValidator {
	return &postValidator{}
}

func (pv *postValidator) PostValidate(post models.Post) error {
	return validation.ValidateStruct(&post,
		validation.Field(
			&post.Content,
			validation.Required.Error("content is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
		),
	)
}
