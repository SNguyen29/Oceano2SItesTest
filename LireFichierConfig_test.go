package main

import (
	"code.google.com/p/gcfg"
	"fmt"
	"log"
	"strings"
	"testing"
)


func TestLireConfig(t *testing.T) {
	
	var cfg Config
	var split string
	
	err := gcfg.ReadFileInto(&cfg,"Ini/configTest.ini")
	if err == nil {
		fmt.Println("aucune erreur readfileinfo")
		fmt.Println(cfg.Ctd.CruisePrefix)
		fmt.Println(cfg.Ctd.InstrumentNumber)
		fmt.Println(cfg.Ctd.StationPrefixLength)
		fmt.Println(cfg.Ctd.TitleSummary)
		fmt.Println(cfg.Ctd.TypeInstrument)
		
		split = cfg.Ctd.Split
		
		var fields []string
		fields = strings.Split(split, ",")
		fmt.Println(fields)
		
		
	}else {
		fmt.Println("erreur readfileinfo")
		fmt.Println(err)
		log.Fatal(err)
	}
	
		
}