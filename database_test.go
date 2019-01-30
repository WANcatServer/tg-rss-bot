package main

import (
	"testing"
)

var Url string = "https://wancat.cc/atom.xml"
var Updated Time = "2018-12-31T16:00:00.000Z"
var tg string = "t.me/lancatlin"

func TestCRUD(t *testing.T) {
	var db Database
	var err error
	db, err = NewjsonDB(`test-data/db.json`)
	defer db.Close()
	var userId ID = db.checkUser(tg)
	var feedId ID = db.checkFeed(Url)
	// insert Feed Data
	if userId == -1 {
		t.Log("Didn't check user")
		if feedId == -1 {
			t.Log("Didn't check feed")
			feedId, err = db.insertFeed(Url, Updated)
			if err != nil {
				t.Error("insert feed fatal: ", err)
			}
		}
		userId, err = db.insertUser(tg)
		if err != nil {
			t.Error("insert user fatal: ", err)
		}
		err = db.addFeed(userId, feedId)
		if err != nil {
			t.Error("add feed into user's list fatal: ", err)
		}
	}
	// feed updated, notice users
	users, err := db.getSubscriber(feedId)
	if err != nil {
		t.Error("Get Subscriber fatal: ", err)
	}
	t.Log("users: ", users)
	// remove
	t.Logf("userId: %d\tfeedId: %d\n", userId, feedId)
	err = db.deleteFeed(userId, feedId)
	if err != nil {
		t.Error("delete feed fatal")
	}
	users, err = db.getSubscriber(feedId)
	if len(users) != 0 {
		t.Error("db.deleteFeed didn't work", users)
	}
	err = db.removeUser(userId)
	if err != nil {
		t.Error("delete user fatal")
	}
}
