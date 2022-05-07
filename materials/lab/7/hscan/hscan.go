package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string
var mutext = &sync.Mutex{}
var mutextword = &sync.Mutex{}

func GuessSingle(sourceHash string, filename string) {
	begin := time.Now()

	n := len(sourceHash)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// add a check and logicial structure
		if n == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
		} else if n == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		} else {
			fmt.Println("error: unknown hash length")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	done := time.Now()
	elapsed := done.Sub(begin)
	fmt.Printf("Duration for %s: %v\n", sourceHash, elapsed.Milliseconds())
}

func GenHashMaps(filename string) {

	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	shalookup = make(map[string]string)
	md5lookup = make(map[string]string)

	for scan.Scan() {
		pass := scan.Text()

		go addto(pass)
	}
	if err := scan.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

func GetMD5(hash string) (string, error) {
	pass, itBeOkey := md5lookup[hash]

	if itBeOkey {
		return pass, nil
	} else {
		return "", errors.New("password does not exist")
	}
}

func addto(pass string) {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))
	mutext.Lock()
	shalookup[hash] = pass
	mutext.Unlock()
	hash = fmt.Sprintf("%x", md5.Sum([]byte(pass)))
	mutextword.Lock()
	md5lookup[hash] = pass
	mutextword.Unlock()
}
