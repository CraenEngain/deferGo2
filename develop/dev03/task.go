package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column       int
	numeric      bool
	reverse      bool
	unique       bool
	month        bool
	ignoreBlanks bool
	check        bool
	humanNumeric bool
)

func init() {
	flag.IntVar(&column, "k", 0, "номер столбца для сортировки (начиная с 1)")
	flag.BoolVar(&numeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить дубликаты строк")
	flag.BoolVar(&month, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&ignoreBlanks, "b", false, "игнорировать конечные пробелы")
	flag.BoolVar(&check, "c", false, "проверить, отсортированы ли данные")
	flag.BoolVar(&humanNumeric, "h", false, "сортировать по понятному человеку числовому значению (например, 2K 1M)")
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("Использование: sortutil [options] inputfile outputfile")
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := readLines(inputFile)
	if err != nil {
		log.Fatalf("Ошибка чтения входного файла: %v", err)
	}

	if check {
		if isSorted(lines) {
			fmt.Println("Файл отсортирован.")
			return
		} else {
			fmt.Println("Файл не отсортирован.")
			return
		}
	}

	sortLines(lines)

	if unique {
		lines = uniqueLines(lines)
	}

	if err := writeLines(outputFile, lines); err != nil {
		log.Fatalf("Ошибка записи выходного файла: %v", err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if ignoreBlanks {
			line = strings.TrimRight(line, " \t")
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func writeLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func sortLines(lines []string) {
	sort.Slice(lines, func(i, j int) bool {
		keyI := getKey(lines[i])
		keyJ := getKey(lines[j])

		var less bool
		if numeric {
			numI, errI := strconv.ParseFloat(keyI, 64)
			numJ, errJ := strconv.ParseFloat(keyJ, 64)
			if errI == nil && errJ == nil {
				less = numI < numJ
			} else {
				less = keyI < keyJ
			}
		} else if month {
			less = compareMonths(keyI, keyJ)
		} else if humanNumeric {
			less = compareHumanReadable(keyI, keyJ)
		} else {
			less = keyI < keyJ
		}

		if reverse {
			return !less
		}
		return less
	})
}

func getKey(line string) string {
	if column > 0 {
		fields := strings.Fields(line)
		if column-1 < len(fields) {
			return fields[column-1]
		}
	}
	return line
}

func uniqueLines(lines []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniqueLines []string

	for _, line := range lines {
		if _, exists := uniqueMap[line]; !exists {
			uniqueMap[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines
}

func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if lines[i-1] > lines[i] {
			return false
		}
	}
	return true
}

func compareMonths(month1, month2 string) bool {
	months := map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
		"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
		"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}
	return months[month1] < months[month2]
}

func compareHumanReadable(val1, val2 string) bool {
	parseHumanReadable := func(val string) float64 {
		unit := map[byte]float64{
			'K': 1e3, 'M': 1e6, 'G': 1e9, 'T': 1e12, 'P': 1e15, 'E': 1e18,
		}
		if len(val) == 0 {
			return 0
		}
		if factor, exists := unit[val[len(val)-1]]; exists {
			num, err := strconv.ParseFloat(val[:len(val)-1], 64)
			if err != nil {
				return 0
			}
			return num * factor
		}
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return num
	}

	return parseHumanReadable(val1) < parseHumanReadable(val2)
}
