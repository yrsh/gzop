package main

import (
	"configParser"
	"fileWR"
	"flag"
	"jsonParser"
	"log"
	"os"
	"ruleParser"
	"strconv"
	//"strings"
	//"sync"
)

//var wg sync.WaitGroup

var conffile = flag.String("conf", "", "config file")
var rulesfile = flag.String("rules", "", "rules file")

func main() {
	//
	flag.Parse()
	conf := configParser.GetConf(*conffile)
	rules := ruleParser.GetConf(*rulesfile)
	log.Print(rules["SLV"])
	//
	/*
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
	*/
	//
	for k, v := range rules {
		for j := range v {
			for i := v[j].Range[0]; i <= v[j].Range[1]; i++ {
				jsonRaw := fileWR.ReadFile(conf.Folders.Input + "/" + v[j].File)
				jsonS, _ := jsonParser.ProcessJSON(jsonRaw, i, conf.Naming.Feature_prop)
				os.MkdirAll(conf.Folders.Output+"/"+strconv.Itoa(i), 0777)
				fileWR.WiteFile(conf.Folders.Output+"/"+
					strconv.Itoa(i)+"/"+k+".json", jsonS)
			}
		}
	}
	//
}
