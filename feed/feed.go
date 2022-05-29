package feed

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	//"syscall/js"
)

const (
	feedurl      = "https://www.youtube.com/feeds/videos.xml?channel_id=<id>"
	xmltojsonurl = "https://api.factmaven.com/xml-to-json/?xml=<xml>"
)

/* Gets videos from all given channels.
   Than organizes them and returns list
   of organized Feed.Entry. */
func GetUserFeed(channels []string) []Video {
	/* List of every feed in given chanels*/
	var feeds []Feed

	//iterate through all given channel IDs
	for _, id := range channels {
		videos := DownloadRSS(id)
		feeds = append(feeds, videos)
	}

	//iterate through each video and give it a ranking
	var videos_out [][]Video

	for _, channel := range feeds {
		for i, video := range channel.Video {
			// TODO actually rank videos
			videos_out = append([i]videos_out, video)
		}
	}
	var videos_out_1d []Video
	for _, videos := range videos_out {
		for _, video := range videos {
			videos_out_1d = append(videos_out_1d, video)
		}
	}

	return videos_out_1d
}

func DownloadRSS(channelid string) Feed {
	//define empty struct
	var ytfeed RSSFeed

	//create download url
	rssurl := strings.ReplaceAll(feedurl, "<id>", channelid)
	url := strings.ReplaceAll(xmltojsonurl, "<xml>", rssurl)

	//get response
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ytfeed.Feed
	}

	defer response.Body.Close()

	// parse xml into earlier defined struct
	bytes, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
		return ytfeed.Feed
	}

	json.Unmarshal(bytes, &ytfeed)
	return ytfeed.Feed
}

//Struct to represent youtube rss feed
type RSSFeed struct {
	Feed Feed `json:"feed"`
}
type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
type Author struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}
type Content struct {
	URL    string `json:"url"`
	Type   string `json:"type"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
type Thumbnail struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
type StarRating struct {
	Count   string `json:"count"`
	Average string `json:"average"`
	Min     string `json:"min"`
	Max     string `json:"max"`
}
type Statistics struct {
	Views string `json:"views"`
}
type Community struct {
	StarRating StarRating `json:"starRating"`
	Statistics Statistics `json:"statistics"`
}
type Group struct {
	Title       string    `json:"title"`
	Content     Content   `json:"content"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Description string    `json:"description"`
	Community   Community `json:"community"`
}
type Video struct {
	VideoID   string    `json:"videoId"`
	ChannelID string    `json:"channelId"`
	Group     Group     `json:"group"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Link      Link      `json:"link"`
	Author    Author    `json:"author"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
}
type Feed struct {
	ChannelID string    `json:"channelId"`
	Link      []Link    `json:"link"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	Published time.Time `json:"published"`
	Video     []Video   `json:"entry"`
}
