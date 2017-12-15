package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const newline = "\n"

type config struct {
	inFilesRaw,
	outFile string

	inFiles []string
}

var cfg = config{
	inFilesRaw: "",
	outFile:    "",

	inFiles: []string{},
}

func init() {
	flag.StringVar(&cfg.inFilesRaw, "in", cfg.inFilesRaw, "CSV file paths to be combined, paths separated by a space")
	flag.StringVar(&cfg.outFile, "out", cfg.outFile, "Combined output file path, will overwrite output location")
}

func main() {
	flag.Parse()
	parseInFiles()

	first := firstInFile(cfg.inFiles)
	copyToOutput(first, cfg.outFile, true)

	latter := latterInFiles(cfg.inFiles)
	for _, f := range latter {
		copyToOutput(f, cfg.outFile, false)
	}

	fmt.Println("Success: CSV Concat completed.")
}

// parseInFiles splits the inFilesRaw into inFiles based on spaces
func parseInFiles() {
	if cfg.inFilesRaw == "" {
		err := errors.New("Error: no input files")
		abort(err)
	}

	cfg.inFiles = strings.Split(cfg.inFilesRaw, " ")
}

// firstInFile returns the first file in cfg.inFiles
func firstInFile(files []string) string {
	if len(files) > 0 {
		return files[0]
	}

	return ""
}

// latterInFiles returns all except for the first file in cfg.inFiles
func latterInFiles(files []string) []string {
	if len(files) > 1 {
		return files[1:]
	}

	return []string{}
}

// copyToOutput copies the source CSV to the target CSV destination by appending to the target file,
// takes into account if it should copy the header (first) row as well
func copyToOutput(source, target string, copyHeader bool) error {
	fOut, err := os.OpenFile(target, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fOut.Close()

	fIn, err := os.Open(source)
	if err != nil {
		return err
	}
	defer fIn.Close()

	firstRow := true

	scanner := bufio.NewScanner(fIn)
	for scanner.Scan() {
		append := true

		if firstRow {
			if !copyHeader {
				append = false
			}
			firstRow = false
		}

		if append {
			if _, err = fOut.WriteString(scanner.Text() + newline); err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// abort outputs error & exits
func abort(e error) {
	fmt.Println(e)
	os.Exit(255)
}
