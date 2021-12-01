package main

import (
	"fmt"

	"newsfeed/newsfeed"
)

func main() {
	newsFeed := newsfeed.NewNewsFeed()
	uid, _ := newsFeed.SignUp("mohit")
	user, _ := newsFeed.Login(uid)
	p, _ := newsFeed.Post(user, "hello world")
	p.Upvote()
	newsFeed.Comment(user, p, "LOL!")

	uid, _ = newsFeed.SignUp("mohit1")
	user, _ = newsFeed.Login(uid)
	p, _ = newsFeed.Post(user, "hello world1")
	p.DownVote()

	uid, _ = newsFeed.SignUp("mohit2")
	user, _ = newsFeed.Login(uid)
	newsFeed.Post(user, "hello world2")

	uid, _ = newsFeed.SignUp("mohit3")
	user, _ = newsFeed.Login(uid)
	newsFeed.Post(user, "hello world3")
	user.Follow(1)
	user.Follow(2)
	user.Follow(3)
	feeds, _ := newsFeed.ShowNewsFeed(user)
	for _, f := range feeds {
		fmt.Println(f.GetId())
		fmt.Println(f.GetText())
		fmt.Println("upvote: ", f.GetUpVotes(), ", Downvotes: ", f.GetDownVotes())
		fmt.Println("Comments: ", f.GetCommentsCnt())
		fmt.Println("posted by: ", f.GetUserId())
	}

}
