package main

import ()

type Database interface {
	insertFeed(string, Time) (Id, error)
	updateFeed(Id, Time) error
	checkFeed(string) Id
	getFeed(Id) (*Feed, error)
	removeFeed(Id) error
	getSubscriber(Id) ([]User, error)

	insertUser(string) (Id, error)
	checkUser(string) Id
	getUser(Id) (*User, error)
	removeUser(Id) error
	deleteFeed(Id, Id) error
	addFeed(Id, Id) error
}
