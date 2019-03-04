package pipeline

import (
	"fmt"
	"os/exec"
	"strings"
)

//screenshot takes an html file and converts it to a png image.
func screenshot(htmlfile string) {

	substrings := strings.Split(htmlfile, "/") //TODO this is platform specific..
	filename := substrings[len(substrings)-1]
	filename = strings.Replace(filename, ".html", ".png", -1)
	filename = "./screenshots/" + filename

	cmd := exec.Command("wkhtmltoimage",
		"--width", "1024",
		"--height", "768",
		"--crop-w", "1024",
		"--crop-h", "768",
		"--disable-javascript",
		"--stop-slow-scripts",
		"--no-debug-javascript",
		"--load-error-handling", "skip",
		"--load-media-error-handling", "skip",
		"--quiet",
		htmlfile, //security risk, must be fixed!
		filename)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("A problem with the conversion of", filename, "due to:", err)
		fmt.Printf("%s\n", stdoutStderr)
	} else {
		fmt.Println("Converted", filename)
	}
}
