//AnalyzeTypeSeabird.go
//Analyze the type of instrument of the data files
package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// read all cnv files and return dimensions
func AnalyzeTypeSeabird(files []string) string {

	var result = "error in analyse type"
	
	// open first file
	fid, err := os.Open(files[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fid.Close()

	scanner := bufio.NewScanner(fid)
	for scanner.Scan() {
		str := scanner.Text()
		
		for i:=0;i<len(cfg.Instrument.Decodetype);i++{
				temp := regexp.MustCompile(cfg.Instrument.Decodetype[i])
				if temp.MatchString(str){
						result = cfg.Instrument.Type[i]
					}
			
			}
	}
	
	return result
}
