package handleFile

import (
	"bufio"
	"fmt"
	"os"
)

//pathExample: "./sample.txt"
func ReadFileLineByLine(path string, output chan string) {
	defer os.Remove(path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot read file from path: %v", path)
	}
	defer file.Close()
	fmt.Printf("reading %v...", path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output <- scanner.Text()

	}
	output <- "READDONE"
}
