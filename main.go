package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type filterlog struct {
	path      string
	separator string
	filters   []string
	results   []string
}

func (fl *filterlog) new(path string, filters string, separator string) {
	fl.path = path
	fl.separator = separator
	if separator != "" && strings.Contains(filters, separator) {
		fl.filters = strings.Split(filters, separator)
		fl.removeEmptyFilters()
	} else {
		fl.filters = []string{filters}
	}
}

func (fl *filterlog) removeEmptyFilters() {
	for i, v := range fl.filters {
		if v == "" {
			fl.filters = append(fl.filters[:i], fl.filters[i+1:]...)
		}
	}
}

func (fl *filterlog) addResult(result string) {
	fl.results = append(fl.results, result)
}

func (fl *filterlog) search() {
	f, err := os.Open(fl.path)
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

		for _, f := range fl.filters {
			if strings.Contains(s, f) {
				fl.addResult(s)
				fmt.Print(s)
				break
			}
		}
	}
}

func (fl *filterlog) exportResults() {
	fileName := "filerlog_" + time.Now().Format("2006-01-02_15:04:05") + ".log"
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	for _, v := range fl.results {
		_, err := file.WriteString(v)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	path := flag.String("path", "", "The path to the log file that should be filtered")
	filters := flag.String("filters", "", "Log filters to search for")
	separator := flag.String("separator", "", "Separate between multiple filters")
	exportResults := flag.Bool("export_results", false, "Which to print the results to a file or not")

	flag.Parse()
	if *path == "" && *filters == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *path == "" {
		panic("path is required")
	}
	if *filters == "" {
		panic("one filter is required at least")
	}

	fl := &filterlog{}
	fl.new(*path, *filters, *separator)
	fl.search()
	if *exportResults {
		fl.exportResults()
	}
}
