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
	Updated Time     `xml:"updated"`
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
	resp, err := http.Get(feed.Url)
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

func Parse(body []byte) Atom {
	var v Atom
	xml.Unmarshal(body, &v)
	return v
}

func CheckNew(body []byte, feed Feed) (bool, error) {
	//檢查是否有新文章
	var v Atom = Parse(body)
	log.Printf("update v:\t%s\nfeed:\t%s\n", v.Updated, feed.Updated)
	return v.Updated != feed.Updated, nil
}
