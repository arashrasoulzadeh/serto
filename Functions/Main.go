package Functions

import (
	"bufio"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

var processor string = "serto"

func Verbose(msg string) {
	color.Green("%s => %s", processor, msg)
}
func GrepOutput(msg string, filter string) {
	scanner := bufio.NewScanner(strings.NewReader(msg))
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, filter) {
			Verbose(scanner.Text())
		}
	}
}

func Error(err string) {
	color.Red("%s => %s", processor, err)
}
func ErrorAndDie(err string) {
	Error(err)
	os.Exit(0)
}
func ClientOS() string {
	return "mac"
}
func DieIfEqual(size int, arr []string, err string) bool {
	if len(arr) == size {
		ErrorAndDie(err)
		return false
	}
	return false
}
func SetProcessor(proc string) {
	processor = proc
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
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

func OpenFileForEdit(path string) {
	ExecuteCommand("nano", path, "", "", true)
}


func help() {

}