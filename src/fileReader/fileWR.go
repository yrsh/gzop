package fileWR

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"structs"
)

func readFile(path string) structs.Geojson {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Printf("config file error: %v\n", e)
		os.Exit(1)
	}
	var geojson structs.Geojson
	json.Unmarshal(file, &geojson)
	return jsontype
}

func witeFile(path string) structs.Geojson {
	//
}

func readFolder() {
	//
}
