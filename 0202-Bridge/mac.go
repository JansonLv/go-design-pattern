package main

import "fmt"

type mac struct {
	printer Printer
}

func (m *mac) print() {
	fmt.Println("mac print")
	m.printer.printFile()
}

func (m *mac) setPrinter(printer Printer) {
	m.printer = printer
}
