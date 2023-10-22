package validator

import (
	"testing"

	"github.com/junshintakeda/internet-forum/models"
)

func TestThreadValidator_ThreadValidate_Success(t *testing.T) {
	validator := NewThreadValidator()

	validThread := models.Thread{
		Title: "Valid title",
	}

	err := validator.ThreadValidate(validThread)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestThreadValidator_ThreadValidate_RequiredTitle(t *testing.T) {
	validator := NewThreadValidator()

	emptyTitleThread := models.Thread{
		Title: "",
	}

	err := validator.ThreadValidate(emptyTitleThread)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "title: title is required."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestThreadValidator_ThreadValidate_MaxLengthTitle(t *testing.T) {
	validator := NewThreadValidator()

	longTitleThread := models.Thread{
		Title: "this is a very long title that exceeds the max length of the title. Max length of the title is 30 characters. this is too long!",
	}

	err := validator.ThreadValidate(longTitleThread)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "title: limited max 30 char."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}
