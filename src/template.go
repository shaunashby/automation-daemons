package main

import (
	"html/template"
	"os"
)

// DFiClientID is a type to hold client information:
type DFiClientID struct {
	dfinum string
}

// Dfinum method returns the client ID:
func (d DFiClientID) Dfinum() (s string) {
	return d.dfinum
}

// DFiCleanIPEntry is a type to hold information about an IP to be provisined:
type DFiCleanIPEntry struct {
	IPAddress string
	*DFiClientID
}

func templateTest() {
	DFiClient := DFiClientID{dfinum: "123456"}

	cleanips := DFiCleanIPEntry{"192.168.2.3", &DFiClient}

	tmpl, err := template.New("cleanip").Parse("\"{{.IPAddress}}\"; # Client {{.DFiClientID.Dfinum}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, cleanips)
	if err != nil {
		panic(err)
	}
}
