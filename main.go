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
	dataPolicies, lang, dataCat, cataloguesName := helpers.ParseFiles()
	//fmt.Printf("%v\n",dataCat)
	cataloguePath := helpers.CategoriesPath(dataCat, cataloguesName)
	res := helpers.PoliciesParse(dataPolicies, lang, cataloguePath)

	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	file1 := "gpo.json"
	file2 := "gpoTree.json"
	jsonTree := helpers.Treegen(res)
	//helpers.Treegen(res)
	ioutil.WriteFile(file1, jsonRes, 0777)
	ioutil.WriteFile(file2, []byte(jsonTree), 0777)
}
