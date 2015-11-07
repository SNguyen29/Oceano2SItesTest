// AnalyzeConstructor.go
package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// define constante for instrument type
type Constructor int

const (
	Seabird Constructor = 1
)

// define regexp
var regIsSeabird = regexp.MustCompile(`^\*\s+(Sea-Bird)`)

// read all cnv files and return dimensions
func AnalyzeConstructor(files []string) Constructor {

	var result Constructor
	// open first file
	fid, err := os.Open(files[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fid.Close()

	scanner := bufio.NewScanner(fid)
	for scanner.Scan() {
		str := scanner.Text()
		
		switch {
		case regIsSeabird.MatchString(str) : 
			result = Seabird
		}
	}
	return result
}

