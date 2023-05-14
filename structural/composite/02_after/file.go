package main

type File interface {
	Name() string
	Size() int
}

type TextFile struct {
	name string
	size int
}

func NewTextFile(name string, size int) *TextFile {
	return &TextFile{
		name: name,
		size: size,
	}
}

func (f TextFile) Name() string {
	return f.name
}

func (f TextFile) Size() int {
	return f.size
}
