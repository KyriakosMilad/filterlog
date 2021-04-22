package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "", "The path to the log file that should be filtered")
	filter := flag.String("filter", "", "Log filters to search for, options available: any text that doesn't contain a comma ',', for multiple options separate them by coma: WARNING,ERROR")

	flag.Parse()

	if *path == "" {
		panic("path is required")
	}

	if *filter == "" {
		panic("filter is required")
	}

	var filters []string

	if strings.Contains(*filter, ",") {
		filters = strings.Split(*filter, ",")

		// remove empty strings
		for k, v := range filters {
			if v == "" {
				filters = append(filters[:k], filters[k+1:]...)
			}
		}
	} else {
		filters = []string{*filter}
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

		for _, f := range filters {
			if strings.Contains(s, f) {
				fmt.Print(s)
				break
			}
		}
	}
}
