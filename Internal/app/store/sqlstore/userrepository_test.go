package sqlstore_test

import (
	"testing"

	model "github.com/glebtorbin/stock-rest-api/Internal/app/models"
	"github.com/glebtorbin/stock-rest-api/Internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "host=localhost user=glebb0 password=3005 dbname=stocks_test sslmode=disable")
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "host=localhost user=glebb0 password=3005 dbname=stocks_test sslmode=disable")
	defer teardown("users")
	s := sqlstore.New(db)
	email := "qwerty@bk.ru"
	_, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
