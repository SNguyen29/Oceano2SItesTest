// GetProfileNumber_test.go
package main

import "testing"

import (
	"fmt"
)

//function for testing GetProfileNumber
func TestGetProfile(t *testing.T){
// variable for test

var ncTest Nc
ncTest.TestInitNC()

TestFile := "Data/csp00101.cnv"
Profile := ncTest.GetProfileNumber(TestFile)
fmt.Println("Profile Number : ",Profile)

}