// AnalyzeConstructor_test.go
package main

import "testing"

import (
	"fmt"
)

//function for testing ReadGlobal
func  TestAnalyzeConstructor(t *testing.T) {

TestFile := []string{"csp00101.cnv"}
	
Construct := AnalyzeConstructor(TestFile)

fmt.Println("Numero Constructeur = ",Construct)

}
