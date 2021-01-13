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
	dataPolicies, lang, dataCat, cataloguesName, present := helpers.ParseFiles()
	//fmt.Printf("%v\n",dataCat)
	cataloguePath := helpers.CategoriesPath(dataCat, cataloguesName)
	res := helpers.PoliciesParse(dataPolicies, lang, cataloguePath, present)
	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	file1 := "gpo.json"
	ioutil.WriteFile(file1, jsonRes, 0777)
	jsonTree := helpers.Treegen(res)
	file2 := "gpoTree.json"
	ioutil.WriteFile(file2, []byte(jsonTree), 0777)
	/*
		for key, value := range present {
			fmt.Println(key, " = ", value)
			fmt.Println("*************************************************************")
		}
	*/
}
