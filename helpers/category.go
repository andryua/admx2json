package helpers

import (
	"strings"
)

var keyPath map[string]string

func categoryPath(key, val string) string {
	path := key
	if strings.EqualFold(val, keyPath[val]) || strings.EqualFold(key, val) || val == "" {
		return path
	}
	path += "|" + val
	return categoryPath(path, keyPath[val])
}

func reverse(item []string) []string {
	newItem := make([]string, len(item))
	for i, j := 0, len(item)-1; i <= j; i, j = i+1, j-1 {
		newItem[i], newItem[j] = item[j], item[i]
	}
	return newItem
}

func CategoriesPath(keyPath, catname map[string]string) map[string]string {
	for key, value := range keyPath {
		keyPath[key] = categoryPath(key, value)
	}

	for key, value := range keyPath {
		tmpArray := []string{}
		if strings.Contains(value, "|") {
			tmpArray = strings.Split(value, "|")
			for i := 0; i < len(tmpArray); i++ {
				if strings.Contains(tmpArray[i], ":") {
					tmpArray[i] = strings.Split(tmpArray[i], ":")[1]
				}
				if catname[tmpArray[i]] != "" {
					tmpArray[i] = catname[tmpArray[i]]
				}
			}
			keyPath[key] = strings.Join(reverse(tmpArray), "|")
		} else {
			if value != "" {
				keyPath[key] = catname[value]
			}
		}
	}
	return keyPath
}
