package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fields    string
	delimiter string
	onlyDelim bool
)

func init() {
	flag.StringVar(&fields, "f", "", "выбрать только эти поля (столбцы)")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&onlyDelim, "s", false, "только строки с разделителем")
}

func main() {
	flag.Parse()

	if fields == "" {
		log.Fatal("Usage: cututil -f <fields> [-d <delimiter>] [-s]")
	}

	fieldIndices, err := parseFields(fields)
	if err != nil {
		log.Fatalf("Error parsing fields: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if onlyDelim && !strings.Contains(line, delimiter) {
			continue
		}

		columns := strings.Split(line, delimiter)
		selectedColumns := selectColumns(columns, fieldIndices)
		fmt.Println(strings.Join(selectedColumns, delimiter))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("ошибка ввода: %v", err)
	}
}

func parseFields(fields string) ([]int, error) {
	var indices []int
	parts := strings.Split(fields, ",")
	for _, part := range parts {
		var start, end int
		n, _ := fmt.Sscanf(part, "%d-%d", &start, &end)
		if n == 1 {
			indices = append(indices, start-1)
		} else if n == 2 {
			for i := start; i <= end; i++ {
				indices = append(indices, i-1)
			}
		} else {
			return nil, fmt.Errorf("неправильная спецификация поля: %s", part)
		}
	}
	return indices, nil
}

func selectColumns(columns []string, indices []int) []string {
	var selected []string
	for _, index := range indices {
		if index >= 0 && index < len(columns) {
			selected = append(selected, columns[index])
		}
	}
	return selected
}
