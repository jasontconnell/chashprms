package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func sha256sum(d []byte) []byte {
	h := sha256.New()
	h.Write(d)
	return h.Sum(nil)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovering from an error")
		}
	}()

	contents := ""
	if len(os.Args) == 2 {
		contents = strings.Join(os.Args[1:], " ")
	}

	sum := sha256sum([]byte(contents))

	fmt.Println(fmt.Sprintf("%x", sum))
}
