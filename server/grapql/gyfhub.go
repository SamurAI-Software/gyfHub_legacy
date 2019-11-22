package grapql

import (
	"time"
)

type Hub struct {
	ID           string
	Logo         string
	Hubers       int
	Status       bool
	Close        bool
	LatestActive time.Time
	Username     string
}

type Message struct {
	ID       string
	GifID    string
	Gif      string
	CreateAt time.Time
	Username string
}

type Follower struct {
	Username    string
	Isfollowing bool
}
