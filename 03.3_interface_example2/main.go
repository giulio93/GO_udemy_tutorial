package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	cmd := os.Args
	f, err := os.Open(cmd[1])
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		fmt.Println(err)
	}

}
