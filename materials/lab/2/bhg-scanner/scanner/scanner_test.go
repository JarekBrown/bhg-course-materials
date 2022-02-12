package scanner

import (
	"testing"
)

func allPorts() ([]int) {
    var ports []int
    for i := 1; i <= 1024; i++ {
        ports = append(ports, i)
    }
    return ports
}

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T) {
	filename := "out.csv"
	target := "scanme.nmap.org"

	got, _ := PortScanner(filename, target, allPorts(), false) // Currently function returns only number of open ports
	want := 2                             // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns
	//consider what would happen if you parameterize the portscanner address and ports to scan

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTotalPortsScanned(t *testing.T) {
	filename := "out.csv"
	target := "scanme.nmap.org"

	open, closed := PortScanner(filename, target, allPorts(), false) // Currently function returns only number of open ports
	want := 1024                                // default value; consider what would happen if you parameterize the portscanner ports to scan

	got := open + closed

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestClosedPort(t *testing.T) {
	filename := "out.csv"
	target := "scanme.nmap.org"

	_, closed := PortScanner(filename, target, allPorts(), false)
	want := 1022

	if closed != want {
		t.Errorf("got %d, wanted %d", closed, want)
	}
}

func TestSomePorts(t *testing.T) {
	filename := "out.csv"
    target := "scanme.nmap.org"

    ports := make([]int, 2)
    ports[0] = 80
    ports[1] = 500

    open,closed := PortScanner(filename, target, ports, false)

    wantOpen := 1
    if open != wantOpen {
        t.Errorf("got %d, wanted %d", open, wantOpen)
    }

    wantClosed := 1
    if closed != wantClosed {
        t.Errorf("got %d, wanted %d", closed, wantClosed)
    }
}
