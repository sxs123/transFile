package controller

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	// "strings"
)

func main() {
	cmd1 := exec.Command("/bin/sh", "-c", "kill -9 `ps -ef |grep myserver|awk '{print $2}'`")
	cmd2 := exec.Command("go", "run", "myserver.go")
	var out bytes.Buffer
	cmd1.Stdout = &out
	err := cmd1.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status of server: %q\n", out.String())
	cmd2.Stdout = &out
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status of server", out.String())

}
