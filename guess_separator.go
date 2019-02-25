package main

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"bufio"
	"io"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s file\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	separators := []string{"\t", "*", "|", "."}

	lineCount, lines := readUpToNLines(os.Args[0], 5)
	// fmt.Printf("line count: %d\n", lineCount)
	// fmt.Printf("lines : %d\n", len(lines))

	counts := createCounts(lines, separators, lineCount)
	fmt.Print(counts)
}

func createCounts(lines, separators []string, lineCount int) [][]int {
	counts := make([][]int, len(separators))

	for sepInex := range separators {
		counts[sepInex] = make([]int, lineCount)
		for lineIndex, line := range lines {
			counts[sepInex][lineIndex] =
				strings.Count(line, separators[sepInex])
		}
	}

	return counts
}

func readUpToNLines(filename string, maxLines int) (int, []string) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Fatal("failed to open the file: ", err)
	}
	defer file.Close()
	lines := make([]string, maxLines)
	reader := bufio.NewReader(file)
	i := 0
	for ; i < maxLines; i++ {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines[i] = line
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("failed to finish reading the file", err)
		}
	}

	return i, lines[:i]
}