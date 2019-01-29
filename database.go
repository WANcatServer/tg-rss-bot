package main

import ()

type Database interface {
	insertFeed(string, Time) error
	updateFeed(Id, Time) error
	checkFeed(string) (Id, error)
	getFeed(Id) (Feed, error)
	removeFeed(Id) error

	insertUser(string) (Id, error)
	checkUser(string) (Id, error)
	getUser(Id) (User, error)
	removeUser(Id) error
	deleteFeed(Id, Id) error
	addFeed(Id, Id) error
}
