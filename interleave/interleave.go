package interleave

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type wordsList []string

// InterleaveFiles interleave the files
func InterleaveFiles(file1, file2, delimiter string) (result []string, err error) {
	s1, err := readLines(file1, delimiter)
	if err != nil {
		return result, err
	}
	s2, err := readLines(file2, delimiter)
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

func readLines(path, delimiter string) (wordsList, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines wordsList
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
		re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
		final := re_leadclose_whtsp.ReplaceAllString(str, "")
		final = re_inside_whtsp.ReplaceAllString(final, " ")
		words := strings.Split(final, delimiter)
		lines = append(lines, words...)
		// remove empty string
		lines.delete("")
	}
	fmt.Println("Lines:", lines)
	return lines, nil
}

func (j *wordsList) delete(selector string) {
	var r wordsList
	for _, str := range *j {
		if str != selector {
			r = append(r, str)
		}
	}
	*j = r
}
