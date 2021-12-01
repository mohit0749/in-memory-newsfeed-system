package comment

import (
	"sync/atomic"
	"time"
)

var seqId uint64

type Comment struct {
	id               uint64
	text             string
	userId           uint64
	postId           uint64
	timeStamp        time.Time
	upVote, downVote uint64
}

func (c *Comment) GetId() uint64 {
	return c.id
}

func (c *Comment) Upvote() (bool, error) {
	atomic.AddUint64(&c.upVote, 1)
	return true, nil
}

func (c *Comment) DownVote() (bool, error) {
	atomic.AddUint64(&c.downVote, 1)
	return true, nil
}

func CreateComment(userId, postId uint64, text string) (*Comment, error) {
	atomic.AddUint64(&seqId, 1)
	c := &Comment{
		id:        seqId,
		text:      text,
		userId:    userId,
		postId:    postId,
		timeStamp: time.Now(),
		upVote:    0,
		downVote:  0,
	}
	return c, nil
}
