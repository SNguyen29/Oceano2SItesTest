//AnalyzeType_test.go
package main

import "testing"

import (
	"fmt"
)

//function for testing ReadGlobal
func  TestAnalyzeTypeSeabird(t *testing.T) {
	
TestFile := []string{"data/csp00101.cnv"}
	
Type := AnalyzeTypeSeabird(TestFile)

fmt.Println("Type = ",Type)

TestFile2 := []string{"data/csp00101.btl"}
	
Type2 := AnalyzeTypeSeabird(TestFile2)

fmt.Println("Type = ",Type2)


}
