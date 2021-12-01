package post

import (
	"sync"
	"sync/atomic"
	"time"
)

var seqId uint64

type Post struct {
	id               uint64
	text             string
	userId           uint64
	timeStamp        time.Time
	upVote, downVote uint64
	comments         sync.Map
	commentsCnt      uint64
}

func (p *Post) StoreComment(cid uint64) {
	p.comments.Store(cid, struct{}{})
	atomic.AddUint64(&p.commentsCnt, 1)
}

func (p Post) GetCommentsCnt() uint64 {
	return p.commentsCnt
}

func (p Post) GetCommentIds() []uint64 {
	c := make([]uint64, 0)
	p.comments.Range(func(key, value interface{}) bool {
		cid, _ := key.(uint64)
		c = append(c, cid)
		return true
	})
	return c
}

func (p Post) GetVotes() int64 {
	return int64(p.upVote - p.downVote)
}

func (p Post) GetUpVotes() uint64 {
	return p.upVote
}

func (p Post) GetDownVotes() uint64 {
	return p.downVote
}

func (p Post) GetText() string {
	return p.text
}

func (p Post) GetUserId() uint64 {
	return p.userId
}

func (p Post) GetTime() int64 {
	return p.timeStamp.Unix()
}

func (p Post) GetId() uint64 {
	return p.id
}

func (p *Post) Upvote() (bool, error) {
	atomic.AddUint64(&p.upVote, 1)
	return true, nil
}

func (p *Post) DownVote() (bool, error) {
	atomic.AddUint64(&p.downVote, 1)
	return true, nil
}

func CreatePost(userid uint64, text string) (*Post, error) {
	atomic.AddUint64(&seqId, 1)
	c := sync.Map{}
	p := &Post{seqId, text, userid, time.Now(), 0, 0, c, 0}
	return p, nil
}
