package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}


func ParseLogLine(line string) (*LogEntry, error) {
	parts := strings.SplitN(line, " ", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid log line format: %s", line)
	}
	return &LogEntry{
		Timestamp: parts[0],
		Level:     parts[1],
		Message:   parts[2],
	}, nil
}

func main() {
	file, err := os.Open("example.log") 
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Filtered log entries with level ERROR:")
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := ParseLogLine(line)
		if err != nil {
			fmt.Printf("Skipping line: %v\n", err)
			continue
		}
		if entry.Level == "ERROR" {
			fmt.Printf("[%s] %s\n", entry.Timestamp, entry.Message)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
