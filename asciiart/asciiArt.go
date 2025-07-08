package asciiart

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func GetAsciiLine(filename string, num int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line, nil
}

func AsciiArt(input, filename string) (string, error) {
	banner := "banners/" + filename + ".txt"
	result := "" // Start with an empty result string

	args := strings.Split(input, "\n")

	// Handle the case where the first line might be empty (new line at the beginning)
	if len(args) > 0 && args[0] == "" {
		result += "\n" // Add a single newline at the start if the input starts with \n
	}

	// Process each word (or line) from the split input
	for _, word := range args {
		// Skip over any empty lines (we've already handled the first empty line above)
		if word == "" {
			result += "\n"
			continue
		}

		// Process the word character by character
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				asciiLine, err := GetAsciiLine(banner, 1+int(letter-' ')*9+i)
				if err != nil {
					return "", err
				}
				result += asciiLine
			}
			result += "\n" // Add a newline after processing each line of characters
		}
	}

	return result, nil
}

// SplitNewLine splits the input strings on newlines (\n or \\n)
func SplitNewLine(input []string) []string {
	var lines []string
	re := regexp.MustCompile(`\n|\\n`)

	for _, arg := range input {
		// Split the string by regex matches of \n or \\n
		split := re.Split(arg, -1)
		lines = append(lines, split...)
	}

	return lines
}
