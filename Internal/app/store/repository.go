package store

import model "github.com/glebtorbin/stock-rest-api/Internal/app/models"

type UserRepository interface {
	Create(*model.User)
	FindByEmail(string) (*model.User, error)
}
