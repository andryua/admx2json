package main

import (
	"admx/helpers"
	"html/template"
	"log"
	"net/http"
)

var res []helpers.AllPolicies

func admjson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(helpers.Treegen(res)))
}
func GPTree(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/gptree.html"))
	var v = make(map[string]interface{})
	v["UnChangeable"] = 1
	//v["Token"] = token
	t.ExecuteTemplate(w, "gptree", v)
}
func main() {
	dataPolicies, lang, dataCat, cataloguesName, present := helpers.ParseFiles()
	//fmt.Printf("%v\n",dataCat)
	cataloguePath := helpers.CategoriesPath(dataCat, cataloguesName)
	res = helpers.PoliciesParse(dataPolicies, lang, cataloguePath, present)
	//jsonRes, err := json.Marshal(res)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//file1 := "gpo.json"
	//ioutil.WriteFile(file1, jsonRes, 0777)
	//jsonTree := helpers.Treegen(res)
	//file2 := "gpoTree.json"
	//ioutil.WriteFile(file2, []byte(jsonTree), 0777)

	http.HandleFunc("/admjson", admjson)
	http.HandleFunc("/", GPTree)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))) //погашення папки (щоб при роботі сервера він знав де брати файли для вебу)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("Error on ListenAndServe:\n")
		log.Fatal(err.Error())
	}

}
