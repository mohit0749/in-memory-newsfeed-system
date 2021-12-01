package userstore

import (
	"fmt"
	"newsfeed/user"
	"sync"
)

type userStore struct {
	users sync.Map
}

func NewUserStore() *userStore {
	return &userStore{sync.Map{}}
}

func (us *userStore) AddUser(u *user.User) error {
	us.users.Store(u.GetId(), u)
	return nil
}

func (us *userStore) GetUser(id uint64) (*user.User, error) {
	value, ok := us.users.Load(id)
	if !ok {
		return nil, fmt.Errorf("user %d does not exists", id)
	}
	u, _ := value.(*user.User)
	return u, nil
}
