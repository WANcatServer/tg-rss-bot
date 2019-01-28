package main

type Id string

type Feed struct {
	id      Id
	url     string
	updated string
}

type Post struct {
	feedId  Id
	title   string
	summary string
	image   string
}

type User struct {
	id    string
	feeds []Id
}
