package user

import (
	"sync"
	"sync/atomic"
)

var seqId uint64

type User struct {
	id       uint64
	name     string
	followee sync.Map
	posts    sync.Map
}

func (u *User) StorePostId(pid uint64) {
	u.posts.Store(pid, struct{}{})
}

func (u User) GetId() uint64 {
	return u.id
}

func (u User) GetFollowee() []uint64 {
	f := make([]uint64, 0)
	u.followee.Range(func(key, value interface{}) bool {
		pId, _ := key.(uint64)
		f = append(f, pId)
		return true
	})
	return f
}

func (u User) GetPosts() []uint64 {
	posts := make([]uint64, 0)
	u.posts.Range(func(key, value interface{}) bool {
		pId, _ := key.(uint64)
		posts = append(posts, pId)
		return true
	})
	return posts
}

func (u *User) Follow(userId uint64) (bool, error) {
	u.followee.Store(userId, struct{}{})
	return true, nil
}

func CreateUser(name string) (*User, error) {
	atomic.AddUint64(&seqId, 1)
	return &User{
		id:       seqId,
		name:     name,
		followee: sync.Map{},
		posts:    sync.Map{},
	}, nil
}
