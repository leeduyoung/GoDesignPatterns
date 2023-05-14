package main

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func (f File) Name() string {
	return f.name
}

func (f File) Size() int {
	return f.size
}
