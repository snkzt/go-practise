package main

import (
	"fmt"
	"os"
)

// Defer is used to ensure that a function call is performed later in a program's execution,
// usually to clean up.
func main() {
	f := createFile("./defer.txt")
	// Defer will be executed at the end of the enclosing function, main in this case,
	// after all the non-defer function, in this case witeFile, has finished.
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
