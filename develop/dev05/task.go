package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
)

func init() {
	flag.IntVar(&after, "A", 0, "print +N lines after match")
	flag.IntVar(&before, "B", 0, "print +N lines before match")
	flag.IntVar(&context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&count, "c", false, "print count of matching lines")
	flag.BoolVar(&ignoreCase, "i", false, "ignore case")
	flag.BoolVar(&invert, "v", false, "invert match")
	flag.BoolVar(&fixed, "F", false, "fixed string match")
	flag.BoolVar(&lineNum, "n", false, "print line number")
}

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("Usage: greputil [options] pattern file")
	}

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("Error reading lines: %v", err)
	}

	matches := grep(lines, pattern)

	if count {
		fmt.Println(len(matches))
	} else {
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}

func readLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func grep(lines []string, pattern string) []string {
	var result []string
	var regex *regexp.Regexp
	var err error

	if ignoreCase {
		pattern = "(?i)" + pattern
	}

	if fixed {
		regex, err = regexp.Compile(regexp.QuoteMeta(pattern))
	} else {
		regex, err = regexp.Compile(pattern)
	}

	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
	}

	contextLines := max(after, before)
	matchCount := 0
	for i, line := range lines {
		match := regex.MatchString(line)
		if invert {
			match = !match
		}

		if match {
			if contextLines > 0 {
				start := max(0, i-contextLines)
				end := min(len(lines), i+contextLines+1)
				for j := start; j < end; j++ {
					if lineNum {
						result = append(result, fmt.Sprintf("%d:%s", j+1, lines[j]))
					} else {
						result = append(result, lines[j])
					}
				}
			} else {
				if lineNum {
					result = append(result, fmt.Sprintf("%d:%s", i+1, line))
				} else {
					result = append(result, line)
				}
			}
			matchCount++
			if contextLines > 0 {
				i += contextLines
			}
		}
	}

	if count {
		result = []string{fmt.Sprintf("%d", matchCount)}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
