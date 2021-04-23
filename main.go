package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	path := flag.String("path", "", "The path to the log file that should be filtered")
	filter := flag.String("filter", "", "Log filters to search for")
	separator := flag.String("separator", "", "Separate between multiple filters")
	outputFile := flag.Bool("output_file", false, "Which to output the result to file or not")

	flag.Parse()

	if *path == "" && *filter == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *path == "" {
		panic("path is required")
	}

	if *filter == "" {
		panic("filter is required")
	}

	var filters []string

	if *separator != "" && strings.Contains(*filter, *separator) {
		filters = strings.Split(*filter, *separator)

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

	// check if outputFile == true here better than checking it every line inside the loop
	if *outputFile {
		fileName := "filer_log_" + time.Now().Format("2006-01-02_15:04:05") + ".log"
		outputFileName, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(outputFileName)

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
					_, err := outputFileName.WriteString(s)

					if err != nil {
						panic(err)
					}
					break
				}
			}
		}
	} else {
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
}
