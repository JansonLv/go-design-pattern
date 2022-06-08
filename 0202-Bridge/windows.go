package main

import "fmt"

type windows struct {
	printer Printer
}

func (m *windows) print() {
	fmt.Println("windows print")
	m.printer.printFile()
}

func (m *windows) setPrinter(printer Printer) {
	m.printer = printer
}
