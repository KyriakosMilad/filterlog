package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "example.log", "The path to the log file that should be filtered")
	level := flag.String("level", "INFO", "Log level to search for, options available are INFO, WARNING, ERROR, TRACE")

	flag.Parse()

	f, err := os.Open(*path)
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

		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
