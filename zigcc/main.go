package main

import (
	"fmt"
	"os/exec"

	"github.com/jrossi/zigtool/comm"
)

func main() {
	_, err := exec.LookPath("zig")
	if err != nil {
		fmt.Println("Zig is not installed or not added to the path environment variable")
		return
	}
	comm.Build("cc")
}
