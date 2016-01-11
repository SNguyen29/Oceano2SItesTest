//AnalyzeConstructor.go
//Analyze the constructor of the data files
package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// define constante for constructor type
type Constructor int

const (
	Seabird Constructor = 1
)


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
		case strings.ContainsAny(cfg.Instrument.Seabird,str) : 
			result = Seabird
		}
	}
	return result
}

