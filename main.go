package main

import (
	"admx/helpers"
	"encoding/json"
	"io/ioutil"
	"strings"

	//"encoding/json"
	"fmt"
	//"io/ioutil"
)

func main() {
	dataPolicies, lang, dataCat, cataloguesname := helpers.ParseFiles()
	cpDataCat := make(map[string]string)
	for key, value := range dataCat {
		if strings.Contains(value, ":") {
			value = strings.Split(value, ":")[1]
		}
		cpDataCat[key] = value
		//fmt.Print(key,",",value,"\n")
	}
	catpath := helpers.CategoriesPath(cpDataCat, cataloguesname)
	res := helpers.PoliciesParse(dataPolicies, lang, catpath)

	jsonres, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fil := "gpo.json"
	ioutil.WriteFile(fil, jsonres, 0777)
}
