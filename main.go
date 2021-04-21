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
	level := flag.String("level", "INFO", "Log level to search for, options available are INFO, WARNING, ERROR, TRACE, for multiple options separate them by coma: WARNING,ERROR")

	flag.Parse()

	var levels []string

	if strings.Contains(*level, ",") {
		levels = strings.Split(*level, ",")

		// remove empty strings
		for k, v := range levels {
			if v == "" {
				levels = append(levels[:k], levels[k+1:]...)
			}
		}
	} else {
		levels = []string{*level}
	}

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

		for _, l := range levels {
			if strings.Contains(s, l) {
				fmt.Println(s)
				break
			}
		}
	}
}
