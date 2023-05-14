package main

type Directory struct {
	name    string
	files   []*File
	subdirs []*Directory
}

func NewDirectory(name string, files []*File, subdirs []*Directory) *Directory {
	return &Directory{
		name:    name,
		files:   files,
		subdirs: subdirs,
	}
}

func (d Directory) Name() string {
	return d.name
}

func (d Directory) Files() []*File {
	return d.files
}

func (d Directory) Directories() []*Directory {
	return d.subdirs
}

func (d Directory) Size() int {
	size := 0

	for _, file := range d.files {
		size += file.Size()
	}

	for _, subdir := range d.subdirs {
		size += subdir.Size()
	}

	return size
}
