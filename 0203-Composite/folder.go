package main

import "fmt"

type folder struct {
	name       string
	components []component
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching for keyword:%s in folder:%s \n", keyword, f.name)
	for _, component := range f.components {
		component.search(keyword)
	}
}
