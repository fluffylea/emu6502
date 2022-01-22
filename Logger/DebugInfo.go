package Logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DebugListingFile *string
var DebugMappingFile *string

var debugListingFileContent *[]string
var debugMapping *map[string]string

// LogDebugInstruction extracts the debug instruction from the debug listing file
// and prints it to the console.
func LogDebugInstruction(pc uint16) {
	if DebugListingFile == nil {
		// Without a listing file we can't do anything
		Warnf("Cannot log debug instruction without a listing file")
		return
	}

	if debugListingFileContent == nil {
		// On first call, load the listing file
		debugListingFileContent = readFile(DebugListingFile)
	}

	if debugMapping == nil {
		if DebugMappingFile == nil {
			// Without a mapping file we can still continue, just without showing label names
			Warnf("Cannot log debug instruction without a mapping file")
			tmpMapping := make(map[string]string)
			debugMapping = &tmpMapping
		} else {
			debugMapping = readMappingFile(DebugMappingFile)
		}
	}

	pcPrefix := fmt.Sprintf(" %04X", pc)
	for _, line := range *debugListingFileContent {
		if strings.HasPrefix(line, pcPrefix) {
			// Add the label name to the address if available
			for key, value := range *debugMapping {
				line = strings.Replace(line, key, fmt.Sprintf("%s(%s)", key, value), -1)
			}
			Debugf("Running: %s", line)
			return
		}
	}
}

// readFile reads a file and returns its content as a slice of strings.
func readFile(fileName *string) *[]string {
	file, _ := os.Open(*fileName)
	scanner := bufio.NewScanner(file)
	tempLines := make([]string, 0)
	for scanner.Scan() {
		tempLines = append(tempLines, scanner.Text())
	}
	return &tempLines
}

// readMappingFile reads the mapping file and returns a map of the labels and their
// corresponding addresses.
func readMappingFile(fileName *string) *map[string]string {
	// File format:
	// $<ADDRESS> | <LABEL>      | <FILE>
	// Returns:
	// <ADDRESS> -> <LABEL>
	file, err := os.Open(*fileName)
	if err != nil {
		Warnf("Could not open mapping file: %s", err)
		return nil
	}
	scanner := bufio.NewScanner(file)
	tempMapping := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			parts := strings.Split(line, "|")
			tempMapping[strings.TrimLeft(strings.TrimSpace(parts[0]), "$")] = strings.TrimSpace(parts[1])
		}
	}
	return &tempMapping
}
