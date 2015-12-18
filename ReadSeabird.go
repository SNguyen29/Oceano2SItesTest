//ReadSeabird.go
//Function for read data file for constructor Seabird

package main


// read cnv files in two pass, the first to get dimensions
// second to get data
func (nc *Nc) ReadSeabird(files []string,optCfgfile string) {
	
	var Instrument string
	var Type string
	Instrument = AnalyzeTypeSeabird(files)
	Type = AnalyzeTypeInstrument(Instrument)
	var cfg Config
	
	switch{
		case Instrument == CTD :

			nc.GetConfigCTD(optCfgfile,cfg,Type)
			// first pass, return dimensions fron cnv files
			nc.Dimensions["TIME"], nc.Dimensions["DEPTH"] = nc.firstPassCTD(files)
		
			// initialize 2D data
			nc.Variables_2D = make(AllData_2D)
			for i, _ := range map_var {
				nc.Variables_2D.NewData_2D(i, nc.Dimensions["TIME"], nc.Dimensions["DEPTH"])
			}
		
			// second pass, read files again, extract data and fill slices
			nc.secondPassCTD(files)
			// write ASCII file
			nc.WriteAsciiCTD(map_format, hdr,cfg,Instrument)
		
			// write netcdf file
			//if err := nc.WriteNetcdf(); err != nil {
			//log.Fatal(err)
			//}
			nc.WriteNetcdf(Instrument)
			
		case Instrument == BTL :
		
			nc.GetConfigBTL(optCfgfile,cfg,Type)
			// first pass, return dimensions fron btl files
			nc.Dimensions["TIME"], nc.Dimensions["DEPTH"] = nc.firstPassBTL(files)
		
			//	// initialize 2D data
			//	nc.Variables_2D = make(AllData_2D)
			//	for i, _ := range map_var {
			//		nc.Variables_2D.NewData_2D(i, nc.Dimensions["TIME"], nc.Dimensions["DEPTH"])
			//	}
		
			// second pass, read files again, extract data and fill slices
			nc.secondPassBTL(files)
			// write ASCII file
			nc.WriteAsciiBTL2(map_format, hdr,cfg,Instrument)
		
			// write netcdf file
			//if err := nc.WriteNetcdf(); err != nil {
			//log.Fatal(err)
			//}
			nc.WriteNetcdf(Instrument)
			}
}