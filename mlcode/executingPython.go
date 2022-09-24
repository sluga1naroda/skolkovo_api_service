package mlcode

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func Execute(param string) string {
	cmd := exec.Command("python", "script.py", param)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)

	cmd.Wait()

	return "hello"
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
