package translator

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Translate() {
	fmt.Println("Requesting the API server...")
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
