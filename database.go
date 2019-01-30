package main

type Database interface {
	Close() error

	insertFeed(string, Time) (ID, error)
	updateFeed(ID, Time) error
	checkFeed(string) ID
	getFeed(ID) (*Feed, error)
	removeFeed(ID) error
	getSubscriber(ID) ([]User, error)

	insertUser(string) (ID, error)
	checkUser(string) ID
	getUser(ID) (*User, error)
	removeUser(ID) error
	deleteFeed(ID, ID) error
	addFeed(ID, ID) error
}
