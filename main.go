package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/turtley12/antigorithm/channelselect"

	"github.com/turtley12/antigorithm/view"

	"github.com/turtley12/antigorithm/feed"
)

func main() {
	c := make(chan struct{}, 0)

	pathname := js.Global().Get("window").Get("location").Get("pathname").String()
	//fmt.Println("pathname: " + pathname)
	if pathname == "/userfeed.html" || pathname == "/userfeed" {
		userFeed()
	} else {
		home()
	}
	<-c
}
func userFeed() {
	userchannels := getUserChannels()

	view.DisplayFeed(feed.GetUserFeed(userchannels))
}
func getUserChannels() []string {
	window := js.Global().Get("window")
	location := window.Get("location")
	search := location.Get("search").String()

	if search == "" {
		location.Call("replace", location.Get("origin"))
	}

	params := strings.ReplaceAll(search, "?channels=", "")
	userchannels := strings.Split(params, ",")
	for _, channel := range userchannels {
		fmt.Println(channel)
	}
	return userchannels
}
func home() {
	js.Global().Set("generate", js.FuncOf(generate))
	channelselect.CreateButtons()
}
func generate(this js.Value, args []js.Value) interface{} {
	channelselect.GetUserSelection()
	return js.ValueOf(nil)
}
