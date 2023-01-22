package files

import (
	"bufio"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// FileToLines reads a file and returns a lines slice
func FileToLines(path string) (lines []string) {
	const maxLineLength = 1000000000

	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, maxLineLength)
	scanner.Buffer(buf, maxLineLength)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	check(err)

	return lines
}

// FileToString reads a file and returns contents as a string
func FileToString(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	check(err)
	return string(bytes)
}

// FileToBytes reads a file and returns contents as a byte slice
func FileToBytes(fileName string) []byte {
	bytes, err := ioutil.ReadFile(fileName)
	check(err)
	return bytes
}

