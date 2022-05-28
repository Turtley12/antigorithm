package main

import (
	"github.com/turtley12/antigorithm/view"
	//"fmt"

	"github.com/turtley12/antigorithm/feed"
)

func main() {
	var userchannels []string
	userchannels = append(userchannels, "UCSju5G2aFaWMqn-_0YBtq5A")
	view.DisplayFeed(feed.GetUserFeed(userchannels))

}
