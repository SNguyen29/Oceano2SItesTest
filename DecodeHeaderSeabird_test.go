// DecodeHeaderSeabird_test.go
package main

import "testing"

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// define regexp
var TestregIsHeader = regexp.MustCompile(`^[*#]`)
var TestregEndOfHeader = regexp.MustCompile(`\*END\*`)
var TestregCruise = regexp.MustCompile(`Cruise\s*:\s*(.*)`)
var TestregShip = regexp.MustCompile(`Ship\s*:\s*(.*)`)
var TestregStation = regexp.MustCompile(`Station\s*:\s*\D*(\d*)`)
var TestregType = regexp.MustCompile(`Type\s*:\s*(.*)`)
var TestregOperator = regexp.MustCompile(`Operator\s*:\s*(.*)`)
var TestregBottomDepth = regexp.MustCompile(`Bottom Depth\s*:\s*(\d*\.?\d+?)\s*\S*`)
var TestregDummyBottomDepth = regexp.MustCompile(`Bottom Depth\s*:\s*$`)
var TestregDate = regexp.MustCompile(`Date\s*:\s*(\d+)/(\d+)/(\d+)`)
var TestregHour = regexp.MustCompile(`[Heure|Hour]\s*:\s*(\d+)[:hH](\d+):(\d+)`)
var TestregLatitude = regexp.MustCompile(`Latitude\s*:\s*(\d+)\s+(\d+.\d+)\s+(\w)`)
var TestregLongitude = regexp.MustCompile(`Longitude\s*:\s*(\d+)\s+(\d+.\d+)\s+(\w)`)
var TestregSystemTime = regexp.MustCompile(`System UpLoad Time =\s+(.*)`)
var TestregNmeaLatitude = regexp.MustCompile(`NMEA Latitude\s*=\s*(\d+\s+\d+.\d+\s+\w)`)
var TestregNmeaLongitude = regexp.MustCompile(`NMEA Longitude\s*=\s*(\d+\s+\d+.\d+\s+\w)`)

// parse header line from .cnv and extract correct information
// use regular expression
// to parse time with non standard format, see:
// http://golang.org/src/time/format.go

func (ncTest *Nc) TestDecodeHeaderSeabird(str string, profileTest float64) {
	switch {
	// decode Systeme Upload Time
		case TestregSystemTime.MatchString(str) : 
			res := TestregSystemTime.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintf(debug, "%s -> ", value)
			// create new Time object, see tools.go
			var t = NewTimeFromString("Jan 02 2006 15:04:05", value)
			v := t.Time2JulianDec()
			ncTest.Variables_1D["TIME"] = append(ncTest.Variables_1D["TIME"].([]float64),v)
	
		case TestregNmeaLatitude.MatchString(str) :
			if v, err := Position2Decimal(str); err == nil {
				ncTest.Variables_1D["LATITUDE"] = append(ncTest.Variables_1D["LATITUDE"].([]float64), v)
			} else {
				ncTest.Variables_1D["LATITUDE"] = append(ncTest.Variables_1D["LATITUDE"].([]float64), 1e36)
			}
			
			
		case TestregNmeaLongitude.MatchString(str) :
			if v, err := Position2Decimal(str); err == nil {
				ncTest.Variables_1D["LONGITUDE"] = append(ncTest.Variables_1D["LONGITUDE"].([]float64), v)
				fmt.Println(v)
			} else {
				ncTest.Variables_1D["LONGITUDE"] = append(ncTest.Variables_1D["LONGITUDE"].([]float64), 1e36)
			}
			
			
		case TestregCruise.MatchString(str) :
			res := TestregCruise.FindStringSubmatch(str)
			value := res[1]
			fmt.Println(value)
			fmt.Fprintln(debug, value)
			ncTest.Attributes["cycle_mesure"] = value

		case TestregShip.MatchString(str) :
			res := TestregShip.FindStringSubmatch(str)
			value := res[1]
			fmt.Fprintln(debug, value)
			ncTest.Attributes["plateforme"] = value
			fmt.Println(value)
			
		case TestregStation.MatchString(str) :
			res := TestregStation.FindStringSubmatch(str)
			value := res[1]
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				fmt.Fprintln(debug, v)
				// ch
				//			format := "%0" + cfg.Ctd.StationPrefixLength + ".0f"
				//			p := fmt.Sprintf(format, profile)
				//			//s := fmt.Sprintf(format, v)
				//			fmt.Println(p, v)
				//			if p != v {
				//				fmt.Printf("Warning: profile for header differ from file name: %s <=> %s\n", p, v)
				//			}
				ncTest.Variables_1D["PROFILE"] = append(ncTest.Variables_1D["PROFILE"].([]float64), profileTest)
			} else {
				ncTest.Variables_1D["PROFILE"] = append(ncTest.Variables_1D["PROFILE"].([]float64), 1e36)
			}
			fmt.Println(value)

		case TestregBottomDepth.MatchString(str) :
			res := TestregBottomDepth.FindStringSubmatch(str)
			value := res[1]
			if v, err := strconv.ParseFloat(value, 64); err == nil {
				fmt.Fprintf(debug, "Bath: %f\n", v)
				ncTest.Variables_1D["BATH"] = append(ncTest.Variables_1D["BATH"].([]float64), v)
			} else {
				fmt.Fprintf(debug, "Bath: %f\n", v)
				ncTest.Variables_1D["BATH"] = append(ncTest.Variables_1D["BATH"].([]float64), 1e36)
			}
			fmt.Println(value)
			
		case TestregDummyBottomDepth.MatchString(str) ://?
			ncTest.Variables_1D["BATH"] = append(ncTest.Variables_1D["BATH"].([]float64), 1e36)
			fmt.Fprintf(debug, "Bath: %g\n", 1e36)
			fmt.Println("1e36")

		case TestregOperator.MatchString(str) :
			res := TestregOperator.FindStringSubmatch(str)
			value := res[1]
			if *optDebug {
				fmt.Println(value)
			}
			fmt.Println(value)
			
	// TODOS: uncomment, add optionnal value from seabird header
		case TestregType.MatchString(str) :
			res := TestregType.FindStringSubmatch(str)
			value := strings.ToUpper(res[1]) // convert to upper case
			var v float64
			switch value {
			case "PHY":
				v = float64(PHY)
			case "GEO":
				v = float64(GEO)
			case "BIO":
				v = float64(BIO)
			default:
				v = float64(UNKNOW)
			}
			//f("Type: %f\n", v)
			ncTest.Variables_1D["TYPECAST"] = append(ncTest.Variables_1D["TYPECAST"].([]float64), v)
			if *optDebug {
				fmt.Println(value)
			}
			ncTest.Extras_s[fmt.Sprintf("TYPE:%d", int(profileTest))] = value
		
	}
}


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
ncTest.TestDecodeHeaderSeabird(StringTest,profileTest)	
}