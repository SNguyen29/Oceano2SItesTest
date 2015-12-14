// DecodeDataSeabird_test.go
package main

import "testing"

import (
	"log"
	"os"
	"bufio"
)


//function for testing Decodedata
func TestDecodeData(t *testing.T){
// variable for test

var ncTest Nc
ncTest.TestInitNC()

TestFile := "Data/FileTestDecodeData.cnv"

var profileTest float64 = 00101

	var line int = 0

	fid, err := os.Open(TestFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fid.Close()
	// fmt.Printf("Read %s\n", file)

	scanner := bufio.NewScanner(fid)
	for scanner.Scan() {
		str := scanner.Text()
		ncTest.DecodeDataSeabird(str,profileTest,TestFile,line)
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}