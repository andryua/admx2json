package helpers

import (
	"regexp"
	"strings"
)

var categormap = map[string]string{}

func categoryMap(v []Category) map[string]string {
	for _, c := range v {
		if _, ok := categormap[c.Name]; ok {
			continue
		}
		if c.ParentCategory.Ref != "" {
			categormap[c.Name] = c.ParentCategory.Ref
		} else {
			categormap[c.Name] = ""
		}
	}
	return categormap
}

var keyPath map[string]string

func categoryPath(key, val string) string {
	path := key
	if strings.EqualFold(val, keyPath[val]) || strings.EqualFold(key, val) || len(val) <= 1 {
		return path
	}
	path += "\\" + val
	return categoryPath(path, keyPath[val])
}

func reverse(item []string) []string {
	newItem := make([]string, len(item))
	for i, j := 0, len(item)-1; i <= j; i, j = i+1, j-1 {
		newItem[i], newItem[j] = item[j], item[i]
	}
	return newItem
}

func CategoriesPath(n PolicyDefinitions, lang map[string]string) map[string]string {
	var rgx, _ = regexp.Compile(`..string.(\S+).`)

	catname := make(map[string]string)
	for _, category := range n.Categories.Category {
		catname[category.Name] = lang[rgx.FindStringSubmatch(category.DisplayName)[1]]
	}
	keyPath = categoryMap(n.Categories.Category)

	tmp := ""
	tmpArray := []string{}

	for key, value := range keyPath {
		tmp = categoryPath(key, value)
		if strings.Contains(tmp, "\\") {
			tmpArray = strings.Split(tmp, "\\")
			for i := 0; i < len(tmpArray); i++ {
				if strings.Contains(tmpArray[i], ":") {
					tmpArray[i] = strings.Split(tmpArray[i], ":")[1]
				}
				if catname[tmpArray[i]] != "" {
					tmpArray[i] = catname[tmpArray[i]]
				}
			}
			keyPath[key] = strings.Join(reverse(tmpArray), "\\")
		} else {
			keyPath[key] = tmp
		}
	}
	return keyPath
}
