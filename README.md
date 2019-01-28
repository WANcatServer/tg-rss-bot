# WANcat RSS Reader 

## Feature
* add / remove rss or atom feed
* import / export OPML
* send new post of feeds
* get feed's last posts
* RSS search engine: try to get RSS feed from a normal URL

> writen in Go  

## Method
|      |     |
| ---- | --- |
| /add | add new feed |
| /remove | remove a feed |
| /import | import a OPML file |
| /export | export feed list to OPML file |
| /get | get a feed's last posts |

# Architecture
* Data Saver: to store all of feeds and users
* RSS checker: get new posts from feeds
* History saver: to take down last post of each feeds
* Post sender: make the posts turn into text
* Feed Search Engine: search a feed from a URL
* TG Wrapper 

structure: 
* Post: to take down a post's title, summary,content, image(if it has)
* User: to remember a user's feed list
* Feed : to take down the feed's URL and last post

## RSS checker:
dependent:  
* HTTP
* Post  
attribute:  

method:  
* GetNewPost(feed list []string) []Post
* get(feed string) Post
* 

