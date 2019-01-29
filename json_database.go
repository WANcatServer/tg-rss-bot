package main

import (
	"encoding/json"
	"io/ioutil"
)

type jsonDB struct {
	feeds []Feed
	users []User
}

func newjsonDB(path string) (db Database) {
	feeds := make([]Feed, 0)
	users := make([]User, 0)
	return &jsonDB{feeds, users}
}

func (db *jsonDB) insertFeed(url string, updated Time) (Id, error) {
	// insert a feed into database
	var feed Feed
	feed.url = url
	feed.updated = updated
	feed.id = Id(len(db.Feeds))
	db.Feeds = append(db.Feeds, feed)
	return feed.id, nil
}

func (db *jsonDB) updateFeed(id Id, time Time) error {
	// update feed's last updated time
	db.Feeds[id].updated = time
	return nil
}

func (db *jsonDB) checkFeed(url string) (id Id) {
	// checkout a feed id by url
	for _, v := range db.Feeds {
		if v.url == url {
			return v.id
		}
	}
	return -1
}

func (db *jsonDB) getFeed(id Id) (feed *Feed, err error) {
	// get a feed by id
	return &db.Feeds[id], nil
}

func (db *jsonDB) removeFeed(id Id) error {
	// remove a feed
	db.Feeds[id].url = ""
	db.Feeds[id].updated = ""
	return nil
}

func (db *jsonDB) getSubscriber(feedId Id) ([]User, error) {
	var users []User = make([]User, 0, len(db.Users))
	for _, user := range db.Users {
		for _, f_id := range user.feeds {
			if f_id == feedId {
				users = append(users, user)
			}
		}
	}
	return users, nil
}

//Users

func (db *jsonDB) insertUser(url string) (Id, error) {
	// insert a new User data into database
	var user User
	user.id = Id(len(db.Users))
	user.url = url
	user.feeds = []Id{}
	db.Users = append(db.Users, user)
	return user.id, nil
}

func (db *jsonDB) checkUser(url string) Id {
	// checkout a user id by url
	for i, v := range db.Users {
		if v.url == url {
			return Id(i)
		}
	}
	return Id(-1)
}

func (db *jsonDB) getUser(id Id) (*User, error) {
	// get a User data
	return &db.Users[id], nil
}

func (db *jsonDB) removeUser(id Id) error {
	// remove a user data
	db.Users[id].url = ""
	db.Users[id].feeds = nil
	return nil
}

func (db *jsonDB) deleteFeed(userId Id, feedId Id) error {
	// delete a feed in user's list
	var deleted int
	user, _ := db.getUser(userId)
	for i, v := range user.feeds {
		if v == feedId {
			deleted = i
			break
		}
	}
	db.Users[userId].feeds = append(user.feeds[:deleted], user.feeds[deleted+1:]...)
	return nil
}

func (db *jsonDB) addFeed(userId Id, feedId Id) error {
	// add a new feed in user's list
	db.Users[userId].feeds = append(db.Users[userId].feeds, feedId)
	return nil
}
