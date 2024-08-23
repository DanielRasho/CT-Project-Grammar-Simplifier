package io

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFileReader(t *testing.T) {
	// Create a temporary file with some content
	content := "line1\nline2\nline3\n"
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Read the file using FileReader
	fr, err := ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer fr.Close()

	var line string
	expectedLines := strings.Split(content, "\n")
	for i, expected := range expectedLines {
		if i == len(expectedLines)-1 {
			break // Last line split by strings.Split is an empty string
		}
		if !fr.NextLine(&line) {
			t.Fatalf("Expected line %d, got no more lines", i+1)
		}
		if line != expected {
			t.Errorf("Expected line %d to be %q, got %q", i+1, expected, line)
		}
	}
	if fr.NextLine(&line) {
		t.Errorf("Expected no more lines, but got %q", line)
	}
}
