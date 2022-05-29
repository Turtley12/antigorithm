package channelselect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"syscall/js"
)

var channels map[string]string

func CreateButtons() {
	channels = getChannels()
	//fmt.Println(channels)
	document := js.Global().Get("document")
	selectiondiv := document.Call("getElementById", "selection")
	fieldset := document.Call("createElement", "fieldset")
	fieldset.Call("setAttribute", "id", "fieldset")

	for name, _ := range channels {
		div := document.Call("createElement", "div")
		div.Call("setAttribute", "class", "button_div")
		check := document.Call("createElement", "input")
		check.Call("setAttribute", "type", "checkbox")
		check.Call("setAttribute", "name", "channel_select")
		check.Call("setAttribute", "class", "channel_select")
		check.Call("setAttribute", "value", name)
		check.Call("setAttribute", "id", strings.ToLower(name))

		label := document.Call("createElement", "label")
		label.Set("htmlFor", name)
		label.Call("appendChild", document.Call("createTextNode", name))

		br := document.Call("createElement", "br")

		div.Call("appendChild", check)
		div.Call("appendChild", label)
		div.Call("appendChild", br)
		fieldset.Call("appendChild", div)
	}
	selectiondiv.Call("appendChild", fieldset)
}

func GetUserSelection() {
	var userchannels []string

	document := js.Global().Get("document")
	for name, id := range channels {
		check := document.Call("getElementById", strings.ToLower(name))
		if check.Get("checked").Bool() {
			userchannels = append(userchannels, id)
		}
	}
	url := createLink(userchannels)
	clipboard := js.Global().Get("navigator").Get("clipboard")
	clipboard.Call("writeText", js.ValueOf(url))

}

func createLink(userchannels []string) string {
	window := js.Global().Get("window")
	current_url := window.Get("location").Get("href").String()
	output_url := current_url + "userfeed.html?channels="
	for _, channel := range userchannels {
		output_url = output_url + channel + ","
	}
	return output_url

}

func getChannels() map[string]string {
	url := "channels.json"
	//get response
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bytes, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println(string(bytes))
	var result map[string]string
	err1 := json.Unmarshal(bytes, &result)
	if err1 != nil {
		fmt.Println(err1)
	}
	return result
}
