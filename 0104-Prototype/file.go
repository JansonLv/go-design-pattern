package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(name string) {
	fmt.Println(name + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name}
}
