package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go file1.txt file2.txt ...")
		os.Exit(1)
	}

	// Create a map to store three-word sequences and their counts
	sequencesCount := make(map[string]int)

	// Process each file or stdin input
	for _, filePath := range os.Args[1:] {
		if filePath == "-" {
			// Read from stdin
			processInput(os.Stdin, sequencesCount)
		} else {
			// Process file
			err := processFile(filePath, sequencesCount)
			if err != nil {
				fmt.Printf("Error processing file %s: %v\n", filePath, err)
			}
		}
	}

	// Sort and print the 100 most common three-word sequences
	printTopSequences(sequencesCount, 100)
}

func processFile(filePath string, sequencesCount map[string]int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	processInput(file, sequencesCount)
	return nil
}

func processInput(input io.Reader, sequencesCount map[string]int) {
	scanner := bufio.NewScanner(input)
	prevWords := []string{}

	// Define a regular expression to match words
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)

	// Read word by word
	for scanner.Scan() {
		line := scanner.Text()
		// Remove punctuation, convert to lowercase
		line = cleanLine(line)
		// Split the line into words using regular expression
		matches := wordRegex.FindAllString(line, -1)

		for _, word := range matches {
			// Append lowercase words to the slice
			word = strings.ToLower(word)
			prevWords = append(prevWords, word)
			if len(prevWords) >= 3 {
				// Build and update three-word sequences
				sequence := strings.Join(prevWords[len(prevWords)-3:], " ")
				sequencesCount[sequence]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
}

func cleanLine(line string) string {
	// Remove punctuation, convert to lowercase
	line = strings.ToLower(line)
	line = strings.Map(func(r rune) rune {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz ", r) {
			return r
		}
		return ' '
	}, line)
	return line
}

func printTopSequences(sequencesCount map[string]int, n int) {
	fmt.Println("Top", n, "three-word sequences:")
	// Sort sequences by count in descending order
	counts := make([]struct {
		seq   string
		count int
	}, 0, len(sequencesCount))
	for seq, count := range sequencesCount {
		counts = append(counts, struct {
			seq   string
			count int
		}{seq, count})
	}
	sortByCountDesc(counts)
	for i, item := range counts {
		if i >= n {
			break
		}
		fmt.Printf("%d - %s\n", item.count, item.seq)
	}
}

func sortByCountDesc(counts []struct {
	seq   string
	count int
},
) {
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].seq < counts[j].seq
		}
		return counts[i].count > counts[j].count
	})
}
