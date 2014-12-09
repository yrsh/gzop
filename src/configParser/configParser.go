package configParser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Conf struct {
	Naming  Naming
	Zooms   Zooms
	Folders Folders
}

type Naming struct {
	Original     bool
	Feature_prop string
}

type Zooms struct {
	Min int
	Max int
}

type Folders struct {
	Rules  bool
	Input  string
	Output string
}

func GetConf(path string) Conf {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("config file error: %v\n", e)
		os.Exit(1)
	}
	var jsontype Conf
	json.Unmarshal(file, &jsontype)
	return jsontype
}
