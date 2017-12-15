# CSV Concat

CSV Concat combines 1+ text files into a single output file. The intended file type is CSV files, but any text file will work as long as first row is header and rest of rows are non-header

Caveats:

* Does NOT validate your CSV files
* Overwrites the file at the output path
* Assumes all files have a header row
* Assumes no row is longer than 65,536 characters

## Arguments

Required

	* input relative file path(s)
	* output relative file path

## Setup

	# Installs to Command Line via your $GOPATH/bin/
	make

## Usage

	csvConcat --in '[INPUT FILES]' --out [OUTPUT_FILE]

## Example

	csvConcat --in 'test/file1.csv test/file2.csv test/file3.csv' --out combined.csv

## Help

	csvConcat -h

## Run Tests

	make test
