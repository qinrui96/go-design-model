package pkg

import (
	"strconv"
	"strings"
)

type Item interface {
	Count() int
}

type File struct {
	Name string
}

func (f *File) Count() int {
	return 1
}

type Folder struct {
	items []Item
}

func (f *Folder) Count() int {
	var accumulate int
	for _, item := range f.items {
		accumulate += item.Count()
	}
	return accumulate
}

func (f *Folder) AddItem(i Item) {
	f.items = append(f.items, i)
}

func NewFolder() *Folder {
	folder := &Folder{}
	for i := 0; i < 10; i++ {
		folder.AddItem(&File{Name: strconv.Itoa(i)})
		fileNameElement := []string{strconv.Itoa(i), "child"}
		fileName := strings.Join(fileNameElement, "")
		folder.AddItem(&Folder{items: []Item{&File{Name: fileName}}})
	}
	return folder
}
