package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// if not enough args provided
	if len(os.Args) < 2 {
		syntax()
		return
	}

	// if help menu is needed
	if os.Args[1] == "-h" {
		help()
		return
	}

	var p []int
	var ipIndex int
	if os.Args[1] == "-p" || os.Args[1] == "-P" {
		p = ports(os.Args[1], os.Args[2])
		ipIndex = 3
	} else {
		p = ports("", "")
		ipIndex = 1
	}
	target := os.Args[ipIndex]
	valid, validTarget := validTarget(target)
	if valid == false {
		fmt.Println("error: target ip/address invalid")
		return
	}
	scan(validTarget, p)
}

func scan(target string, ports []int) {
	dt := time.Now()
	fmt.Printf("Starting GoMap at %s\n", dt.Format(time.UnixDate))
	fmt.Printf("Scan Report for %s :\n", target)
	fmt.Printf("PORT\tSTATE\n")

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", target, port)
		_, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			fmt.Printf("%d\t%s\t\n", port, "open")
		} else {
			fmt.Printf("%d\t%s\t\n", port, "closed")
		}
	}
}

// check if the target input is a valid format
func validTarget(target string) (bool, string) {
	validIP := net.ParseIP(target)
	validURL, err := url.Parse(target)
	host := validURL.Host
	if (err != nil || host == "") && validIP == nil {
		return false, ""
	} else if validIP == nil { // if it is in URL format use host
		return true, host
	} else {
		return true, target
	}
}

// returns ports according to the option specified
func ports(option, input string) []int {
	p := extractNumbers(input)
	var ports []int

	if option == "-p" {
		return p
	} else if option == "-P" {
		start := p[0]
		end := p[1]
		if end < start {
			fmt.Println("error: end value for port range smaller than start range")
		}
		ports = generatePortRange(start, end)
	} else {
		ports = generatePortRange(1, 1024)
	}
	return ports
}

// returns port numbers in input string
func extractNumbers(input string) []int {
	// use regex to find requested ports
	re := regexp.MustCompile("[0-9]+")
	p := re.FindAllString(input, -1)
	var ports []int

	// convert []string p to []int ports
	for _, b := range p {
		tmp, _ := strconv.Atoi(b)
		ports = append(ports, tmp)
	}
	return ports
}

// creates array
func generatePortRange(start, end int) []int {
	var ports []int
	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}
	return ports
}

// syntax menu
func syntax() {
	fmt.Println("Usage: gomap [options] [target ip/address]")
	fmt.Println("See 'gomap -h' for usage information")
}

// help menu
func help() {
	fmt.Print("\nGo Implementation of nmap\n\n")
	fmt.Println("Available options:")
	fmt.Println("  -h: see this help window")
	fmt.Println("  -p: specify port(s) to be scanned (default without this flag is [1:1024]")
	fmt.Println("  -P: specify a range of ports to be scanned")
	fmt.Println("Examples:")
	fmt.Println("  gomap -p 80 https://scanme.nmap.org")
	fmt.Println("  gomap -p 80,22,42 https://scanme.mnmap.org")
	fmt.Println("  gomap -P 1-100 https://scanme.nmap.org")
	fmt.Println("NOTE: if using a URL, a scheme (http/https) is required! (i.e. https://example.com)")
}
