package main

import (
	"io/ioutil"
	"log"
	"testing"
)

var testFeed Feed = Feed{0, "http://www.wancat.cc/atom.xml", "2018-12-31T16:00:00.000Z"}

func TestRequest(t *testing.T) {
	body, err := Request(testFeed)
	if err != nil {
		t.Error("Request Error: ", err)
	}
	log.Printf("%s\n", body[0:200])
}

func TestCheck(t *testing.T) {
	body, err := ioutil.ReadFile("test-data/atom.xml")
	log.Printf("body: %s\n", body[0:50])
	hasNew, err := CheckNew(body, testFeed)
	if err != nil {
		t.Error(err)
	}
	if !hasNew {
		t.Error("Check fail: should be 'true' but except 'false'")
	}
	testFeed.Updated = "2019-01-20T05:23:25.070Z"
	hasNew, err = CheckNew(body, testFeed)
	if err != nil {
		t.Error(err)
	}
	if hasNew {

		t.Error("Check fail: should be 'false' but except 'ture'")
	}
}
