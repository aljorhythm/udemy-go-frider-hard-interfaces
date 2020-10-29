package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR")
	} else {
		io.Copy(os.Stdout, bytes.NewReader(data))
	}
}
