package fileHandling

import (
	"bufio"
	"log"
	"os"
)

func ReadFile() (s string) {
	file, err := os.Open("./fulldata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
