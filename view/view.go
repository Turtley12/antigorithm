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
		video_div := document.Call("createElement", "div")
		video_div.Call("setAttribute", "id", video_info.ID)

		video_thumb := document.Call("createElement", "img")
		video_thumb.Call("setAttribute", "src", video_info.Group.Thumbnail.URL)

		video_thumb.Call("setAttribute", "width", getThumbWidth(video_info))
		video_thumb.Call("setAttribute", "height", getThumbWidth(video_info))
		video_div.Call("appendChild", video_thumb)

		container.Call("appendChild", video_div)
	}

}

func getThumbHeight(video feed.Video) string {
	//original_height, _ := strconv.Atoi(video.Group.Thumbnail.Height)

	return "auto" //strconv.Itoa(original_height / scalefactor)
}
func getThumbWidth(video feed.Video) string {
	//original_width, _ := strconv.Atoi(video.Group.Thumbnail.Width)
	return "50%"
	//return strconv.Itoa(original_width / scalefactor)
}
