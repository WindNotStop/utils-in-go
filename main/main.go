package main

import (
	"github.com/WindNotStop/utils-in-go/concurrency/channel"
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("./trace.out")
	trace.Start(f)
	defer trace.Stop()
	channel.TeeChannel()
}
