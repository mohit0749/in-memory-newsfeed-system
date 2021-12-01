package newsfeed

import (
	"fmt"
	"newsfeed/comment"
	"newsfeed/post"
	"newsfeed/user"
	"sort"
	"sync"
)

type NewsFeed struct {
	users    sync.Map
	posts    sync.Map
	comments sync.Map
}

func NewNewsFeed() NewsFeed {
	return NewsFeed{}
}

func (nf *NewsFeed) SignUp(name string) (uint64, error) {
	u, err := user.CreateUser(name)
	if err != nil {
		return 0, err
	}
	nf.users.Store(u.GetId(), u)
	return u.GetId(), nil
}

func (nf *NewsFeed) Login(id uint64) (*user.User, error) {
	value, ok := nf.users.Load(id)
	if !ok {
		return nil, fmt.Errorf("user %d does not exists", id)
	}
	u, _ := value.(*user.User)
	return u, nil
}

func (nf NewsFeed) getUser(uid uint64) *user.User {
	value, _ := nf.users.Load(uid)
	u, _ := value.(*user.User)
	return u
}

func (nf NewsFeed) getPost(uid uint64) *post.Post {
	value, _ := nf.posts.Load(uid)
	p, _ := value.(*post.Post)
	return p
}

func (nf *NewsFeed) ShowNewsFeed(u *user.User) ([]*post.Post, error) {
	fIds := u.GetFollowee()
	posts := make([]*post.Post, 0)
	for _, uid := range fIds {
		u := nf.getUser(uid)
		postIds := u.GetPosts()
		for _, pid := range postIds {
			posts = append(posts, nf.getPost(pid))
		}
	}
	sort.Slice(posts, func(i, j int) bool {
		voteCntI := posts[i].GetVotes()
		voteCntJ := posts[j].GetVotes()
		if voteCntI == voteCntJ {
			return posts[i].GetTime() < posts[j].GetTime()
		}
		if posts[i].GetCommentsCnt() > posts[j].GetCommentsCnt() {
			return true
		}
		return voteCntI > voteCntJ
	})
	userPosts := u.GetPosts()
	uposts := make([]*post.Post, 0)
	for _, pid := range userPosts {
		uposts = append(uposts, nf.getPost(pid))
	}
	sort.Slice(uposts, func(i, j int) bool {
		voteCntI := uposts[i].GetVotes()
		voteCntJ := uposts[j].GetVotes()
		if voteCntI == voteCntJ {
			return uposts[i].GetTime() < uposts[j].GetTime()
		}
		if posts[i].GetCommentsCnt() > posts[j].GetCommentsCnt() {
			return true
		}
		return voteCntI > voteCntJ
	})
	return append(posts, uposts...), nil
}


func (nf *NewsFeed) Post(u *user.User, text string) (*post.Post, error) {
	p, err := post.CreatePost(u.GetId(), text)
	if err != nil {
		return nil, err
	}
	nf.posts.Store(p.GetId(), p)
	u.StorePostId(p.GetId())
	return p, nil
}

func (nf *NewsFeed) Comment(u *user.User, post *post.Post, text string) (*comment.Comment, error) {
	c, err := comment.CreateComment(u.GetId(), post.GetId(), text)
	if err != nil {
		return c, err
	}
	post.StoreComment(c.GetId())
	nf.comments.Store(c.GetId(), c)
	return c, nil
}
