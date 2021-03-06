package functions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"github.com/phayes/freeport"
)

var processor = "serto"
var clear map[string]func() //create a map for storing clear funcs

/**
verbose output to stdout
*/
func Verbose(msg interface{}) {
	color.Green("%s => %s", processor, msg)
}

/**
grep output of stdout
*/
func GrepOutput(msg string, filter string) {
	scanner := bufio.NewScanner(strings.NewReader(msg))
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, filter) {
			Verbose(scanner.Text())
		}
	}
}

/**
show an error
*/
func Error(err string) {
	color.Red("%s => %s", processor, err)
}

/**
show an error and exit
*/
func ErrorAndDie(err string) {
	Error(err)
	os.Exit(0)
}

/**
die with a error message if array is equal to given size
*/
func DieIfEqual(size int, arr []string, err string) bool {
	if len(arr) == size {
		ErrorAndDie(err)
		return false
	}
	return true
}

/**
set processor label
*/
func SetProcessor(proc string) {
	processor = proc
}

/**
check if file exists
*/
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
execute a command in os
*/
func ExecuteCommand(name string, arg string, arg2 string, arg3 string, standalone bool) []byte {

	if standalone {
		cmd := exec.Command("cmd", "/C", name, arg)
		cmd.Start()
	} else {
		cmd := exec.Command(name, arg, arg2, arg3)
		output, err := cmd.Output()

		if err != nil {
			ErrorAndDie(err.Error())
			return nil
		}
		return output
	}

	return nil
}

/**
open a file with editor
*/
func OpenFileForEdit(path string) {
	ExecuteCommand("nano", path, "", "", true)
}

/**
check if a tcp port is open or not
*/
func CheckPort(port string) bool {

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		Error(err.Error())
		return false
	}

	ln.Close()
	return true
}

/**
get a free port (struct)
*/
type FreePostStruct struct {
	Port int
}

/**
get a free tcp port
*/
func FreePort() FreePostStruct {

	free_port := FreePostStruct{}
	port, err := freeport.GetFreePort()
	if err != nil {
		free_port.Port = 0
		return free_port
	}
	// port is ready to listen on
	free_port.Port = port
	return free_port
}

type IpGeolocationStruct struct {
	COUNTRY_NAME string `json:"geoplugin_countryName"`
	COUNTRY_CODE string `json:"geoplugin_countryCode"`
	LAT          string `json:"geoplugin_latitude"`
	LON          string `json:"geoplugin_longitude"`
	IP           string `json:"geoplugin_request"`
}

/**
get an ip geo location data
*/
func IpGeolocation(ip_address string) IpGeolocationStruct {
	ip := GETRequest("http://www.geoplugin.net/json.gp?ip=" + ip_address)
	var jsonData IpGeolocationStruct
	err := json.Unmarshal([]byte(ip), &jsonData)

	if err != nil {
		ErrorAndDie("there was an error getting data.")
	}
	return jsonData
}

/**
call clear terminal
*/
func CallClear() {
	// Clears the screen
	screen.Clear()
	screen.MoveTopLeft()
}

// pretty print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

/**
is json flag added
*/
func IsJsonOutput() bool {
	args := os.Args
	return Contains(args, "--json")
}

/**
get argument if set and default string if not
*/
func GetArgOrDefault(index int, default_value string) string {
	args := os.Args
	if len(args) >= index {
		value := args[index-1]
		if strings.HasPrefix(value, "--") {
			if len(args) > index {
				return GetArgOrDefault(index+1, default_value)
			} else {
				return default_value
			}
		} else {
			return args[index-1]
		}
	}
	return default_value
}

/**
indicate that this method does not support json output
*/
func NoJsonSupport() {
	if IsJsonOutput() {
		ErrorAndDie("this method does not support JSON output")
	}
}
