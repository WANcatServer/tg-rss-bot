package main

type JSON_DB struct {
	feeds []Feed
	users []User
}

func newJSON_DB(path string) (db Database) {
	feeds := make([]Feed, 0)
	users := make([]User, 0)
	db = JSON_DB{feeds, users}
	return
}

func (db *JSON_DB) insertFeed(url string, updated Time) error {
	// insert a feed into database
	var feed Feed
	feed.url = url
	feed.updated = updated
	feed.id = Id(len(db.feeds))
	db.feeds = append(db.feeds, feed)
	return nil
}

func (db *JSON_DB) updateFeed(id Id, time Time) error {
	// update feed's last updated time
	db.feeds[id].updated = time
	return nil
}

func (db *JSON_DB) checkFeed(url string) (id Id, err error) {
	// checkout a feed id by url
	for _, v := range db.feeds {
		if v.url == url {
			return v.id, nil
		}
	}
	return -1, nil
}

func (db *JSON_DB) getFeed(id Id) (feed Feed, err error) {
	// get a feed by id
	return db.feeds[id], nil
}

func (db *JSON_DB) removeFeed(id Id) error {
	// remove a feed
	db.feeds[id].url = ""
	db.feeds[id].updated = ""
	return nil
}

//Users

func (db *JSON_DB) insertUser(url string) (Id, error) {
	// insert a new User data into database
	var user User
	user.id = Id(len(db.users))
	user.url = url
	user.feeds = []Id{}
	db.users = append(db.users, user)
	return user.id, nil
}

func (db *JSON_DB) checkUser(url string) (Id, error) {
	// checkout a user id by url
	for i, v := range db.users {
		if v.url == url {
			return Id(i), nil
		}
	}
	return Id(-1), nil
}

func (db *JSON_DB) getUser(id Id) (User, error) {
	// get a User data
	return db.users[id], nil
}

func (db *JSON_DB) removeUser(id Id) error {
	// remove a user data
	db.users[id].url = ""
	db.users[id].feeds = nil
	return nil
}

func (db *JSON_DB) deleteFeed(userId Id, feedId Id) error {
	// delete a feed in user's list
	var feeds *[]Id = &db.users[userId].feeds
	var newfeed []Id = make([]Id, len(*feeds)-1)
	var deleted bool = false
	for i, v := range *feeds {
		if v == feedId {
			deleted = true
		}
		if deleted {
			newfeed[i] = v
		} else {
			newfeed[i-1] = v
		}
	}
	*feeds = newfeed
	return nil
}

func (db *JSON_DB) addFeed(userId Id, feedId Id) error {
	// add a new feed in user's list
	db.users[userId].feeds = append(db.users[userId].feeds, feedId)
	return nil
}
