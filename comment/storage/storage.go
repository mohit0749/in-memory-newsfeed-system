package storage

import (
	"newsfeed/comment"
	"sync"
)

type CommentStore interface {
	AddComment(post *comment.Comment) error
	GetComment(int64) (*comment.Comment, error)
}

var defaultStore CommentStore
var singleton sync.Once

func InitCommentStore(store CommentStore) {
	singleton.Do(func() {
		defaultStore = store
	})
}

func GetCommentStore() CommentStore {
	return defaultStore
}
