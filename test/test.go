package main

import (
	"os/exec"
	"runtime"
	"time"
)

func main() {
	timer := time.NewTicker(time.Minute)
	runtime.GOMAXPROCS(1)

	for range timer.C {
		cmd := exec.Command("/bin/bash", "/Users/star/dev_env/bin/clean")
		_ = cmd.Start()
	}

}
