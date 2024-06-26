package model

import "testing"

func TestUser(*testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
