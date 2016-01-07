//AnalyzeTypeSeabird.go
//Analyze the type of instrument of the data files
package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// define constante for instrument type
 
const (
	CTD 	string = "CTD"
	BTL 	string = "BTL"
	LADCP 	string = "LADCP"
	SADCP	string = "SADCP"
	THERMO 	string = "THERMO"
	XBT		string = "XBT"
)

// define regexp
var regIsCnv = regexp.MustCompile(`(\*END\*)`)
var regIsBottle = regexp.MustCompile(`^\s+(Bottle)`)

// read all cnv files and return dimensions
func AnalyzeTypeSeabird(files []string) string {

	// initialize result
	var result string = ""

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
			result = CTD
		
		case regIsBottle.MatchString(str) :
			result = BTL
		//add case for other instrument	
		}	
	}
	return result
}
