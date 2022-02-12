package main

import (
	"bhg-scanner/scanner"
	"fmt"
	"time"
)

// creates a slice with numbers 1 through 1024
func allPorts() []int {
	var ports []int
	for i := 1; i <= 1024; i++ {
		ports = append(ports, i)
	}
	return ports
}

func main() {
	target := "scanme.nmap.org"
	output := false   //set to true to print
	file := "out.csv" //file name for output (MUST BE CSV)

	start := time.Now()

	scanner.PortScanner(file, target, allPorts(), output)

	elapsed := time.Since(start)
	fmt.Printf("Port Scanner took %v\n", elapsed)
}
