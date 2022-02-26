// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"shodan/shodan"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: main <ports|search>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	if os.Args[1] == "ports" {
		portList(s)
	} else{
		search(s)
	}
}

// I would pipe this out to some junk file, it's a long list
func portList(client *shodan.Client){
	ports, err := client.ListPorts()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Ports them crawlerz be looking for:")
	for _, port := range ports {
		fmt.Print(string(port))
	}
}

func search(client *shodan.Client) {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main search <searchterm>")
	}
	hostSearch, err := client.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ",host.IPString,"====")
		h,_ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ",host.IPString,"====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}


	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}
}