//ConfigLADCP.go
//File for config a instrument LADCP

package main

import (
	//"code.google.com/p/gcfg"
	//"fmt"
	//"log"
	//"strconv"
	//"strings"
)

type ladcp struct {

	CruisePrefix        string
	StationPrefixLength string
	TypeInstrument      string
	InstrumentNumber    string
	TitleSummary        string
	
	
}

func (nc *Nc) GetConfigLADCP(configFile string,cfg Config,Type string) {

}