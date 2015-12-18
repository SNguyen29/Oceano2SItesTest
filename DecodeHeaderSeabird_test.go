// DecodeHeaderSeabird_test.go
package main

import "testing"

//Function for init before the test
func (ncTest *Nc) TestInitNC(){
	
		// define map from netcdf structure
	ncTest.Dimensions = make(map[string]int)
	ncTest.Attributes = make(map[string]string)
	ncTest.Extras_f = make(map[string]float64)
	ncTest.Extras_s = make(map[string]string)
	ncTest.Variables_1D = make(map[string]interface{})

	// initialize map entry from nil interface to empty slice of float64
	ncTest.Variables_1D["PROFILE"] = []float64{}
	ncTest.Variables_1D["TIME"] = []float64{}
	ncTest.Variables_1D["LATITUDE"] = []float64{}
	ncTest.Variables_1D["LONGITUDE"] = []float64{}
	ncTest.Variables_1D["BATH"] = []float64{}
	//	nc.Variables_1D["TYPECAST"] = []float64{}
	ncTest.Roscop = codeRoscopFromCsv(code_roscop)

	// add some global attributes for profile, change in future
	ncTest.Attributes["data_type"] = "OceanSITES profile data"
}

//function for testing Decodeheader 
func TestDecodeHeader(t *testing.T){
// variable for test

var ncTest Nc
ncTest.TestInitNC()
var profileTest float64 = 00101

//var StringTest string = "* System UpLoad Time = Jul 20 2015 06:15:13"
//var StringTest string = "* NMEA Latitude = 19 58.55 S"
var StringTest string = "* NMEA Longitude = 168 00.45 E"
//var StringTest string = "** Cruise : CASSIOPEE"
//var StringTest string = "** Station : 00101"
//var StringTest string = "** Bottom Depth: 4937"
//var StringTest string = "** Operator:  JG-GE-ESL"
ncTest.DecodeHeaderSeabird(StringTest,profileTest)	
}