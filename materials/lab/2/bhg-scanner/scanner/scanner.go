// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Usage:
// PortScanner uses go channels and DialTimeout from the net package to scan a specified target address/ports
// Ex.
// PortScanner("output_file.csv",127.0.0.1, [21,22,80,443], true))
// ^ this would scan ports 21, 22, 80, and 443 on 127.0.0.1 and output the results to "output_file.csv"

package scanner

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// Ports contains the results of a port scan
type Ports struct {
	open    int            // number of open ports
	closed  int            // number of closed ports
	status  map[int]string // stores the status of scanned ports
	targets []int          // target ports
}

// newStatus adds the port status to the status map in a Ports object
// the count of open and closed ports is updated
func (ports *Ports) newStatus(status bool, num int) {
	if status == true {
		ports.open++
		ports.status[num] = "OPEN"
	} else {
		ports.closed++
		ports.status[num] = "closed"
	}
}

// writes port status out to a csv file
func (ports *Ports) toCSV(filename string) {
	file, err := os.Create(filename) // create (or overwrite) an output file
	if err != nil {
		log.Println("could not create file")
	}
	defer file.Close()

	_, err = file.WriteString("port,status\n") // write header for file
	if err != nil {
		log.Fatalln("error header to file: ", err)
	}

	for i := 1; i <= len(ports.status); i++ {
		file.WriteString(fmt.Sprintf("%v,%v\n", i, ports.status[i])) // write each port status
	}
}

func worker(ports, results chan int, target string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			results <- (p * -1)
			continue
		}
		conn.Close()
		results <- p
	}
}

// scans ports of the specified target address (target ports as well), writing to specified file when output is true
// returns the number of open ports and the number of closed ports
func PortScanner(fname, address string, targetPorts []int, output bool) (int, int) {

	ports := make(chan int, 100) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	var scanned Ports
	scanned.targets = targetPorts
	scanned.status = make(map[int]string)

	// create workers and scan the target
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}
	go func() {
		for i := range scanned.targets {
			ports <- targetPorts[i]
		}
	}()

	// read the status of each scanned port
	for i := 0; i < len(scanned.targets); i++ {
		port := <-results
		if port >= 0 {
			scanned.newStatus(true, port) // port is open
		} else {
			scanned.newStatus(false, (port * -1)) // port is closed
		}
	}
	close(ports)
	close(results)

	// if output desired, print to the csv file
	if output {
		scanned.toCSV(fname)
	}

	return scanned.open, scanned.closed
}
