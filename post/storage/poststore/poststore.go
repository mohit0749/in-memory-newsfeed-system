package poststore

import (
	"fmt"
	"newsfeed/post"
	"sync"
)

type postStore struct {
	posts sync.Map
}

func NewPostStore() *postStore {
	return &postStore{sync.Map{}}
}

func (ps *postStore) AddPost(u *post.Post) error {
	ps.posts.Store(u.GetId(), u)
	return nil
}

func (ps *postStore) GetPost(id uint64) (*post.Post, error) {
	value, ok := ps.posts.Load(id)
	if !ok {
		return nil, fmt.Errorf("post %d does not exists", id)
	}
	u, _ := value.(*post.Post)
	return u, nil
}
