package models

import "time"

type User struct {
	Id             int
	Email          string
	Username       string
	Password       string
	RepeatPassword string
	ExpiresAt      time.Time
	IsAuth         bool
}

type Post struct {
	Id          int
	Title       string
	Description string
	Image       string
	Category    []string
	Author      string
	Likes       int
	Dislikes    int
	CreateAt    time.Time
}

type Message struct {
	Id          int
	PostId      int
	CommentId   int
	Author      string
	ReactAuthor string
	Message     string
	Active      int
	CreateAt    time.Time
}

type Comment struct {
	Id         int
	PostId     int
	Creator    string
	Text       string
	Likes      int
	Dislikes   int
	IsAuth     bool
	Created_at time.Time
}

type Like struct {
	UserID       int
	PostID       int
	Islike       int
	CommentID    int
	CountLike    int
	Countdislike int
}

type Category struct {
	Name string
}
