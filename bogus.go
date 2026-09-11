package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"

	"github.com/bldmgr/bogus/pkg/namesgenerator"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

// randomString generates a cryptographically secure random alphanumeric string of the specified length.
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

// generateBranchName generates a Docker-style random name with a 6-character random suffix.
func generateBranchName() string {
	return fmt.Sprintf("%s-%s", namesgenerator.GetRandomName(0), randomString(6))
}

func main() {
	cmdName := "git"
	cmdArgs := []string{"checkout", "-b", generateBranchName()}

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
