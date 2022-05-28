package view

import (
	//"strconv"
	"syscall/js"

	"github.com/turtley12/antigorithm/feed"
)

const scalefactor int = 4

func DisplayFeed(videos []feed.Video) {
	document := js.Global().Get("document")
	container := document.Call("getElementById", "container")

	for _, video_info := range videos {
		video_hyperlink := document.Call("createElement", "a")
		video_hyperlink.Call("setAttribute", "href", video_info.Link.Href)

		video_div := document.Call("createElement", "div")
		video_div.Call("setAttribute", "id", video_info.ID)
		video_div.Call("setAttribute", "class", "video")

		video_thumb := document.Call("createElement", "img")
		video_thumb.Call("setAttribute", "src", video_info.Group.Thumbnail.URL)
		video_div.Call("appendChild", video_thumb)

		video_title := document.Call("createElement", "figcaption")
		text := document.Call("createTextNode", video_info.Group.Title)
		video_title.Call("appendChild", text)
		video_div.Call("appendChild", video_title)

		video_hyperlink.Call("appendChild", video_div)
		container.Call("appendChild", video_hyperlink)
	}

}
