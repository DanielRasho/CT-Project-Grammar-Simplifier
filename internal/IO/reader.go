package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReaderTXT lee un archivo y retorna una lista de líneas filtradas
func ReaderTXT(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue // Ignora líneas vacías y líneas que comienzan con '#'
		}
		lines = append(lines, trimmedLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer el archivo: %w", err)
	}

	return lines, nil
}
