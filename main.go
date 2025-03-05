package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

// getOffsets searches for all occurrences of target in data,
// returning their starting indices (only if index >= 8).
func getOffsets(data []byte, target string) []int {
	var offsets []int
	targetBytes := []byte(target)
	start := 0
	for {
		index := bytes.Index(data[start:], targetBytes)
		if index == -1 {
			break
		}
		actualIndex := start + index
		if actualIndex >= 8 {
			offsets = append(offsets, actualIndex)
		}
		start = actualIndex + len(targetBytes)
	}
	return offsets
}

// formatBytes returns a string with each byte formatted as two-digit hex.
func formatBytesToHex(b []byte) string {
	parts := make([]string, len(b))
	for i, v := range b {
		parts[i] = fmt.Sprintf("%02X", v)
	}
	return strings.Join(parts, " ")
}

// byteToChar returns a printable character for a given byte.
// Non-printable bytes (except for newline 0x0A and carriage return 0x0D)
// are replaced with a dot. Newline and carriage return are replaced with "\n" and "\r".
func byteToChar(b byte) string {
	if (b < 32 || b > 126) && b != 10 && b != 13 {
		return "."
	} else if b == 10 {
		return "\\n"
	} else if b == 13 {
		return "\\r"
	} else {
		return string(b)
	}
}

// formatBytesChar converts a byte slice to a string using byteToChar for each byte.
func formatBytesToChar(b []byte) string {
	var result string
	for _, v := range b {
		result += byteToChar(v)
	}
	return result
}

// readBytesBefore returns the n bytes that precede the given offset.
func readBytesBefore(data []byte, offset, n int) []byte {
	start := offset - n
	if start < 0 {
		start = 0
	}
	return data[start:offset]
}

// readBytesAfter returns the n bytes that follow the given offset.
func readBytesAfter(data []byte, offset, n int) []byte {
	end := offset + n
	if end > len(data) {
		end = len(data)
	}
	return data[offset:end]
}

func main() {
	inputFile := flag.String("infile", "", "path to the input file")
	targetString := flag.String("string", "", "target string to search for in the input file")
	before := flag.Int("before", 8, "number of bytes to read before each offset")
	beforeFormat := flag.String("beforeFormat", "hex", "format of bytes before each offset: hex or char")
	after := flag.Int("after", 16, "number of bytes to read after each offset")
	afterFormat := flag.String("afterFormat", "hex", "format of bytes after each offset: hex or char")
	flag.Parse()

	if *inputFile == "" || *targetString == "" {
		fmt.Println("usage: byteinspect -input <file> -target <string> [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("error reading file %s: %v\n", *inputFile, err)
		os.Exit(1)
	}

	offsets := getOffsets(data, *targetString)

	var inspectionLines []string
	for _, offset := range offsets {
		var beforeStr string
		if *before > 0 {
			precedingData := readBytesBefore(data, offset, *before)
			if *beforeFormat == "char" {
				beforeStr = formatBytesToChar(precedingData)
			} else {
				beforeStr = formatBytesToHex(precedingData)
			}
			beforeStr += "\t"
		}

		var afterStr string
		if *after > 0 {
			followingData := readBytesAfter(data, offset, *after)
			if *afterFormat == "char" {
				afterStr = formatBytesToChar(followingData)
			} else {
				afterStr = formatBytesToHex(followingData)
			}
		}

		line := fmt.Sprintf("%s%08X\t%s", beforeStr, offset, afterStr)
		inspectionLines = append(inspectionLines, line)
	}

	// sort.Strings(inspectionLines)

	for _, line := range inspectionLines {
		fmt.Println(line)
	}
}
