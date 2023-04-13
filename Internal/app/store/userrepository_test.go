package store_test

import (
	"testing"

	model "github.com/glebtorbin/stock-rest-api/Internal/app/models"
	"github.com/glebtorbin/stock-rest-api/Internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, "host=localhost user=glebb0 password=3005 dbname=stocks_test sslmode=disable")
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "testmail@test.go",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, "host=localhost user=glebb0 password=3005 dbname=stocks_test sslmode=disable")
	defer teardown("users")
	email := "testmail@test.go"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Email: "testmail@test.go",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
