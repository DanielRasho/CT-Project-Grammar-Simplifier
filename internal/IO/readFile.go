package io

import (
	"bufio"
	"os"
)

type FileReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

// readFile opens the file and returns a FileReader instance
func ReadFile(path string) (*FileReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

// NextLine reads the next line from the file and stores it in the provided string pointer
func (fr *FileReader) NextLine(line *string) bool {
	if fr.scanner.Scan() {
		*line = fr.scanner.Text()
		return true
	}
	return false
}

func (fr *FileReader) Close() error {
	return fr.file.Close()
}
