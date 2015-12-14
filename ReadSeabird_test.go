//ReadSeabird_test.go

package main

import "testing"

//function for testing ReadGlobal
func  TestReadSeabird(t *testing.T) {

TestFile := []string{"Data/csp00101.cnv"}	
var ncTest Nc
ncTest.TestInitNC()
var config Config
TestFile, optCfgfile := GetOptions()
ncTest.GetConfigCTD(optCfgfile,config)	
ncTest.ReadSeabird(TestFile,optCfgfile)
}