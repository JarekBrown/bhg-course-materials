package main

import (
	"fmt"
	"hscan/hscan"
	"os"
)

func main() {

	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	//* I used rockyou.txt to find these
	drmike1 := "90f2c9c53f66540e67349e0ab83d8cd0"                                 //p@ssword
	drmike2 := "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032" //letmein

	// NON CODE - TODO
	// Download and use bigger password file from: https://weakpass.com/wordlist/tiny  (want to push yourself try /small ; to easy? /big )

	if len(os.Args) < 2 {
		fmt.Println("usage: main [filename]")
		return
	}
	//TODO Grab the file to use from the command line instead; look at previous lab (e.g., #3 ) for examples of grabbing info from command line
	file := os.Args[1]

	hscan.GuessSingle(md5hash, file)
	hscan.GuessSingle(sha256hash, file)
	hscan.GuessSingle(drmike1, file)
	hscan.GuessSingle(drmike2, file)

	hscan.GenHashMaps(file)

	hscan.GetSHA(sha256hash)
	hscan.GetMD5(sha256hash)
	hscan.GetSHA(drmike2)
	hscan.GetMD5(drmike1)
}
