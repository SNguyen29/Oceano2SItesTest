//Analyse.go
//global analyse of the file

package main

type structfile struct{
	Constructeur 	Constructor
	Instrument		string
	TypeInstrument	string	
}

func AnalyzeFile(files []string) structfile{
	
	var result structfile
	
	result.Constructeur = AnalyzeConstructor(files)
	
	result.Instrument = AnalyzeTypeSeabird(files)
	
	result.TypeInstrument = AnalyzeTypeInstrument(result.Instrument)
	
	return result
}