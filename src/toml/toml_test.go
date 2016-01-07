//test toml

package main

import (
	"fmt"
	"testing"
	"github.com/BurntSushi/toml"
	//"regexp"
	"strings"
)

var file = "filetoml.toml"

type configtoml struct{
	Progname	string
	Progversion	string
	Configfile 	string
	Roscopfile 	string
	Seabird struct{
		Cnv			string
		Bottle		string
		Prefix		string
		Header1		string
		Header2		string
		Split1		string
		Split2		string
		Cruise 		string
		Ship 		string
		Station 	string
		Type 		string
		Operator	string
		BottomDepth string
		Date 		string
		Hour 		string
		SystemTime 	string
		Latitude 	string
		Longitude 	string
		}
}
	
var cc configtoml

func  TestTOML(t *testing.T) {
	
	//  read config file
	if _, err := toml.DecodeFile(file, &cc); err != nil {
		fmt.Println(err)
		return
	}
	var re = (cc.Seabird.Header1)
	var temp string = "** Cruise : CASSIOPEE"
	//var temp2 string = "* NMEA Latitude = 19 58.55 S"
	
	
	if(strings.Contains(temp,re)){
		fmt.Println("ok")
		ter := strings.Split(temp,": ")
		fmt.Println(ter[1])

		}else{
			fmt.Println("NO")
			}
	
}