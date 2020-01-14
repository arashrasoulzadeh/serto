package functions

import (
	"github.com/caarlos0/spin"
	"io/ioutil"
	"net/http"
)

/**
send a get request
*/
func GETRequest(url string) string {
	s := spin.New("%s GET request...")
	s.Start()
	defer s.Stop()
	response, err := http.Get(url)
	if err != nil {
		return ""
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return ""
		}
		return string(contents)
	}

}
