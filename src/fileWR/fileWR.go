package fileWR

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"structs"
)

func ReadFile(path string) structs.Geojson {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Printf("read file error: %v\n", e)
		os.Exit(1)
	}
	var geojson structs.Geojson
	json.Unmarshal(file, &geojson)
	return geojson
}

func WiteFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		log.Printf("write file error: %v\n", err)
		os.Exit(1)
	}
}

func ReadFolder(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("dir error: %v\n", err)
		os.Exit(1)
	}
	return files
}
