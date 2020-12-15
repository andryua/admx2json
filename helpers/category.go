package helpers

import (
	"strings"
)

var keyPath = make(map[string]string)

func catPath(key, value string) string {
	path := key
	if value == "" || strings.EqualFold(value, keyPath[value]) {
		return path
	}
	path += "|" + value
	return catPath(path, keyPath[value])
}

func reverse(item []string) []string {
	newItem := make([]string, len(item))
	for i, j := 0, len(item)-1; i <= j; i, j = i+1, j-1 {
		newItem[i], newItem[j] = item[j], item[i]
	}
	return newItem
}

func CategoriesPath(keyPath, catname map[string]string) map[string]string {
	cpKey := make(map[string]string)
	for key, value := range keyPath {
		cpKey[key] = strings.Join(reverse(strings.Split(catPath(key, value), "|")), "|")
		//cpKey[key] = categoryPath(key, value)
	}

	for key, value := range cpKey {
		//fmt.Println(key,":",value)
		if strings.Contains(value, "|") {
			tmpArray := strings.Split(value, "|")
			for i := 0; i < len(tmpArray); i++ {
				if _, ok := catname[tmpArray[i]]; ok {
					tmpArray[i] = catname[tmpArray[i]]
				}
			}
			cpKey[key] = strings.Join(tmpArray, "|")
		}
	}
	return cpKey
}
