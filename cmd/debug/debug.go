package debug

import (
	"fmt"
	"runtime"
)

func PrintMem() {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("%f MB\n", float64(rtm.Alloc)/1024./1024.)
}
