package site

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type UrlGenerator interface {
	GenerateUrl() string
}

func RetrieveHtmlFromURL(url string) (string, error) {
	resp, err := http.Get(generator.GenerateUrl())
	if err != nil {
		//log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	return fmt.Sprintf("%s", bytes), err
}

func HandleTaskForHtmlRetrieval(incomingURLs <-chan string, outputHTML chan<- string, outputError chan<- error) {
	defer close(outputHTML)
	for url, ok := <- incomingURLs; ok; {
		if !ok {
			break // channel probably closed
		}
		go func() {
			htmldata, err := RetrieveHtmlFromURL(url)
			if err != nil {
				outputError <- err
			} else {
				outputHTML <- htmldata
			}
		}()
	}
}
