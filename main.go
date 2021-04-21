package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("example.log")
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				panic(err)
			}
		}

		if strings.Contains(s, "INFO") {
			fmt.Println(s)
		}
	}
}
