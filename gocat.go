package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:

Example:
	gocat file1 file2 file3 ...

`, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, filePath := range flag.Args() {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file: %q\nError: %v",
				filePath, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintf(os.Stderr, "Error while printin file: %q\nError: %v",
				filePath, err)
		}

		if err = file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error while closing file: %q\nError: %v",
				filePath, err)
		}
	}
}
