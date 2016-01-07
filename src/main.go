
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	
) 

const PROGNAME string = "Test"
const PROGVERSION string = "0.1.0"

type Data_2D struct {
	data [][]float64
}

type AllData_2D map[string]Data_2D

type Nc struct {
	Dimensions   map[string]int
	Variables_1D map[string]interface{}
	Variables_2D AllData_2D
	Attributes   map[string]string
	Extras_f     map[string]float64 // used to store max of profiles value
	Extras_s     map[string]string  // used to store max of profiles type
	Roscop       map[string]RoscopAttribute
}


// configuration file
var cfgname string = "ini/oceano2oceansites.ini"
var code_roscop string = "code_roscop.csv"

// file prefix for --all option: "-all" for all parameters, "" empty by default
var prefixAll = ""

// usefull macro
var p = fmt.Println
var f = fmt.Printf

// Create an empty map.
var map_var = map[string]int{}
var map_format = map[string]string{}
var data = make(map[string]interface{})
var hdr []string
var cfg Config

// use for debug mode
var debug io.Writer = ioutil.Discard

// use for echo mode
var echo io.Writer = ioutil.Discard

var nc Nc

// main body
func main() {
	
	var files []string
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if os.Getenv("ApplicationTest") != "" {
		cfgname = os.Getenv("ApplicationTest")
	}
	if os.Getenv("ROSCOP") != "" {
		code_roscop = os.Getenv("ROSCOP")
	}
	
	files, optCfgfile := GetOptions()

	Cons := AnalyzeConstructor(files)
	
	switch{
		case Cons == Seabird :
			nc.ReadSeabird(files,optCfgfile)
		}
	
}
