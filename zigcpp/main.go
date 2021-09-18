package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jrossi/zigtool/comm"
)

func main() {
	_, err := exec.LookPath("zig")
	if err != nil {
		fmt.Println("Zig is not installed or not added to the path environment variable")
		return
	}
	os.Getwd()
	comm.Build("c++")
}
