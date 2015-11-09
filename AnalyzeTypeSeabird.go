//AnalyzeTypeSeabird.go
package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// define constante for instrument type
type InstrumentType int

const (
	CTD InstrumentType = 3
	BTL                = 5
)

// define regexp
var regIsCnv = regexp.MustCompile(`(\*END\*)`)
var regIsBtottle = regexp.MustCompile(`^\s+(Bottle)`)

// read all cnv files and return dimensions
func AnalyzeTypeSeabird(files []string) InstrumentType {

	// initialize result
	var result InstrumentType = 0

	// open first file
	fid, err := os.Open(files[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fid.Close()

	scanner := bufio.NewScanner(fid)
	for scanner.Scan() {
		str := scanner.Text()
		
		switch{
		case regIsCnv.MatchString(str) :
			result = 3
		
		case regIsBtottle.MatchString(str) :
			result = 5
		}
	}
	return result
}
