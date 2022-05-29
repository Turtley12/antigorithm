package main

import (
	"fmt"
	"syscall/js"

	"github.com/turtley12/antigorithm/channelselect"

	"github.com/turtley12/antigorithm/view"

	//"fmt"

	"github.com/turtley12/antigorithm/feed"
)

func main() {
	pathname := js.Global().Get("window").Get("location").Get("pathname").String()
	fmt.Println("pathname: " + pathname)
	if pathname == "/userfeed.html" || pathname == "/userfeed" {
		userFeed()
	} else {
		home()
	}
}
func userFeed() {
	var userchannels []string
	userchannels = append(userchannels, "UCSju5G2aFaWMqn-_0YBtq5A")
	view.DisplayFeed(feed.GetUserFeed(userchannels))
}
func home() {
	channelselect.CreateButtons()
}
