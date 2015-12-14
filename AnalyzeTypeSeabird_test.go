//AnalyzeType_test.go
package main

import "testing"

import (
	"fmt"
)

//function for testing ReadGlobal
func  TestAnalyzeTypeSeabird(t *testing.T) {
	
TestFile := []string{"Data/csp00101.cnv"}
	
Type := AnalyzeTypeSeabird(TestFile)

fmt.Println("Type = ",Type)

}
