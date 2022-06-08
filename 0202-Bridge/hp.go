package main

import "fmt"

type hp struct {
}

func (e *hp) printFile() {
	fmt.Println("hp printFile")
}
