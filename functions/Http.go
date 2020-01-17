package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/caarlos0/spin"
)

/**
send a get request
*/
func GETRequest(url string) string {
	if !IsJsonOutput() {
		s := spin.New("%s GET request...")
		s.Start()
		defer s.Stop()
	}
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

/**
public ip json struct
*/
type PublicIpStruct struct {
	IP string `json:"ip"`
}

/**
get public ip
*/
func GetPublicIP() PublicIpStruct {
	ip := GETRequest("https://api.ipify.org/?format=json")
	var jsonData PublicIpStruct
	err := json.Unmarshal([]byte(ip), &jsonData)
	if err != nil {
		ErrorAndDie("there was an error getting public ip.")
	}
	return jsonData
}

/**
http static serve on given port
*/
func ServeStaticHttp(port int) {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", fs)

	Verbose(strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
