//ConfigTHERMO.go
//File for config a instrument THERMO

package main

import (
	//"code.google.com/p/gcfg"
	//"fmt"
	//"log"
	//"strconv"
	//"strings"
)

type thermo struct {

	CruisePrefix        string
	StationPrefixLength string
	TypeInstrument      string
	InstrumentNumber    string
	TitleSummary        string
	
}

func (nc *Nc) GetConfigTHERMO(configFile string,cfg Config,Type string) {

}