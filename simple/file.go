package simple

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}

func (file *File) Close() {
	fmt.Println("Close File", file.Name)
}
