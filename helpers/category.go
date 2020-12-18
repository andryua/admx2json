package helpers

import (
	"strings"
)

var keyPath map[string]string

func catPath(k1, v1 string) string {
	path := k1
	if v1 == "" || strings.EqualFold(v1, keyPath[v1]) {
		return path
	}
	path += "|" + v1
	return catPath(path, keyPath[v1])
}

func reverse(item []string) []string {
	newItem := make([]string, len(item))
	for i, j := 0, len(item)-1; i <= j; i, j = i+1, j-1 {
		newItem[i], newItem[j] = item[j], item[i]
	}
	return newItem
}

func CategoriesPath(inputKeyPath, catname map[string]string) map[string]string {
	keyPath = inputKeyPath
	cpKey := make(map[string]string)
	for key, value := range keyPath {
		cpKey[key] = strings.Join(reverse(strings.Split(catPath(key, value), "|")), "|")
	}
	for key, value := range cpKey {
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
