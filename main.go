package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdout, "required file path")
		helpPrint()
		os.Exit(1)
	}

	filePath := os.Args[1]

	err := readFromFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read from file, err: %s\n", err.Error())
		helpPrint()
		os.Exit(1)
	}
}

func readFromFile(path string) error {
	if path != "" {
		var inputReader io.Reader

		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("error opening file %s: %v", path, err)
		}
		defer file.Close()
		inputReader = file
		fmt.Printf("reading from file: %s\n", path)

		scanner := bufio.NewScanner(inputReader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error opening file %s: %v", path, err)
		}
	}

	return nil
}

func helpPrint() {
	fmt.Fprintln(os.Stdout, "Usage: parking_app <file_path>")
}
