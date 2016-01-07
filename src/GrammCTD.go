//GrammCTD.go
//File for with the regular expression for CTD type instrument and function for read CTD files

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"regexp"
	"strconv"
	"strings"

)


//regular expressions
var regIsHeader = regexp.MustCompile(`^[*#]`)
//var regEndOfHeader = regexp.MustCompile(`\*END\*`)

//function
// read .cnv files and return dimensions
func (nc *Nc) firstPassCTD(files []string) (int, int) {
	
	//variable init
	var	pres float64 = 0
	var depth float64 = 0
	var	maxDepth float64 = 0
	var maxPres	float64 = 0
	var maxPresAll float64 = 0
	var line int = 0
	var maxLine int = 0
	
	fmt.Fprintf(echo, "First pass: ")
	// loop over each files passed throw command line
	for _, file := range files {
		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()

		profile := nc.GetProfileNumber(file)
		scanner := bufio.NewScanner(fid)
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeader.MatchString(str)
			if !match {
				values := strings.Fields(str)
				// read the pressure
				if pres, err = strconv.ParseFloat(values[map_var["PRES"]], 64); err != nil {
					log.Fatal(err)
				}
				// read the depth
				if depth, err = strconv.ParseFloat(values[map_var["DEPTH"]], 64); err != nil {
					log.Fatal(err)
				} else {
					//p(math.Floor(depth))
				}
			}
			if pres > maxPres {
				maxPres = pres
				maxDepth = depth
				line = line + 1
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Fprintf(debug, "Read %s size: %d max pres: %4.f\n", file, line, maxPres)

		if line > maxLine {
			maxLine = line
		}
		// store the maximum pressure and maximum depth value per cast
		nc.Extras_f[fmt.Sprintf("PRES:%d", int(profile))] = maxPres
		nc.Extras_f[fmt.Sprintf("DEPTH:%d", int(profile))] = math.Floor(maxDepth)
		nc.Extras_s[fmt.Sprintf("TYPECAST:%s", int(profile))] = "n/a"
		if maxPres > maxPresAll {
			maxPresAll = maxPres
		}
		// reset value for next loop
		maxPres = 0
		maxDepth = 0
		pres = 0
		line = 0
	}

	fmt.Fprintf(echo, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxPresAll)
	fmt.Fprintf(debug, "First pass: %d files read, maximum pressure found: %4.0f db\n", len(files), maxPresAll)
	fmt.Fprintf(debug, "First pass: size %d x %d\n", len(files), maxLine)
	return len(files), maxLine
}

func (nc *Nc) secondPassCTD(files []string) {

	fmt.Fprintf(echo, "Second pass ...\n")

	// initialize profile and pressure max
	var nbProfile int = 0

	// loop over each files passed throw command line
	for _, file := range files {
		var line int = 0

		fid, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fid.Close()
		// fmt.Printf("Read %s\n", file)

		profile := nc.GetProfileNumber(file)
		scanner := bufio.NewScanner(fid)
		downcast := true
		for scanner.Scan() {
			str := scanner.Text()
			match := regIsHeader.MatchString(str)
			if match {
				nc.DecodeHeaderSeabird(str, profile)
			} else {
				// fill map data with information contain in read line str
				nc.DecodeDataSeabird(str, profile, file, line)

				if downcast {
					// fill 2D slice
					for _, key := range hdr {
						if key != "PRFL" {
							//fmt.Println("Line: ", line, "key: ", key, " data: ", data[key])
							nc.Variables_2D[key].data[nbProfile][line] = data[key].(float64)
						}
					}
					// exit loop if reach maximum pressure for the profile
					if data["PRES"] == nc.Extras_f[fmt.Sprintf("PRES:%d", int(profile))] {
						downcast = false
					}
				} else {
					// store last julian day for end profile
					nc.Extras_f[fmt.Sprintf("ETDD:%d", int(profile))] = data["ETDD"].(float64)
					//fmt.Println(presMax)
				}
				line++
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// increment sclice index
		nbProfile += 1

		// store last julian day for end profile
		nc.Extras_f[fmt.Sprintf("ETDD:%d", int(profile))] = data["ETDD"].(float64)
		//fmt.Println(presMax)
	}
	fmt.Fprintln(debug, nc.Variables_1D["PROFILE"])
}