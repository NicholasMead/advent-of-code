package mockfs

import (
	"strings"
)

type FileSystem interface {
	Move(dir string) error

	CurrentDir() Directory
	FindAll(selector func(Directory) bool) <-chan Directory
}

type fs struct {
	root Directory
	path []string
}

// CurrentDir implements FileSystem
func (fs *fs) CurrentDir() Directory {
	dir := fs.root

	for _, p := range fs.path {
		if sub, err := dir.GetSubDirectory(p); err != nil {
			panic(err)
		} else {
			dir = sub
		}
	}

	return dir
}

// FindAll implements FileSystem
func (f *fs) FindAll(selector func(Directory) bool) <-chan Directory {
	output := make(chan Directory)

	go func() {
		defer close(output)
		queue := []Directory{f.root}

		for {
			switch {
			case len(queue) == 0:
				return
			case selector(queue[0]):
				output <- queue[0]
			}

			queue = append(queue, queue[0].GetSubDirectories()...)
			queue = queue[1:]
		}
	}()

	return output
}

// Move implements FileSystem
func (fs *fs) Move(path string) error {
	switch {
	case path == "":
		return nil

	case path[0] == '/':
		fs.path = []string{}
		fs.Move(path[1:])

	case path == "..":
		fs.path = fs.path[:len(fs.path)-1]

	case strings.Contains(path, "/"):
		subPaths := strings.Split(path, "/")
		for _, sub := range subPaths {
			fs.Move(sub)
		}

	default:
		if _, err := fs.CurrentDir().GetSubDirectory(path); err != nil {
			return err
		} else {
			fs.path = append(fs.path, path)
		}
	}

	return nil
}

func CreateFileSystem() FileSystem {
	return &fs{CreateDirectory(""), []string{}}
}
