package storage

import (
	"newsfeed/user"
	"sync"
)

type UserStore interface {
	AddUser(user *user.User) error
	GetUser(uint64) (*user.User, error)
}

var defaultStore UserStore
var singleton sync.Once

func InitUserStore(store UserStore) {
	singleton.Do(func() {
		defaultStore = store
	})
}

func GetUserStore() UserStore {
	return defaultStore
}
