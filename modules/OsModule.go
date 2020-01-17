package modules

import (
	"fmt"
	"runtime"
	"time"

	"github.com/arashrasoulzadeh/serto.git/functions"
)

func ParseOSModule(args []string) {
	functions.SetProcessor("OS")
	has_command := functions.DieIfEqual(2, args, "Please Enter Sub Command.")
	if has_command {
		command := args[2]
		switch command {
		case "monitor":
			NewMonitor(1)
			break
		case "free":
			FreeSpace(args)
			break
		}

	}
}

/**
Monitor Struct
*/
type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64

	NumGC        uint32
	NumGoroutine int
}

/**
start monitoring
*/
func NewMonitor(duration int) {
	functions.NoJsonSupport()
	var m Monitor
	var rtm runtime.MemStats
	var interval = time.Duration(duration) * time.Second
	for {
		<-time.After(interval)

		// Read full mem stats
		runtime.ReadMemStats(&rtm)

		// Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		m.Alloc = rtm.Alloc
		m.TotalAlloc = rtm.TotalAlloc
		m.Sys = rtm.Sys
		m.Mallocs = rtm.Mallocs
		m.Frees = rtm.Frees

		// Live objects = Mallocs - Frees
		m.LiveObjects = m.Mallocs - m.Frees

		// GC Stats
		m.PauseTotalNs = rtm.PauseTotalNs
		m.NumGC = rtm.NumGC

		// Just encode to json and print
		functions.CallClear()
		functions.PrettyPrint(&m)
		// fmt.Println(string(b))
	}
}

/**
show Free Space
*/
func FreeSpace(args []string) {
	path := functions.GetArgOrDefault(4, "/")
	disk := functions.DiskUsage(path)
	if functions.IsJsonOutput() {
		functions.PrettyPrint(disk)
	} else {
		functions.Verbose(fmt.Sprintf("Total : %.2f GB", float64(disk.All)/float64(functions.GB)))
		functions.Verbose(fmt.Sprintf("Used  : %.2f GB", float64(disk.Used)/float64(functions.GB)))
		functions.Verbose(fmt.Sprintf("Free  : %.2f GB", float64(disk.Free)/float64(functions.GB)))
	}

}
