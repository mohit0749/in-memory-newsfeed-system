package commentstore

import (
	"fmt"
	"newsfeed/comment"
	"sync"
)

type commentStore struct {
	comments sync.Map
}

func NewCommentStore() *commentStore {
	return &commentStore{sync.Map{}}
}

func (cs *commentStore) AddComment(c *comment.Comment) error {
	cs.comments.Store(c.GetId(), c)
	return nil
}

func (cs *commentStore) GetComment(id int64) (*comment.Comment, error) {
	value, ok := cs.comments.Load(id)
	if !ok {
		return nil, fmt.Errorf("comment %d does not exists", id)
	}
	u, _ := value.(*comment.Comment)
	return u, nil
}
