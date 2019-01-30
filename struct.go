package main

type ID int
type Time string

type Feed struct {
	Id      ID
	Url     string
	Updated Time
}

type Post struct {
	FeedId  ID
	Title   string
	Summary string
	Image   string
}

type User struct {
	Id    ID
	Url   string
	Feeds []ID
}
