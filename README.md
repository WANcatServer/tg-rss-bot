# WANcat RSS Reader 

## Feature
* add / remove rss or atom feed
* import / export OPML
* send new post of feeds
* get feed's last posts
* RSS search engine: try to get RSS feed from a normal URL

> writen in Go  

## Method
| Method | Description |
| ---- | --- |
| /add | add new feed |
| /remove | remove a feed |
| /import | import a OPML file |
| /export | export feed list to OPML file |
| /get | get a feed's last posts |

# Architecture
* - [ ] Database: to store all of feeds and users
* - [x] RSS checker: get new posts from feeds
* - [ ] Post sender: make the posts turn into text
* - [ ] Feed Search Engine: search a feed from a URL
* - [ ] TG Wrapper 

structure: 
* Post: to take down a post's title, summary,content, image(if it has)
* User: to remember a user's feed list
* Feed : to take down the feed's URL and last post
* Atom: the atom file parse result

## RSS checker:

dependent:  
* HTTP
* Post 
* Feed
* Database

attribute:  

method:  
* CheckNew(body []byte, feed Feed) bool
* Request(feed Feed) []byte
* Parse(body []byte) Atom
* CheckAll(feeds []Feed) []Post

