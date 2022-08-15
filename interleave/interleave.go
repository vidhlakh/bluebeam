package interleave

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type wordsList []string

// InterleaveFiles interleave the files
func InterleaveFiles(file1, file2 string) (result []string, err error) {
	s1, err := readLines(file1)
	if err != nil {
		return result, err
	}
	s2, err := readLines(file2)
	if err != nil {
		return result, err
	}
	i := 0
	// Time complexity max(len(s1,s2))
	for i < len(s1) && i < len(s2) {
		result = append(result, s1[i], s2[i])
		i++
	}
	for i < len(s1) {
		result = append(result, s1[i])
		i++
	}
	for i < len(s2) {
		result = append(result, s2[i])
		i++
	}
	return result, nil
}

func readLines(path string) (wordsList, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines wordsList
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsSymbol(c)
		}
		words := strings.FieldsFunc(str, f)
		lines = append(lines, words...)
		// remove strings with single length symbols
		lines.removeSingleSymbols()
	}
	fmt.Println("Lines:", lines, len(lines))
	return lines, nil
}

func (j *wordsList) removeSingleSymbols() {
	var r wordsList
	for _, str := range *j {

		s := []rune(str)
		if len(s) == 1 && unicode.IsSymbol(s[0]) {
			continue
		} else {
			r = append(r, str)
		}
	}
	*j = r
}
