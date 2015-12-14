// GetProfileNumber.go
//Function for get the profil number of a data file for Seabird Constructor
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"regexp"
	"os"
	"bufio"
)

var regprefix = regexp.MustCompile(`cruisePrefix\s*=\s*(.*)`)

func (nc *Nc) GetProfileNumber(str string) float64 {
	var value float64
	var err error
	var CruisePrefix string = GetCruisePrefix()
	if strings.Contains(str,CruisePrefix) {
		res := strings.Split(str,CruisePrefix)
		res = strings.Split(res[1],".")
		if value, err = strconv.ParseFloat(res[0], 64); err == nil {
			// get profile name, eg: csp00101
			nc.Extras_s[fmt.Sprintf("PRFL_NAME:%d", int(value))] = res[0]
		} else {
			log.Fatal(err)
		}

	} else {
		log.Fatal("func GetProfileNumber", err)
	}
	return value

}

func GetCruisePrefix() string{

	var prefix string
	file, err := os.Open(cfgname) // For read access.
		if err != nil {
			log.Fatal(err)
			}
		defer file.Close()
	scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			str := scanner.Text()
			if regprefix.MatchString(str){
				res := regprefix.FindStringSubmatch(str)
				prefix = res[1]
				fmt.Println(prefix)
				}
			}

	return prefix
}