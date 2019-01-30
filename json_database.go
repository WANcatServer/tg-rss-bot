package main

import (
	"encoding/json"
	"io/ioutil"
)

type jsonDB struct {
	path  string `json:"-"`
	Feeds []Feed `json:"Feeds"`
	Users []User `json:"Users"`
}

func NewjsonDB(path string) (db Database, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var jsondb jsonDB
	err = json.Unmarshal(file, &jsondb)
	jsondb.path = path
	return &jsondb, nil
}

func (db *jsonDB) Close() error {
	var data, err = json.Marshal(*db)
	err = ioutil.WriteFile(db.path, data, 644)
	if err != nil {
		return err
	}
	db = nil
	return nil
}

func (db *jsonDB) insertFeed(Url string, Updated Time) (ID, error) {
	// insert a feed into database
	var feed Feed
	feed.Url = Url
	feed.Updated = Updated
	feed.Id = ID(len(db.Feeds))
	db.Feeds = append(db.Feeds, feed)
	return feed.Id, nil
}

func (db *jsonDB) updateFeed(Id ID, time Time) error {
	// update feed's last updated time
	db.Feeds[Id].Updated = time
	return nil
}

func (db *jsonDB) checkFeed(Url string) (Id ID) {
	// checkout a feed id by url
	for _, v := range db.Feeds {
		if v.Url == Url {
			return v.Id
		}
	}
	return -1
}

func (db *jsonDB) getFeed(Id ID) (feed *Feed, err error) {
	// get a feed by id
	return &db.Feeds[Id], nil
}

func (db *jsonDB) removeFeed(Id ID) error {
	// remove a feed
	db.Feeds[Id].Url = ""
	db.Feeds[Id].Updated = ""
	return nil
}

func (db *jsonDB) getSubscriber(feedId ID) ([]User, error) {
	var users []User = make([]User, 0, len(db.Users))
	for _, user := range db.Users {
		for _, f_id := range user.Feeds {
			if f_id == feedId {
				users = append(users, user)
			}
		}
	}
	return users, nil
}

//Users

func (db *jsonDB) insertUser(Url string) (ID, error) {
	// insert a new User data into database
	var user User
	user.Id = ID(len(db.Users))
	user.Url = Url
	user.Feeds = []ID{}
	db.Users = append(db.Users, user)
	return user.Id, nil
}

func (db *jsonDB) checkUser(Url string) ID {
	// checkout a user id by url
	for i, v := range db.Users {
		if v.Url == Url {
			return ID(i)
		}
	}
	return ID(-1)
}

func (db *jsonDB) getUser(Id ID) (*User, error) {
	// get a User data
	return &db.Users[Id], nil
}

func (db *jsonDB) removeUser(Id ID) error {
	// remove a user data
	db.Users[Id].Url = ""
	db.Users[Id].Feeds = nil
	return nil
}

func (db *jsonDB) deleteFeed(userId ID, feedId ID) error {
	// delete a feed in user's list
	var deleted int
	user, _ := db.getUser(userId)
	for i, v := range user.Feeds {
		if v == feedId {
			deleted = i
			break
		}
	}
	db.Users[userId].Feeds = append(user.Feeds[:deleted], user.Feeds[deleted+1:]...)
	return nil
}

func (db *jsonDB) addFeed(userId ID, feedId ID) error {
	// add a new feed in user's list
	db.Users[userId].Feeds = append(db.Users[userId].Feeds, feedId)
	return nil
}
