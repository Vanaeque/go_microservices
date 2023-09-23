package main

import "fmt"

type Storage[T any] struct {
	storageType string
	MapItems    map[int]T
	SliceItems  []T
}

func NewStorage[itemType any](storageType string) (error, *Storage[itemType]) {
	if storageType != "map" && storageType != "slice" {
		return fmt.Errorf("Incorrect type"), &Storage[itemType]{}
	} else {
		s := Storage[itemType]{}
		s.MapItems = make(map[int]itemType)
		s.SliceItems = make([]itemType, 1000, 1000)
		s.storageType = storageType
		return nil, &s
	}
}

func (s *Storage[any]) AddItem(id int, item any) {
	if s.storageType == "map" {
		s.MapItems[id] = item
	} else {
		s.SliceItems[id] = item
	}
}

func (s *Storage[any]) GetItem(id int) any {
	if s.storageType == "map" {
		return s.MapItems[id]
	} else {
		return s.SliceItems[id]
	}
}

type Book struct {
	title  string
	author string
	text   string
}

type Library struct {
	name    string
	genFunc func(string) int
	Storage *Storage[Book]
}

func NewLibrary(name string, genFunc func(string) int, storage *Storage[Book]) *Library {
	return &Library{name, genFunc, storage}
}

func (l *Library) AddBook(b Book) {
	id := l.genFunc(b.title)
	l.Storage.AddItem(id, b)
}

func (l *Library) GetBook(title string) Book {
	id := l.genFunc(title)
	return l.Storage.GetItem(id)
}

func someFunc(s string) int {
	return len(s)
}

func main() {
	_, storage1 := NewStorage[Book]("map")
	lib1 := NewLibrary("somelib1", someFunc, storage1)

	lib1.AddBook(Book{"Recepies", "Granny", "blabla"})
	lib1.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})
	lib1.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})
	lib1.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})
	lib1.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})

	fmt.Println(lib1.GetBook("Recepies"))
	fmt.Println(lib1.GetBook("Recepies. Pt.4"))

	_, storage2 := NewStorage[Book]("slice")
	lib2 := NewLibrary("somelib1", someFunc, storage2)

	lib2.AddBook(Book{"Recepies", "Granny", "blabla"})
	lib2.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})
	lib2.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})
	lib2.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})
	lib2.AddBook(Book{"Recepies. Pt.2", "Granny", "blabla"})

	fmt.Println(lib2.GetBook("Recepies"))
	fmt.Println(lib2.GetBook("Recepies. Pt.4"))
}
