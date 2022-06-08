package main

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	windowsComputer := &windows{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	windowsComputer.setPrinter(epsonPrinter)
	windowsComputer.print()
}
