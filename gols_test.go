package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func setupWithoutHidden() (files []string, testDir string) {
	// Create a test directory, 5 regular test files
	tmp, _ := ioutil.TempDir("", "test-dir")
	testDir = tmp
	for i := 0; i < 5; i++ {
		file, _ := ioutil.TempFile(testDir, "test-file")
		files = append(files, filepath.Base(file.Name()))
	}
	return
}

func setupWithHidden() (files []string, testDir string) {
	// Create a test directory, 5 regular test files,
	// and 1 hidden file
	tmp, _ := ioutil.TempDir("", "test-dir")
	testDir = tmp
	for i := 0; i < 5; i++ {
		file, _ := ioutil.TempFile(testDir, "test-file")
		files = append(files, filepath.Base(file.Name()))
	}
	hidden, _ := ioutil.TempFile(testDir, ".file")
	files = append(files, filepath.Base(hidden.Name()))
	return
}

func TestReturnFileNamesWithoutHiddenFiles(t *testing.T) {
	files, testDir := setupWithoutHidden()
	// Clean up after the test run
	defer os.RemoveAll(testDir)
	// Sort the test files and join them for the
	// expected result
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })
	actual, err := returnFileNames(testDir, false)
	if err != nil {
		t.Fatalf("returnFileNames() failed: %s", err)
	}
	if fmt.Sprintf("%v", files) != fmt.Sprintf("%v", actual) {
		t.Fatalf("expected %s, got %s", files, actual)
	}
}

func TestreturnFileNamesWithHiddenFiles(t *testing.T) {
	files, testDir := setupWithHidden()
	// Clean up after the test run
	defer os.RemoveAll(testDir)
	// Sort the test files
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })
	actual, err := returnFileNames(testDir, true)
	if err != nil {
		t.Fatalf("returnFileNames() failed: %s", err)
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
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })
	expected := strings.Join(files, " ")
	err := printResults(files, false)
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
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })
	expected := strings.Join(files, "\n")
	err := printResults(files, true)
	if err != nil {
		t.Fatalf("printResults() failed: %s", err)
	}
	if expected != buf.String() {
		t.Fatalf("expected %s, got %s", expected, buf.String())
	}
}
