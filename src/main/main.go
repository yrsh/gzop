package main

import (
	"fileWR"
	"flag"
	"jsonParser"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var minZoom = flag.Int("minz", 3, "min zoom")
var maxZoom = flag.Int("maxz", 18, "max zoom")
var jsonFolder = flag.String("inp", "", "input folder with jsons")
var outFolder = flag.String("out", "", "output folder for jsons")

func main() {
	flag.Parse()
	//
	log.Println("porcessing JSONs")
	files := fileWR.ReadFolder(*jsonFolder)
	for i := *minZoom; i <= *maxZoom; i++ {
		for j := range files {
			wg.Add(1)
			go func(f string, i int) {
				jsonRaw := fileWR.ReadFile(*jsonFolder + "/" + f)
				jsonS := jsonParser.ProcessJSON(jsonRaw, i)
				os.MkdirAll(*outFolder+"/"+strconv.Itoa(i), 0777)
				fileWR.WiteFile(*outFolder+"/"+
					strconv.Itoa(i)+"/"+strings.Replace(f, ".geojson", ".json", -1), jsonS)
				defer func() {
					wg.Done()
				}()
			}(files[j].Name(), i)
		}
	}
	wg.Wait()
	defer func() {
		log.Println("processing complete")
	}()
}
