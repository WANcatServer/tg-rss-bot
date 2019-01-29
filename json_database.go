package main

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
	feed.id = Id(len(db.feeds))
	db.feeds = append(db.feeds, feed)
	return feed.id, nil
}

func (db *jsonDB) updateFeed(id Id, time Time) error {
	// update feed's last updated time
	db.feeds[id].updated = time
	return nil
}

func (db *jsonDB) checkFeed(url string) (id Id) {
	// checkout a feed id by url
	for _, v := range db.feeds {
		if v.url == url {
			return v.id
		}
	}
	return -1
}

func (db *jsonDB) getFeed(id Id) (feed *Feed, err error) {
	// get a feed by id
	return &db.feeds[id], nil
}

func (db *jsonDB) removeFeed(id Id) error {
	// remove a feed
	db.feeds[id].url = ""
	db.feeds[id].updated = ""
	return nil
}

func (db *jsonDB) getSubscriber(feedId Id) ([]User, error) {
	var users []User = make([]User, 0, len(db.users))
	for _, user := range db.users {
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
	user.id = Id(len(db.users))
	user.url = url
	user.feeds = []Id{}
	db.users = append(db.users, user)
	return user.id, nil
}

func (db *jsonDB) checkUser(url string) Id {
	// checkout a user id by url
	for i, v := range db.users {
		if v.url == url {
			return Id(i)
		}
	}
	return Id(-1)
}

func (db *jsonDB) getUser(id Id) (*User, error) {
	// get a User data
	return &db.users[id], nil
}

func (db *jsonDB) removeUser(id Id) error {
	// remove a user data
	db.users[id].url = ""
	db.users[id].feeds = nil
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
	db.users[userId].feeds = append(user.feeds[:deleted], user.feeds[deleted+1:]...)
	return nil
}

func (db *jsonDB) addFeed(userId Id, feedId Id) error {
	// add a new feed in user's list
	db.users[userId].feeds = append(db.users[userId].feeds, feedId)
	return nil
}
