package pipeline

import (
	"bild/util"
	"fmt"
	"io/ioutil"
	"net/http"
)
import "strings"

//Requestor fetches html based on a url and stores it in a file.
func requestor(url string) (string, error) {
	fmt.Println("Working on", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to request", url, err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body", err)
		return "", err
	}

	htmlFile := "/tmp/" + strings.Replace(url, "http://", "", -1) + ".html"

	util.WriteFile(htmlFile, body)
	return htmlFile, nil
}
