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

func TestPrintDirWithoutHiddenFiles(t *testing.T) {
	files, testDir := setupWithoutHidden()
	defer os.RemoveAll(testDir)
	// Create a buffer to write to and
	// set output to the buffer
	var buf bytes.Buffer
	Config.output = &buf
	// Clean up after the test run
	// Sort the test files and join them for the
	// expected result
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })
	expected := strings.Join(files, "\n")
	err := printDir(testDir, false)
	if err != nil {
		t.Fatalf("printDir() failed: %s", err)
	}
	fmt.Println(buf.String())
	// Need to call #TrimRight for the result
	// because fmt.Fprintln appends a \n after each
	// line, including the last line, but strings.Join does
	// not append the <sep> after the last item.
	if strings.Compare(strings.TrimRight(buf.String(), "\n"), expected) != 0 {
		t.Fatalf("expected %s, got %s", expected, buf.String())
	}
}

func TestPrintDirWithHiddenFiles(t *testing.T) {
	files, testDir := setupWithHidden()
	defer os.RemoveAll(testDir)
	// Create a buffer to write to and
	// set output to the buffer
	var buf bytes.Buffer
	Config.output = &buf
	// Clean up after the test run
	// Sort the test files and join them for the
	// expected result
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })
	expected := strings.Join(files, "\n")
	err := printDir(testDir, true)
	if err != nil {
		t.Fatalf("printDir() failed: %s", err)
	}
	fmt.Println(buf.String())
	// Need to call #TrimRight for the result
	// because fmt.Fprintln appends a \n after each
	// line, including the last line, but strings.Join does
	// not append the <sep> after the last item.
	if strings.Compare(strings.TrimRight(buf.String(), "\n"), expected) != 0 {
		t.Fatalf("expected %s, got %s", expected, buf.String())
	}
}
