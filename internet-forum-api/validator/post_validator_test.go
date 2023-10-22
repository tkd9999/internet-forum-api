package validator

import (
	"testing"

	"github.com/junshintakeda/internet-forum/models"
)

func TestPostValidator_PostValidate_Success(t *testing.T) {
	validator := NewPostValidator()

	validPost := models.Post{
		Content: "Valid content",
	}

	err := validator.PostValidate(validPost)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestPostValidator_PostValidate_RequiredContent(t *testing.T) {
	validator := NewPostValidator()

	emptyContentPost := models.Post{
		Content: "",
	}

	err := validator.PostValidate(emptyContentPost)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "content: content is required."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestPostValidator_PostValidate_MaxLengthContent(t *testing.T) {
	validator := NewPostValidator()

	longContentPost := models.Post{
		Content: "this is a very long content that exceeds the max length of the content. Max length of the content is 100 characters. this is too long!",
	}

	err := validator.PostValidate(longContentPost)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "content: limited max 100 char."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}
