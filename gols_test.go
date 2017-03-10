package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func setupWithoutHidden() (files []os.FileInfo, testDir string) {
	// Create a test directory, 5 regular test files
	tmp, _ := ioutil.TempDir("", "test-dir")
	testDir = tmp
	var filenames []string
	for i := 0; i < 5; i++ {
		file, _ := ioutil.TempFile(testDir, "test-file")
		filenames = append(filenames, file.Name())
	}
	for _, file := range filenames {
		info, _ := os.Stat(file)
		files = append(files, info)
	}
	return
}

func setupWithHidden() (files []os.FileInfo, testDir string) {
	// Create a test directory, 5 regular test files,
	// and 1 hidden file
	tmp, _ := ioutil.TempDir("", "test-dir")
	testDir = tmp
	var filenames []string
	for i := 0; i < 5; i++ {
		file, _ := ioutil.TempFile(testDir, "test-file")
		filenames = append(filenames, file.Name())
	}
	hidden, _ := ioutil.TempFile(testDir, ".file")
	filenames = append(filenames, hidden.Name())
	for _, file := range filenames {
		info, _ := os.Stat(file)
		files = append(files, info)
	}
	return
}

func TestreturnFilesWithoutHiddenFiles(t *testing.T) {
	files, testDir := setupWithoutHidden()
	// Clean up after the test run
	defer os.RemoveAll(testDir)
	// Sort the test files and join them for the
	// expected result
	actual, err := returnFiles(testDir, false)
	if err != nil {
		t.Fatalf("returnFiles() failed: %s", err)
	}
	if fmt.Sprintf("%v", files) != fmt.Sprintf("%v", actual) {
		t.Fatalf("expected %s, got %s", files, actual)
	}
}

func TestreturnFilesWithHiddenFiles(t *testing.T) {
	files, testDir := setupWithHidden()
	// Clean up after the test run
	defer os.RemoveAll(testDir)
	// Sort the test files
	actual, err := returnFiles(testDir, true)
	if err != nil {
		t.Fatalf("returnFiles() failed: %s", err)
	}
	if fmt.Sprintf("%v", files) != fmt.Sprintf("%v", actual) {
		t.Fatalf("expected %s, got %s", files, actual)
	}
}

func TestprintResultsWithoutLongOutput(t *testing.T) {
	files, testDir := setupWithoutHidden()
	// Clean up after the test run
	defer os.RemoveAll(testDir)
	// Create a buffer to write to and
	// set output to the buffer
	var buf bytes.Buffer
	Config.output = &buf
	// Sort the test files
	var f []string
	for _, file := range files {
		f = append(f, file.Name())
	}
	expected := strings.Join(f, " ")
	actual, _ := returnFiles(testDir, false)
	err := printResults(actual, false)
	if err != nil {
		t.Fatalf("printResults() failed: %s", err)
	}
	if expected != buf.String() {
		t.Fatalf("expected %s, got %s", expected, buf.String())
	}
}

func TestprintResultsWithLongOutput(t *testing.T) {
	files, testDir := setupWithoutHidden()
	// Clean up after the test run
	defer os.RemoveAll(testDir)
	// Create a buffer to write to and
	// set output to the buffer
	var buf bytes.Buffer
	Config.output = &buf
	// Sort the test files
	var f []string
	for _, file := range files {
		f = append(f, file.Name())
	}
	expected := strings.Join(f, " ")
	actual, _ := returnFiles(testDir, false)
	err := printResults(actual, false)
	if err != nil {
		t.Fatalf("printResults() failed: %s", err)
	}
	if expected != buf.String() {
		t.Fatalf("expected %s, got %s", expected, buf.String())
	}
}
