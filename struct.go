package main

type Id int
type Time string

type Feed struct {
	id      Id
	url     string
	updated Time
}

type Post struct {
	feedId  Id
	title   string
	summary string
	image   string
}

type User struct {
	id    Id
	url   string
	feeds []Id
}
