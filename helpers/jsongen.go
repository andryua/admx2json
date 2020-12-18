package helpers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var CachedJson string

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

	if CachedJson != "" {
		return CachedJson
	}
	var dpath []string
	var fpath []string
	for _, pol := range dataGPO {
		if strings.Contains(pol.DisplayName, "/") {
			strings.ReplaceAll(pol.DisplayName, "/", "-")
		}
		if pol.Class == "Both" {
			x := "User" + "/" + strings.ReplaceAll(pol.Category, "|", "/") + "/"
			y := "Machine" + "/" + strings.ReplaceAll(pol.Category, "|", "/") + "/"
			dpath = append(dpath, x)
			dpath = append(dpath, y)
			xx := "User" + "/" + strings.ReplaceAll(pol.Category, "|", "/") + "/" + pol.DisplayName
			yy := "Machine" + "/" + strings.ReplaceAll(pol.Category, "|", "/") + "/" + pol.DisplayName
			fpath = append(fpath, xx)
			fpath = append(fpath, yy)
		} else {
			x := pol.Class + "/" + strings.ReplaceAll(pol.Category, "|", "/") + "/"
			dpath = append(dpath, x)
			y := pol.Class + "/" + strings.ReplaceAll(pol.Category, "|", "/") + "/" + pol.DisplayName
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
			//if j == len(s)-1 {
			//	ss += s[j]
			//} else {
			ss += s[j] + "/"
			//	}
			if strings.Contains(dpath[i], ss) {
				drr = append(drr, ss)
			}
		}
	}
	dr = removeDuplicates(drr)
	sort.Strings(dr)
	//fmt.Println(dr)

	web = append(web, "[{")
	web = append(web, "\"title\": \"Administartive Templates\",")
	web = append(web, "\"folder\": true,", "\"expanded\": true,")
	web = append(web, "\"children\": [")

	//sl1 := 0
	//kid :=0
	k := 0
	for i := 0; i < len(dr); i++ {
		sl := strings.Split(dr[i], "/")
		//	if i+1 == len(dr) {
		//		sl1 = len(sl)
		//	} else {
		//		sl1 = len(strings.Split(dr[i+1], "/"))
		//	}
		tmp := ""
		//tmp1 := ""
		web = append(web, tmp+"{")
		web = append(web, tmp+"\"title\": "+"\""+sl[len(sl)-2]+"\",")
		if strings.EqualFold(dr[i], "User/") || strings.EqualFold(dr[i], "Machine/") {
			web = append(web, "\"expanded\": true,")
		}
		web = append(web, tmp+"\"folder\": true,"+"\"children\": [")

		for _, pol := range dataGPO {
			//dir, file := filepath.Split(dr[i])
			//fmt.Println("dr[i]: ",dr[i]," dir: ",dir, " sl[len(sl)-2]: ", sl[len(sl)-2])
			tm := strings.Split(dr[i], "/")
			//if len(dr[i]) <= 2 {
			//	tm =
			//}
			s1 := strings.Join(tm[1:], "/")
			if last := len(s1) - 1; last >= 0 && s1[last] == '/' {
				s1 = s1[:last]
			}
			s2 := strings.ReplaceAll(pol.Category, "|", "/")
			//fmt.Println("s1: ", s1," pol.cat: ",pol.Category)
			if strings.EqualFold(s1, s2) && (strings.EqualFold(tm[0], pol.Class) || strings.EqualFold(pol.Class, "Both")) { // && strings.EqualFold(file,pol.DisplayName) {
				k++
				web = append(web, tmp+"	{")
				web = append(web, tmp+"\"id\": "+"\""+strconv.Itoa(k)+"\",")
				dn, _ := json.Marshal(pol.DisplayName)
				web = append(web, tmp+"\"title\": "+string(dn)+",")
				sup, _ := json.Marshal(pol.SupportedOn)
				web = append(web, tmp+"\"support\": "+string(sup)+",")
				desc, _ := json.Marshal(pol.ExplainText)
				web = append(web, tmp+"\"description\": "+string(desc)+",")
				web = append(web, tmp+"\"category\": \""+s2+"\",")
				web = append(web, tmp+"\"class\": \""+pol.Class+"\",")
				web = append(web, tmp+"\"values\": ")
				x, err := json.Marshal(pol.Values)
				if err != nil {
					fmt.Println(err)
					return ""
				}
				web = append(web, tmp+"	"+string(x))
				web = append(web, tmp+" },")
			} else {
				continue
			}
		}

		//if sl1-len(sl) != 1 {
		web = append(web, tmp+"] },")
		//}

		/*m := len(tmp) - len("")
		if m > 0 {
			x := "	"
			for c := 0; c < m; c++ {
				x = tmp[0 : len(tmp)-m]
				web = append(web, " "+x+"] },")
			}
		}*/
	}
	web = append(web, "] },")
	/*
		web = append(web, "] }")
		web = append(web, "] }")
		web = append(web, "] }")
	*/
	web = append(web, "]")
	for i := 0; i < len(web); i++ {
		if strings.Contains(web[i], "]") { //&& strings.Contains(web[i],"},") {
			web[i-1] = strings.Replace(web[i-1], "},", "}", -1)
		}
	}
	str := ""
	for i := 0; i < len(web); i++ {
		web[i] = strings.Replace(web[i], "	", " ", -1)
		str += web[i] + "\n"
	}
	CachedJson = str
	return str
}
