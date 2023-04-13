package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "qwerty@bk.ru",
		Password: "qwerty",
	}
}
