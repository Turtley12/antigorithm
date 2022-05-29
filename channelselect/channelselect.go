package channelselect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	//"syscall/js"
)

func CreateButtons() {
	channels := getChannels()
	fmt.Println(channels)
	/*selectiondiv := js.Global().Get("document").Call("findElementById", "selection")
	document := js.Global().Get("document")
	for channel, _ := ke*/
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
