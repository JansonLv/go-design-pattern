package main

type computer interface {
	print()
	setPrinter(Printer)
}

type Printer interface {
	printFile()
}
