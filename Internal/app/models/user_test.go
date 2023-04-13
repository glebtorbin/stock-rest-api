package model_test

import (
	"testing"

	model "github.com/glebtorbin/stock-rest-api/Internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.Validate())
}

func TestUserBeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
