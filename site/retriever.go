package site

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

type UrlGenerator interface {
	GenerateUrl() string
}

func RetrieveHtml(generator UrlGenerator) (string, error) {
	resp, err := http.Get(generator.GenerateUrl())
	if err != nil {
		//log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	return fmt.Sprintf("%s", bytes), err
}
