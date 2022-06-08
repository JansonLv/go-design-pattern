package main

import "fmt"

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(name string) {
	fmt.Println(name + f.name)
	for _, i := range f.children {
		i.print(name)
	}
}

func (f *folder) clone() inode {
	children := make([]inode, 0, len(f.children))
	for _, v := range f.children {
		children = append(children, v.clone())
	}
	return &folder{name: f.name, children: children}
}
