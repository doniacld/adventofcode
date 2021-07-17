package cypher

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// ReadAndExtract reads the file line by line and applies the given function to extract what we want
func ReadAndExtract(path string, idx int) (Cypher, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    // initialization
	c := newCypher()
	nbs := make([]int, 0)
	i := 0

    // scan lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elt, err := convertToInt(scanner.Text())
		if err != nil {
			return c, err
		}

        // if we are in the preamble range, store the value in the preamble map
		if i < idx {
			err = c.fillPreamble(elt)
		}
		if err != nil {
			return c, err
		}

        // store all the numbers
		nbs = append(nbs, elt)
		i++
	}

	c.Numbers = nbs
	return c, scanner.Err()
}

// fills the preamble map
func (c Cypher) fillPreamble(elt int) error {
	c.Preamble.insert(elt)
	return nil
}

// convertToInt converts the line into an int
func convertToInt(line string) (int, error) {
	elt, err := strconv.Atoi(line)
	if err != nil {
		return elt, err
	}
	return elt, nil
}
