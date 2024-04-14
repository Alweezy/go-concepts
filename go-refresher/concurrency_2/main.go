package main

import (
	"fmt"
	"time"
)

var (
	defaultTags = []string{"SystemUser", "User", "NewUser", "System"}
)

type User struct {
	Id, Name, LastName, Status string
	Tags                       []*Tag
}

type Tag struct {
	Name, Type string
}

type Post struct {
	Title  string
	Status string
	UserId string
}

func main() {
	blocking()
}

/*
Main goroutine will be blocked till second goroutine
sends a message letting the main goroutine know that it has finished its work
and so the main goroutine can continue
*/

func blocking() {
	user := &User{}
	done := make(chan bool) // unbuffered channel

	go func() {
		fmt.Println("[Second-GoRoutine] Building User ...")
		buildUser(user)
		fmt.Println("[Second-GoRoutine] Finished building user ...")
		done <- true

		fmt.Println("[Second-Goroutine] Setting default user tags ...")
		setDefaultTags(user)
	}()

	fmt.Println("[Main-GoRoutine] Importing Posts ...")
	posts := importPosts()
	fmt.Println("[Main-GoRoutine] Finished importing Posts!")
	fmt.Println("[Main-GoRoutine] ------Waiting----")
	<-done

	mergeUserPosts(user, posts)
	fmt.Println("Done!")
	fmt.Printf("User %v\n", user)

	for _, post := range posts {
		fmt.Printf("Post %v\n", post)
	}
}

func mergeUserPosts(user *User, posts []*Post) {
	fmt.Println("[Main-GoRoutine] Merging user posts ...")
	for _, post := range posts {
		post.UserId = user.Id
	}
	fmt.Println("[Main-GoRoutine] Finished merging user posts!")
}

func importPosts() []*Post {
	time.Sleep(1 * time.Second)
	titles := []string{"Post 1", "Random Post", "Second Post"}
	var posts []*Post

	for _, title := range titles {
		posts = append(posts, &Post{Title: title, Status: "draft"})
	}

	return posts
}

func buildUser(user *User) {
	time.Sleep(2 * time.Second)
	user.Name = "John"
	user.LastName = "Ale"
	user.Status = "active"
	user.Id = "1"
}

func setDefaultTags(user *User) {
	time.Sleep(1 * time.Second)
	for _, tagName := range defaultTags {
		user.Tags = append(user.Tags, &Tag{Name: tagName, Type: "System"})
	}
}
