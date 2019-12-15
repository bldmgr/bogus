package main

import (
	"bufio"
	"fmt"
	"github.com/docker/docker/pkg/namesgenerator"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cmdName := "git"
	cmdArgs := []string{"checkout", "-b", namesgenerator.GetRandomName(0)}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error", err)
		os.Exit(1)
	}

	cmdoutput := bufio.NewScanner(cmdReader)
	go func() {
		for cmdoutput.Scan() {
			fmt.Printf("output | %s\n", cmdoutput.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Start Error", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Wait Error", err)
		os.Exit(1)
	}

}

