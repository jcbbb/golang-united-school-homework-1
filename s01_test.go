package structs

import (
	"testing"
)

func TestFullName(t *testing.T) {
	user := NewUser("Doe", "John")
	if user.FullName() != "Doe John" {
		t.Errorf("user fullname is incorrect. Got: %v, needed: %v", user.FullName(), "Doe John")
	}
}

func TestIsUserType(t *testing.T) {
	user := NewUser("Doe", "John")
	if ok := isUser(user); !ok {
		t.Errorf("user type is incorrect")
	}
}

func TestResetUser(t *testing.T) {
	user := NewUser("Doe", "John")
	ResetUser(user)
	if user.firstName != "" && user.lastName != "" {
		t.Errorf("Expected user reset, but got: %v", user)
	}
}

func TestProcessUser(t *testing.T) {
	user := NewUser("Doe", "John")
	fullName := user.FullName()
	newFullName := ProcessUser(user)

	if newFullName == fullName {
		t.Errorf("user lastName and firstName was not reset")
	}
}
