package filereader

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadAndExtract reads the file line by line and applies the given function to extract what we want
func ReadAndExtract(path string, f func(s string) error) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err := f(scanner.Text())
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

// ExtractStrings returns a slice of all the lines in a file
func ExtractStrings(path string) ([]string, error) {
	var lines []string
	err := ReadAndExtract(path, func(s string) error {
		lines = append(lines, s)
		return nil
	})
	return lines, err
}

// ExtractInt reads the file line by line and returns a list of int
func ExtractInt(path string) ([]int, error) {
	var lines []int
	err := ReadAndExtract(path, func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		lines = append(lines, i)
		return nil
	})
	return lines, err
}

// Extract all the values into an array
func ExtractStringsArray(path string) ([][]string, error) {
	chars := make([][]string, 0)
	err := ReadAndExtract(path, func(s string) error {
		chars = append(chars, strings.Split(s, ""))
		return nil
	})
	return chars, err
}
