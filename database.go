package main

import (
	"database/sql"
	//	"github.com/mattn/go-sqlite3"
)

type Database interface {
	insertFeed(Feed) error
	removeFeed(string) error
	getFeed([]string) ([]Feed, error)
	insertUser(User) error
	updateUser(*User) error
	getUser() ([]User, error)
	removeUser(*Feed) error
}

type SQLiteDB struct {
	db sql.DB
}

type JSON_DB struct {
	feeds []Feed
	users []User
}

func (db *JSON_DB) insertFeed(feed Feed) {
	feed.id = Id(len(db.feeds))
	db.feeds = append(db.feeds, feed)
}

func (db *JSON_DB) updateFeed(id Id, time Time) {
	db.feeds[id].updated = time
}

func (db *JSON_DB) get() {

}

func (db *JSON_DB) remove() {

}
