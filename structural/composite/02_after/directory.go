package main

type Directory struct {
	name    string
	files   []File
	subdirs []*Directory
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func (d Directory) Name() string {
	return d.name
}

func (d Directory) Files() []File {
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

func (d *Directory) AddFile(file File) {
	d.files = append(d.files, file)
}

func (d *Directory) AddDirectory(subdir *Directory) {
	d.subdirs = append(d.subdirs, subdir)
}
