package storage

import (
	"newsfeed/post"
	"sync"
)

type PostStore interface {
	AddPost(post *post.Post) error
	GetPost(uint64) (*post.Post, error)
}

var defaultStore PostStore
var singleton sync.Once

func InitPostStore(store PostStore) {
	singleton.Do(func() {
		defaultStore = store
	})
}

func GetPostStore() PostStore {
	return defaultStore
}
