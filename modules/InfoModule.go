package modules

import (
	"encoding/json"
	"github.com/arashrasoulzadeh/serto.git/functions"
	"os"
)

/**
init the http module
*/
func ParseInfoModule(args []string) {
	functions.SetProcessor("Info")
	has_command := functions.DieIfEqual(2, args, "Please Enter Sub Command.")
	if has_command {
		command := args[2]
		switch command {
		case "ip":
			MyPublicIp()
			break
		}

	}
}

/**
public ip json struct
*/
type myPublicIpStruct struct {
	IP string `json:"ip"`
}

/**
show your public ip
*/
func MyPublicIp() {
	args := os.Args
	geo := functions.IpGeolocationStruct{}
	if len(args) == 3 {
		ip := functions.GETRequest("https://api.ipify.org/?format=json")
		var jsonData myPublicIpStruct
		err := json.Unmarshal([]byte(ip), &jsonData)

		if err != nil {
			functions.ErrorAndDie("there was an error getting data.")
		}
		functions.Verbose("Your Public ip is " + jsonData.IP)
		geo = functions.IpGeolocation(jsonData.IP)
	} else {
		ip := args[3]
		geo = functions.IpGeolocation(ip)
	}
	functions.Verbose("Country : " + geo.COUNTRY_NAME)
	functions.Verbose("Code    : " + geo.COUNTRY_CODE)
	functions.Verbose("Lat     : " + geo.LAT)
	functions.Verbose("Lon     : " + geo.LON)
}
