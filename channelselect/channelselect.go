package channelselect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

func CreateButtons() {
	channels := getChannels()
	fmt.Println(channels)
	document := js.Global().Get("document")
	selectiondiv := document.Call("getElementById", "selection")
	fieldset := document.Call("createElement", "fieldset")

	for name, _ := range channels {
		check := document.Call("createElement", "input")
		check.Call("setAttribute", "type", "radio")
		check.Call("setAttribute", "name", "channel_slect")
		check.Call("setAttribute", "value", name)
		text := document.Call("createTextNode", name)
		check.Call("appendChild", text)
		fieldset.Call("appendChild", check)
	}
	selectiondiv.Call("appendChild", fieldset)
}

func GetUserSelection() {

}

func getChannels() map[string]interface{} {
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
	fmt.Println(string(bytes))
	var result map[string]interface{}
	err1 := json.Unmarshal(bytes, &result)
	if err1 != nil {
		fmt.Println(err1)
	}
	return result
}
