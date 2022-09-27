package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create a new sub-directory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	check(err)
	// Defer the removal of a directories when dreating a temporary one.
	// os.RemoveAll will delete a whole directory tree (similar to rm -rf)
	defer os.RemoveAll("subdir")
	// Helper function
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")
	// Create a hierarchy of directories
	// MkdirAll is similar to the command-line mkidr -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")
	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// Chdir changes the current working directory like cd.
	err = os.Chdir("subdir/parent/child")
	check(err)
	c, err = os.ReadDir(".") // subdir/parent/child when list the current directory
	check(err)
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..") // cd back
	check(err)
	fmt.Println("Visiting subdir")
	// Visit a directory recursively, including all its sub-directories.
	err = filepath.Walk("subdir", visit)
}

// This function is called for every file or diredctory found recursively
// by filepath.Walk
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
