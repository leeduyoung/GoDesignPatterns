package main

import "fmt"

func main() {
	// Create files
	file1 := NewFile("file1.txt", 100)
	file2 := NewFile("file2.txt", 200)
	file3 := NewFile("file3.txt", 300)

	// Create Directories
	subdir1 := NewDirectory("subdir1", []*File{file1}, []*Directory{})
	subdir2 := NewDirectory("subdir2", []*File{file2}, []*Directory{subdir1})

	// Create Root Directory
	root := NewDirectory("root", []*File{file3}, []*Directory{subdir2})

	fmt.Println(root.Size())
}
