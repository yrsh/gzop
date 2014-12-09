package main

import (
	"configParser"
	"fileWR"
	"flag"
	"jsonParser"
	"log"
	"os"
	"strconv"
	//"strings"
	"sync"
)

var wg sync.WaitGroup

var conffile = flag.String("conf", "", "config file")

func main() {
	//
	flag.Parse()
	conf := configParser.GetConf(*conffile)
	//
	log.Println("porcessing JSONs")
	files := fileWR.ReadFolder(conf.Folders.Input)
	for i := conf.Zooms.Min; i <= conf.Zooms.Max; i++ {
		for j := range files {
			wg.Add(1)
			go func(f string, i int) {
				jsonRaw := fileWR.ReadFile(conf.Folders.Input + "/" + f)
				jsonS, name := jsonParser.ProcessJSON(jsonRaw, i, conf.Naming.Feature_prop)
				os.MkdirAll(conf.Folders.Output+"/"+strconv.Itoa(i), 0777)
				fileWR.WiteFile(conf.Folders.Output+"/"+
					strconv.Itoa(i)+"/"+name, jsonS)
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
