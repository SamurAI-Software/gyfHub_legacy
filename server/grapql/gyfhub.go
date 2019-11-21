package grapql

import "time"

type Profile struct {
	Email    string
	Username string
	Mobile   string
	Avatar   string
	UserID   string
}

type Hub struct {
	ID           string
	Name         string
	Logo         string
	Hubers       int
	Status       bool
	Close        bool
	LatestActive time.Time
	UserID       string
}

type Message struct {
	ID       string
	GifID    string
	Gif      string
	CreateAt time.Time
	UserID   string
}

type Follower struct {
	ID          string
	IsFollowing bool
	UserID      string
}
