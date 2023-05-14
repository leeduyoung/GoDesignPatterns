package main

import "fmt"

func main() {
	// Create files
	file1 := NewTextFile("file1.txt", 100)
	file2 := NewTextFile("file2.txt", 200)
	file3 := NewTextFile("file3.txt", 300)

	// Create Directories
	subdir1 := NewDirectory("subdir1")
	subdir2 := NewDirectory("subdir2")

	// Add files to directories
	subdir1.AddFile(file1)
	subdir1.AddDirectory(subdir2)
	subdir2.AddFile(file2)

	root := NewDirectory("root")
	root.AddFile(file3)
	root.AddDirectory(subdir1)

	fmt.Println(root.Size())
}
