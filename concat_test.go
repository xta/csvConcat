package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	output := "tmp_123456.csv"
	removeIfExists(output)

	cfg.inFilesRaw = "test/file1.csv test/file2.csv test/file3.csv"
	cfg.outFile = output

	main()

	lineCount, err := fileLineCount(output)
	if err != nil {
		t.Fatal(err)
	}

	if lineCount != 18 {
		t.Fatal("Error in main, got unexpected line count:", lineCount)
	}

	removeIfExists(output)
}

func TestFirstInFile(t *testing.T) {
	files := []string{"abc.csv", "def.csv"}

	if firstInFile(files) != "abc.csv" {
		t.Error("Error: did not get correct first file")
	}

	// TODO: refactor firstInFile() to return error
	invalidFiles := []string{}
	if firstInFile(invalidFiles) != "" {
		t.Error("Error: did not get expected")
	}
}

func TestLatterInFiles(t *testing.T) {
	files := []string{"abc.csv", "def.csv", "ghi.csv"}
	latter := latterInFiles(files)

	if len(latter) != 2 {
		t.Error("Error: did not get correct latter files")
	}

	if contains(latter, "abc.csv") {
		t.Error("Error: got first element in latter files")
	}

	moreFiles := []string{"abc.csv", "def.csv", "ghi.csv", "jkl.csv", "mno.csv"}
	moreLatter := latterInFiles(moreFiles)
	if len(moreLatter) != 4 {
		t.Error("Error: did not get correct latter files")
	}

	if contains(moreLatter, "abc.csv") {
		t.Error("Error: got first element in latter files")
	}

	oneFile := []string{"abc.csv"}
	noLatter := latterInFiles(oneFile)

	if len(noLatter) != 0 {
		t.Error("Error: did not get correct latter files")
	}
}

func TestCopyToOutputCopyHeader(t *testing.T) {
	input := "test/file1.csv"
	output := "test/tmp_123.csv"

	removeIfExists(output)

	err := copyToOutput(input, output, true)
	if err != nil {
		t.Fatal(err)
	}

	lineCount, err := fileLineCount(output)
	if err != nil {
		t.Fatal(err)
	}

	if lineCount != 5 {
		t.Fatal("Error while copyToOutput, got unexpected line count:", lineCount)
	}

	removeIfExists(output)
}

func TestCopyToOutputDontCopyHeader(t *testing.T) {
	input := "test/file1.csv"
	output := "test/tmp_123.csv"

	removeIfExists(output)

	err := copyToOutput(input, output, false)
	if err != nil {
		t.Fatal(err)
	}

	lineCount, err := fileLineCount(output)
	if err != nil {
		t.Fatal(err)
	}

	if lineCount != 4 {
		t.Fatal("Error while copyToOutput, got unexpected line count:", lineCount)
	}

	removeIfExists(output)
}

// Helpers

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func fileLineCount(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := f.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func removeIfExists(path string) {
	os.Remove(path)
}
