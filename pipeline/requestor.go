package pipeline

import (
	"bild/util"
	"fmt"
	"io/ioutil"
	"net/http"
)
import "strings"

//Requestor fetches html based on a url and stores it in a file.
func requestor(url string) {
	fmt.Println("Working on", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Fail to request", url)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body", err)
		return
	}

	name := strings.Replace(url, "http://", "", -1)

	util.WriteFile("./tmp/"+name+".html", body)
}
