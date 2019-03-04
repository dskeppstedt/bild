package pipeline

import "sync"

//Pipeline keeps track the progress of a single screenshot job.
func Pipeline(url string, wg *sync.WaitGroup) {
	requestor(url)

	wg.Done()
}
