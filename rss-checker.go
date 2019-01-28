package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Atom struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Link    string   `xml:"id"`
	Updated string   `xml:"updated"`
	Entry   []Entry  `xml:"entry"`
}

type Entry struct {
	Title     string
	Link      string `xml:"id"`
	Published string
	Summary   string
}

func Request(feed Feed) (body []byte, err error) {
	//抓取資料
	resp, err := http.Get(feed.url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil || err != io.EOF {
		return
	}
	return body, nil
}

func CheckNew(body []byte, feed Feed) (result bool, err error) {
	result = false
	var v Atom
	err = xml.Unmarshal(body, &v)
	log.Printf("update v:\t%s\nfeed:\t%s\n", v.Updated, feed.updated)
	result = v.Updated != feed.updated
	return result, nil
}