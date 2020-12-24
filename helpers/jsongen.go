package helpers

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

//var CachedJson string

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}
func Treegen(dataGPO []AllPolicies) string {

	//if CachedJson != "" {
	//	return CachedJson
	//}
	var dpath []string
	var fpath []string
	for _, pol := range dataGPO {
		if strings.Contains(pol.DisplayName, "/") {
			strings.ReplaceAll(pol.DisplayName, "/", "|")
		}
		if pol.Class == "Both" {
			x := "User" + "/" + pol.Category + "/"
			y := "Machine" + "/" + pol.Category + "/"
			dpath = append(dpath, x)
			dpath = append(dpath, y)
			xx := "User" + "/" + pol.Category + "/" + strconv.Itoa(pol.ID)
			yy := "Machine" + "/" + pol.Category + "/" + strconv.Itoa(pol.ID)
			fpath = append(fpath, xx)
			fpath = append(fpath, yy)
		} else {
			x := pol.Class + "/" + pol.Category + "/"
			dpath = append(dpath, x)
			y := pol.Class + "/" + pol.Category + "/" + strconv.Itoa(pol.ID)
			fpath = append(fpath, y)
		}
	}

	var web []string
	var dr []string
	var drr []string
	dpath = removeDuplicates(dpath)
	sort.Strings(dpath)
	fpath = removeDuplicates(fpath)
	sort.Strings(fpath)
	//fmt.Println(dpath)
	for i := 0; i < len(dpath); i++ {
		s := strings.Split(dpath[i], "/")
		ss := ""
		for j := 0; j < len(s); j++ {
			if j == len(s)-1 {
				ss += s[j]
			} else {
				ss += s[j] + "/"
			}
			if strings.Contains(dpath[i], ss) {
				drr = append(drr, ss)
			}
		}
	}
	//dr = removeDuplicates(dpath)
	dr = removeDuplicates(drr)
	sort.Strings(dr)
	//sort.Strings(fp)

	web = append(web, "[{")
	web = append(web, "\"title\": \"Administartive Templates\",")
	web = append(web, "\"folder\": true,", "\"expanded\": true,")
	web = append(web, "\"children\": [")

	sl1 := 0
	for i := 0; i < len(dr); i++ {
		fmt.Println(dr[i])
		sl := strings.Split(dr[i], "/")
		if i+1 == len(dr) {
			sl1 = len(sl)
		} else {
			sl1 = len(strings.Split(dr[i+1], "/"))
		}
		tmp := ""
		tmp1 := ""
		for z := 0; z < len(sl)-1; z++ {
			tmp += "	"
		}
		for z := 0; z < sl1-1; z++ {
			tmp1 += "	"
		}
		web = append(web, tmp+"{")
		web = append(web, tmp+"\"title\": "+"\""+sl[len(sl)-2]+"\",")
		if dr[i] == "User/" || dr[i] == "Machine/" {
			web = append(web, "\"expanded\": true,")
		}
		web = append(web, tmp+"\"folder\": true,"+"\"children\": [")

		for _, pol := range dataGPO {
			//s1 := strings.Join(sl[1:], "/")
			s1, _ := filepath.Split(dr[i])
			if last := len(s1) - 1; last >= 0 && s1[last] == '/' {
				s1 = s1[1:last]
			}
			splt := strings.Split(s1, "/")
			splt1 := strings.Join(splt[1:], "/")
			//x,_ := strconv.Atoi(file)
			if strings.EqualFold(sl[0], pol.Class) || strings.EqualFold(pol.Class, "Both") {
				if strings.EqualFold(splt1, pol.Category) { //&& x == pol.ID {
					//fmt.Println(dr[i])
					web = append(web, tmp+"	{")
					web = append(web, tmp+"\"id\": "+"\""+strconv.Itoa(pol.ID)+"\",")
					dn, _ := json.Marshal(pol.DisplayName)
					web = append(web, tmp+"\"title\": "+string(dn)+",")
					sup, _ := json.Marshal(pol.SupportedOn)
					web = append(web, tmp+"\"support\": "+string(sup)+",")
					desc, _ := json.Marshal(pol.ExplainText)
					web = append(web, tmp+"\"description\": "+string(desc)+",")
					web = append(web, tmp+"\"category\": \""+pol.Category+"\",")
					web = append(web, tmp+"\"class\": \""+sl[0]+"\",")
					web = append(web, tmp+"\"values\": ")
					x, err := json.Marshal(pol.Values)
					if err != nil {
						fmt.Println(err)
						return ""
					}
					web = append(web, tmp+"	"+string(x))
					web = append(web, tmp+" },")
				}
			} else {

				continue
			}
		}

		if sl1-len(sl) != 1 {
			web = append(web, tmp+"] },")
		}

		m := len(tmp) - len(tmp1)
		if m > 0 {
			x := "	"
			for c := 0; c < m; c++ {
				x = tmp[0 : len(tmp)-m]
				web = append(web, " "+x+"] },")
			}
		}
	}
	web = append(web, "		] },")
	web = append(web, "	] }")
	web = append(web, "] }")
	web = append(web, "]")
	for i := 0; i < len(web); i++ {
		if strings.Contains(web[i], "]") { //&& strings.Contains(web[i],"},") {
			web[i-1] = strings.Replace(web[i-1], "},", "}", -1)
		}
	}
	//str := ""
	//for i := 0; i < len(web); i++ {
	//	web[i] = strings.Replace(web[i], "	", " ", -1)
	//	str += web[i] + "\n"
	//}
	//CachedJson = str
	return strings.Join(web, "\n") //str
}
