package teststore

import (
	model "github.com/glebtorbin/stock-rest-api/Internal/app/models"
	"github.com/glebtorbin/stock-rest-api/Internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {

}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
