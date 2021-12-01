package newsfeed

import (
	"newsfeed/comment/storage/commentstore"
	"newsfeed/post/storage/poststore"
	"newsfeed/user/storage/userstore"
	"sort"

	"newsfeed/comment"
	commentstorage "newsfeed/comment/storage"
	"newsfeed/post"
	poststorage "newsfeed/post/storage"
	"newsfeed/user"
	userstorage "newsfeed/user/storage"
)

type NewsFeed struct {
	users    userstorage.UserStore
	posts    poststorage.PostStore
	comments commentstorage.CommentStore
}

func NewNewsFeed() NewsFeed {
	commentstorage.InitCommentStore(commentstore.NewCommentStore())
	poststorage.InitPostStore(poststore.NewPostStore())
	userstorage.InitUserStore(userstore.NewUserStore())
	return NewsFeed{userstorage.GetUserStore(), poststorage.GetPostStore(), commentstorage.GetCommentStore()}
}

func (nf *NewsFeed) SignUp(name string) (uint64, error) {
	u, err := user.CreateUser(name)
	if err != nil {
		return 0, err
	}
	err = nf.users.AddUser(u)
	return u.GetId(), err
}

func (nf *NewsFeed) Login(id uint64) (*user.User, error) {
	return nf.users.GetUser(id)
}

func (nf *NewsFeed) ShowNewsFeed(u *user.User) ([]*post.Post, error) {
	fIds := u.GetFollowee()
	posts := make([]*post.Post, 0)
	for _, uid := range fIds {
		u, err := nf.users.GetUser(uid)
		if err != nil {
			return posts, err
		}
		postIds := u.GetPosts()
		for _, pid := range postIds {
			post, err := nf.posts.GetPost(pid)
			if err != nil {
				return posts, err
			}
			posts = append(posts, post)
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
		post, err := nf.posts.GetPost(pid)
		if err != nil {
			return uposts, err
		}
		uposts = append(uposts, post)
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
	nf.posts.AddPost(p)
	u.StorePostId(p.GetId())
	return p, nil
}

func (nf *NewsFeed) Comment(u *user.User, post *post.Post, text string) (*comment.Comment, error) {
	c, err := comment.CreateComment(u.GetId(), post.GetId(), text)
	if err != nil {
		return c, err
	}
	post.StoreComment(c.GetId())
	nf.comments.AddComment(c)
	return c, nil
}
