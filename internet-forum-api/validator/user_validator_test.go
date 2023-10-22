package validator

import (
	"testing"

	"github.com/junshintakeda/internet-forum/models"
)

func TestUserValidator_UserValidate_Success(t *testing.T) {
	validator := NewUserValidator()

	validUser := models.User{
		Email:    "Valid@email.com",
		Password: "Valid password",
	}

	err := validator.UserValidate(validUser)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestUserValidator_UserValidate_RequiredEmail(t *testing.T) {
	validator := NewUserValidator()

	emptyEmailUser := models.User{
		Email:    "",
		Password: "Valid password",
	}

	err := validator.UserValidate(emptyEmailUser)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "email: email is required."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestUserValidator_UserValidate_MaxLengthEmail(t *testing.T) {
	validator := NewUserValidator()

	longEmailUser := models.User{
		Email:    "veryveryverylonglonglong@email.com",
		Password: "Valid password",
	}

	err := validator.UserValidate(longEmailUser)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "email: limited max 30 char."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestUserValidator_UserValidate_InvalidEmailFormat(t *testing.T) {
	validator := NewUserValidator()

	invalidEmailUser := models.User{
		Email:    "invalid email format",
		Password: "Valid password",
	}

	err := validator.UserValidate(invalidEmailUser)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "email: is not valid email format."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestUserValidator_UserValidate_RequiredPassword(t *testing.T) {
	validator := NewUserValidator()

	emptyPasswordUser := models.User{
		Email:    "Valid@email.com",
		Password: "",
	}

	err := validator.UserValidate(emptyPasswordUser)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "password: password is required."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestUserValidator_UserValidate_MinLengthPassword(t *testing.T) {
	validator := NewUserValidator()

	shortPasswordUser := models.User{
		Email:    "Valid@email.com",
		Password: "short",
	}

	err := validator.UserValidate(shortPasswordUser)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "password: limited min 6 max 30 char."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestUserValidator_UserValidate_MaxLengthPassword(t *testing.T) {
	validator := NewUserValidator()

	longPasswordUser := models.User{
		Email:    "Valid@email.com",
		Password: "very very very long password that exeeds the max length of the password. Max length of the password is 30 characters. This is too long!",
	}

	err := validator.UserValidate(longPasswordUser)
	if err == nil {
		t.Errorf("Expected an error, but got no error")
	}

	expectedErrorMessage := "password: limited min 6 max 30 char."
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}
