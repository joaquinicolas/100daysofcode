// This is a code sample to show how to avoid go routines leak by using
// buffered channels
// If we had used unbuffered channel, the two slower goroutines would have
// got stuck trying to sent the response on a channel with no receiver.
package leak

import (
	"io/ioutil"
	"net/http"
)

var url = [3]string{
	"https://www.google.com.ar/",
	"https://www.google.co.uk/webhp",
	"https://www.google.com/",
}

// MirroredQuery runs three parallel queries
// returns the quickest response
func MirroredQuery() string {
	responses := make(chan string, 3)
	go doRequest(responses, url[0])
	go doRequest(responses, url[1])
	go doRequest(responses, url[2])

	return <-responses
}

// doRequest returns the downloaded content from a url throught a output channel
func doRequest(data chan<- string, url string) {
	data <- request(url)
}

//
func request(url string) string {
	data, err := http.Get(url)
	if err != nil {
		return "no content"
	}
	defer data.Body.Close()
	resp, _ := ioutil.ReadAll(data.Body)

	return string(resp)
}
