package teststore_test

import (
	"testing"

	model "github.com/glebtorbin/stock-rest-api/Internal/app/models"

	"github.com/glebtorbin/stock-rest-api/Internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "qwerty@bk.ru"
	_, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
