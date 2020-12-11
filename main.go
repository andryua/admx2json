package main

import (
	"admx/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	n, lang := helpers.ParseFiles()
	catpath := helpers.CategoriesPath(n, lang)
	res := helpers.PoliciesParse(n, lang, catpath)

	jsonres, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fil := "gpo.json"
	ioutil.WriteFile(fil, jsonres, 0777)
}
