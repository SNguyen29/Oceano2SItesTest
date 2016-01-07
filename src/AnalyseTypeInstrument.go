//AnalyzeTypeInstrument.go
//Analyze the type of instrument in general case
package main

import (
	"strings"
)

// define constante for Type

const (
	Profil string = "profile"
	TimeSeries string = "time-series"
	Trajectoire string = "trajectory"
)

var	TabProfil = [2]string{CTD,BTL}
var	TabTimeSeries = [0]string{}
var	TabTrajectoire = [0]string{}


func AnalyzeTypeInstrument(inst string) string{
	
	var Vprofil bool = false
	var VTimes bool = false
	var Vtraj bool = false
	var result string
	
	for i:=0;i<len(TabProfil);i++{
			if strings.EqualFold(inst,TabProfil[i]){
				Vprofil = true
			}
				
		}
	for i:=0;i<len(TabTimeSeries);i++{
			if strings.EqualFold(inst,TabTimeSeries[i]){
				VTimes = true
			}
				
		}
	for i:=0;i<len(TabTrajectoire);i++{
			if strings.EqualFold(inst,TabTrajectoire[i]){
				Vtraj = true
			}
				
		}
		
	switch{
		case Vprofil : result = Profil
		case VTimes : result = TimeSeries
		case Vtraj : result = Trajectoire
	}
	
	return result
}
