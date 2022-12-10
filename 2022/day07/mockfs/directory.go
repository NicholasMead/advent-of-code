package mockfs

import "errors"

type Directory interface {
	AddFile(name string, size int)
	AddDirectory(dir Directory)

	GetName() string
	GetFileSize(name string) (int, error)
	GetSubDirectory(name string) (Directory, error)
	GetSubDirectories() []Directory
	GetSize() int

	
}

type directory struct {
	name   string
	files  map[string]int
	subDir map[string]Directory
}

func (d *directory) GetSubDirectories() []Directory {
	output := []Directory{}

	for _, d := range d.subDir {
		output = append(output, d)
	}

	return output
}

func (d *directory) GetSize() int {
	size := 0
	for _, f := range d.files {
		size += f
	}
	for _, d := range d.subDir {
		size += d.GetSize()
	}
	return size
}

func (d *directory) GetName() string {
	return d.name
}

func (d *directory) AddDirectory(dir Directory) {
	d.subDir[dir.GetName()] = dir
}

func (d *directory) GetSubDirectory(name string) (Directory, error) {
	if dir, found := d.subDir[name]; !found {
		return nil, errors.New("Directory not found")
	} else {
		return dir, nil
	}
}

func CreateDirectory(name string) Directory {
	return &directory{
		name:   name,
		files:  make(map[string]int),
		subDir: make(map[string]Directory),
	}
}

func (d *directory) AddFile(name string, size int) {
	d.files[name] = size
}

func (d *directory) GetFileSize(name string) (int, error) {
	if file, found := d.files[name]; !found {
		return -1, errors.New("File not found")
	} else {
		return file, nil
	}
}
