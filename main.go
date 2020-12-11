package main

import (
	"admx/helpers"
	"encoding/json"
	"io/ioutil"

	//"encoding/json"
	"fmt"
	//"io/ioutil"
)

func main() {
	//data, lang, dataCat, cataloguesname := helpers.ParseFiles()
	data, lang, dataCat, cataloguesname := helpers.ParseFiles()
	//fmt.Printf("%v",dataCat)
	catpath := helpers.CategoriesPath(dataCat, cataloguesname)
	//fmt.Printf("%v", catpath)
	res := helpers.PoliciesParse(data, lang, catpath)

	jsonres, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fil := "gpo.json"
	ioutil.WriteFile(fil, jsonres, 0777)
}
