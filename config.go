//config.go
//File with struct config who need to be change when a new type of instrument is add

package main

type Config struct {
	Global struct {
		Author string
		Debug  bool
		Echo   bool
	}
	Cruise struct {
		CycleMesure string
		Plateforme  string
		Callsign    string
		Institute   string
		Pi          string
		Timezone    string
		BeginDate   string
		EndDate     string
		Creator     string
	}
	Ctd ctd
	Btl btl
	Xbt xbt
	Thermo thermo
	Ladcp ladcp
	Sadcp sadcp
	//add new type of instrument
}

