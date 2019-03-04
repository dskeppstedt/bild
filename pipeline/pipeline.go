package pipeline

import "sync"

//Pipeline keeps track the progress of a single screenshot job.
func Pipeline(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	html, err := requestor(url)
	if err != nil {
		return
	}
	screenshot(html)

}
