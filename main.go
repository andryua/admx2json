package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fatih/structs"
	"regexp"

	//"github.com/fatih/structs"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"unicode"
)

//-------ADMX structs------

type PolicyDefinitions struct {
	XMLName          xml.Name         `xml:"policyDefinitions"`
	Text             string           `xml:",chardata"`
	Xsd              string           `xml:"xsd,attr"`
	Xsi              string           `xml:"xsi,attr"`
	Revision         string           `xml:"revision,attr"`
	SchemaVersion    string           `xml:"schemaVersion,attr"`
	Xmlns            string           `xml:"xmlns,attr"`
	PolicyNamespaces PolicyNamespaces `xml:"policyNamespaces"`
	Resources        Resources        `xml:"resources"`
	SupportedOn      struct {
		Text     string `xml:",chardata"`
		Products struct {
			Text    string `xml:",chardata"`
			Product []struct {
				Text         string `xml:",chardata"`
				Name         string `xml:"name,attr"`
				DisplayName  string `xml:"displayName,attr"`
				MajorVersion []struct {
					Text         string `xml:",chardata"`
					Name         string `xml:"name,attr"`
					DisplayName  string `xml:"displayName,attr"`
					VersionIndex string `xml:"versionIndex,attr"`
				} `xml:"majorVersion"`
			} `xml:"product"`
		} `xml:"products"`
		Definitions struct {
			Text       string `xml:",chardata"`
			Definition []struct {
				Text        string `xml:",chardata"`
				Name        string `xml:"name,attr"`
				DisplayName string `xml:"displayName,attr"`
				And         struct {
					Text      string `xml:",chardata"`
					Reference []struct {
						Text string `xml:",chardata"`
						Ref  string `xml:"ref,attr"`
					} `xml:"reference"`
				} `xml:"and"`
				Or struct {
					Text      string `xml:",chardata"`
					Reference []struct {
						Text string `xml:",chardata"`
						Ref  string `xml:"ref,attr"`
					} `xml:"reference"`
					Range struct {
						Text            string `xml:",chardata"`
						Ref             string `xml:"ref,attr"`
						MinVersionIndex string `xml:"minVersionIndex,attr"`
						MaxVersionIndex string `xml:"maxVersionIndex,attr"`
					} `xml:"range"`
				} `xml:"or"`
			} `xml:"definition"`
		} `xml:"definitions"`
	} `xml:"supportedOn"`
	Categories Categories `xml:"categories"`
	Policies   Policies   `xml:"policies"`
}

type Resources struct {
	Text                string            `xml:",chardata"`
	MinRequiredRevision string            `xml:"minRequiredRevision,attr"`
	StringTable         StringTable       `xml:"stringTable"`
	PresentationTable   PresentationTable `xml:"presentationTable"`
}

type PolicyNamespaces struct {
	Text   string `xml:",chardata"`
	Target Target `xml:"target"`
	Using  Using  `xml:"using"`
}

type Target struct {
	Text      string `xml:",chardata"`
	Prefix    string `xml:"prefix,attr"`
	Namespace string `xml:"namespace,attr"`
}

type Using struct {
	Text      string `xml:",chardata"`
	Prefix    string `xml:"prefix,attr"`
	Namespace string `xml:"namespace,attr"`
}

type Categories struct {
	Text     string     `xml:",chardata"`
	Category []Category `xml:"category"`
}

type Category struct {
	Text           string         `xml:",chardata"`
	Name           string         `xml:"name,attr"`
	DisplayName    string         `xml:"displayName,attr"`
	ExplainText    string         `xml:"explainText,attr"`
	ParentCategory ParentCategory `xml:"parentCategory"`
}

type Policies struct {
	Text   string   `xml:",chardata"`
	Policy []Policy `xml:"policy"`
}

type Policy struct {
	Text           string         `xml:",chardata"`
	Name           string         `xml:"name,attr"`
	Class          string         `xml:"class,attr"`
	DisplayName    string         `xml:"displayName,attr"`
	ExplainText    string         `xml:"explainText,attr"`
	Key            string         `xml:"key,attr"`
	Presentation   string         `xml:"presentation,attr"`
	ValueName      string         `xml:"valueName,attr"`
	ParentCategory ParentCategory `xml:"parentCategory"`
	SupportedOn    SupportedOn    `xml:"supportedOn"`
	EnabledList    EnabledList    `xml:"enabledList"`
	DisabledList   DisabledList   `xml:"disabledList"`
	Elements       Elements       `xml:"elements"`
	EnabledValue   EnabledValue   `xml:"enabledValue"`
	DisabledValue  DisabledValue  `xml:"disabledValue"`
}

type ParentCategory struct {
	Text string `xml:",chardata"`
	Ref  string `xml:"ref,attr"`
}

type SupportedOn struct {
	Text        string      `xml:",chardata"`
	Ref         string      `xml:"ref,attr"`
	Definitions Definitions `xml:"definitions"`
}

type Definitions struct {
	Text       string       `xml:",chardata"`
	Definition []Definition `xml:"definition"`
}

type Definition struct {
	Text        string `xml:",chardata"`
	DisplayName string `xml:"displayName,attr"`
	Name        string `xml:"name,attr"`
}

type EnabledList struct {
	Text string `xml:",chardata"`
	Item []Item `xml:"item"`
}

type DisabledList struct {
	Text string `xml:",chardata"`
	Item []Item `xml:"item"`
}

type Elements struct {
	Chardata    string        `xml:",chardata"`
	Boolean     []Boolean     `xml:"boolean"`
	Textv       []Textv       `xml:"text"`
	Enum        []Enum        `xml:"enum"`
	List        []List        `xml:"list"`
	Decimal     []Decimal     `xml:"decimal"`
	LongDecimal []LongDecimal `xml:"longDecimal"`
	MultiText   []MultiText   `xml:"multiText"`
}

type Boolean struct {
	Text       string     `xml:",chardata"`
	ID         string     `xml:"id,attr"`
	ValueName  string     `xml:"valueName,attr"`
	TrueValue  TrueValue  `xml:"trueValue"`
	FalseValue FalseValue `xml:"falseValue"`
	TrueList   TrueList   `xml:"trueList"`
	FalseList  FalseList  `xml:"falseList"`
	Key        string     `xml:"key,attr"`
}

type FalseValue struct {
	Text    string  `xml:",chardata"`
	Decimal Decimal `xml:"decimal"`
}

type TrueValue struct {
	Text    string  `xml:",chardata"`
	Decimal Decimal `xml:"decimal"`
}

type TrueList struct {
	Text string `xml:",chardata"`
	Item []Item `xml:"item"`
}

type FalseList struct {
	Text string `xml:",chardata"`
	Item []Item `xml:"item"`
}

type Textv struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"id,attr"`
	ValueName  string `xml:"valueName,attr"`
	Required   string `xml:"required,attr"`
	Expandable string `xml:"expandable,attr"`
	Key        string `xml:"key,attr"`
}

type Enum struct {
	Text      string  `xml:",chardata"`
	ID        string  `xml:"id,attr"`
	ValueName string  `xml:"valueName,attr"`
	Required  string  `xml:"required,attr"`
	Item      []Item  `xml:"item"`
	Textv     []Textv `xml:"text"`
}

type ItemVL struct {
	Text      string `xml:",chardata"`
	Key       string `xml:"key,attr"`
	ValueName string `xml:"valueName,attr"`
	Value     Value  `xml:"value"`
}

type ValueList struct {
	Text   string   `xml:",chardata"`
	Itemvl []ItemVL `xml:"item"`
}

type Item struct {
	Text        string    `xml:",chardata"`
	DisplayName string    `xml:"displayName,attr"`
	Value       Value     `xml:"value"`
	Key         string    `xml:"key,attr"`
	ValueName   string    `xml:"valueName,attr"`
	ValueList   ValueList `xml:"valueList"`
	Required    string    `xml:"required"`
}

type Value struct {
	Text    string  `xml:",chardata"`
	Decimal Decimal `xml:"decimal"`
	StringV StringV `xml:"string"`
	Delete  string  `xml:"delete"`
}

type List struct {
	Text            string `xml:",chardata"`
	ID              string `xml:"id,attr"`
	Key             string `xml:"key,attr"`
	ExplicitValue   string `xml:"explicitValue,attr"`
	ValuePrefix     string `xml:"valuePrefix,attr"`
	Additive        string `xml:"additive,attr"`
	ClientExtension string `xml:"clientExtension,attr"`
}

type MultiText struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"id,attr"`
	ValueName  string `xml:"valueName,attr"`
	MaxStrings string `xml:"maxStrings,attr"`
	MaxLength  string `xml:"maxLength,attr"`
	Required   string `xml:"required,attr"`
}

type EnabledValue struct {
	Text        string      `xml:",chardata"`
	StringV     string      `xml:"string"`
	LongDecimal LongDecimal `xml:"longDecimal"`
	Decimal     Decimal     `xml:"decimal"`
}

type DisabledValue struct {
	Text        string      `xml:",chardata"`
	StringV     string      `xml:"string"`
	LongDecimal LongDecimal `xml:"longDecimal"`
	Decimal     Decimal     `xml:"decimal"`
}

type LongDecimal struct {
	Text        string `xml:",chardata"`
	Value       string `xml:"value,attr"`
	ID          string `xml:"id,attr"`
	ValueName   string `xml:"valueName,attr"`
	MaxValue    string `xml:"maxValue,attr"`
	Required    string `xml:"required,attr"`
	MinValue    string `xml:"minValue,attr"`
	StoreAsText string `xml:"storeAsText,attr"`
}

type Decimal struct {
	Text            string `xml:",chardata"`
	Value           string `xml:"value,attr"`
	ID              string `xml:"id,attr"`
	ValueName       string `xml:"valueName,attr"`
	MaxValue        string `xml:"maxValue,attr"`
	Required        string `xml:"required,attr"`
	MinValue        string `xml:"minValue,attr"`
	StoreAsText     string `xml:"storeAsText,attr"`
	Key             string `xml:"key,attr"`
	ClientExtension string `xml:"clientExtension,attr"`
}

//-------ADML structs--------
type PolicyDefinitionResources struct {
	XMLName       xml.Name  `xml:"policyDefinitionResources"`
	Text          string    `xml:",chardata"`
	Xsd           string    `xml:"xsd,attr"`
	Xsi           string    `xml:"xsi,attr"`
	Revision      string    `xml:"revision,attr"`
	SchemaVersion string    `xml:"schemaVersion,attr"`
	Xmlns         string    `xml:"xmlns,attr"`
	DisplayName   string    `xml:"displayName"`
	Description   string    `xml:"description"`
	Resources     Resources `xml:"resources"`
}
type PresentationTable struct {
	Text         string `xml:",chardata"`
	Presentation []struct {
		Chardata string `xml:",chardata"`
		ID       string `xml:"id,attr"`
		CheckBox []struct {
			Text           string `xml:",chardata"`
			RefId          string `xml:"refId,attr"`
			DefaultChecked string `xml:"defaultChecked,attr"`
		} `xml:"checkBox"`
		ComboBox []struct {
			Text       string   `xml:",chardata"`
			RefId      string   `xml:"refId,attr"`
			NoSort     string   `xml:"noSort,attr"`
			Label      string   `xml:"label"`
			Suggestion []string `xml:"suggestion"`
		} `xml:"comboBox"`
		DropdownList []struct {
			Text        string `xml:",chardata"`
			RefId       string `xml:"refId,attr"`
			DefaultItem string `xml:"defaultItem,attr"`
			NoSort      string `xml:"noSort,attr"`
		} `xml:"dropdownList"`
		Text    []string `xml:"text"`
		ListBox struct {
			Text  string `xml:",chardata"`
			RefId string `xml:"refId,attr"`
		} `xml:"listBox"`
		DecimalTextBox struct {
			Text         string `xml:",chardata"`
			RefId        string `xml:"refId,attr"`
			DefaultValue string `xml:"defaultValue,attr"`
			SpinStep     string `xml:"spinStep,attr"`
		} `xml:"decimalTextBox"`
		LongDecimalTextBox struct {
			Text         string `xml:",chardata"`
			RefId        string `xml:"refId,attr"`
			DefaultValue string `xml:"defaultValue,attr"`
			SpinStep     string `xml:"spinStep,attr"`
		} `xml:"longDecimalTextBox"`
		TextBox struct {
			Text         string `xml:",chardata"`
			RefId        string `xml:"refId,attr"`
			Label        string `xml:"label"`
			DefaultValue string `xml:"defaultValue"`
		} `xml:"textBox"`
		MultiTextBox struct {
			Text  string `xml:",chardata"`
			RefId string `xml:"refId,attr"`
		} `xml:"multiTextBox"`
	} `xml:"presentation"`
}

type StringTable struct {
	Text    string    `xml:",chardata"`
	StringV []StringV `xml:"string"`
}

type StringV []struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}

//--------Structs for result JSON---------------
type AllPolicies struct {
	Name        string   `json:"name"`
	Class       string   `json:"class"`
	DisplayName string   `json:"displayName"`
	ExplainText string   `json:"explainText"`
	Category    string   `json:"category"`
	SupportedOn string   `json:"supportedOn"`
	Values      []Values `json:"values"`
}

type Values struct {
	Type          string `json:"type,omitempty"`
	ValueName     string `json:"valueName,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	Key           string `json:"key,omitempty"`
	Required      string `json:"required,omitempty"`
	MaxValue      string `json:"maxValue,omitempty"`
	MinValue      string `json:"minValue,omitempty"`
	Value         string `json:"value,omitempty"`
	DisabledValue string `json:"disabledValue,omitempty""`
	EnabledValue  string `json:"enabledValue,omitempty"`
	TrueValue     string `json:"trueValue,omitempty"`
	FalseValue    string `json:"falseValue,omitempty"`
	ValuePrefix   string `json:"valuePrefix,omitempty"`
}

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func categoryMap(v []Category) map[string]string {
	fullpath := make(map[string]string)
	for _, c := range v {
		if c.ParentCategory.Ref != "" {
			fullpath[c.Name] = c.ParentCategory.Ref
		} else {
			fullpath[c.Name] = c.Name
		}
	}
	return fullpath
}

var keyPath map[string]string

//var fullPath map[string]string
/*
func categoryPath (s string) string {
	if s == "" {
		return s
	} else {
		return  "/" + keyPath[s] + "/" + categoryPath(keyPath[s])
	}
}
*/
func categoryPath(key, val string) string {
	path := key
	if val == keyPath[val] || key == val {
		return path
	}

	path += "/" + val
	return categoryPath(path, keyPath[val])

}

func reverse(item []string) []string {
	newItem := make([]string, 0, len(item))
	for i := len(item) - 1; i >= 0; i-- {
		newItem = append(newItem, item[i])
	}
	return newItem
}

func main() {
	var n PolicyDefinitions
	var m PolicyDefinitionResources
	var it Values
	lang := make(map[string]string)
	res := make([]AllPolicies, len(n.Policies.Policy))
	var admx []AllPolicies
	root := "/home/DN301081KAI/go/src/admx/gpo"
	enUs := "/home/DN301081KAI/go/src/admx/gpo/en-US"
	fnamesX, err := ioutil.ReadDir(root)
	fnamesL, err := ioutil.ReadDir(enUs)
	fnames := make(map[string]string)
	for _, fnameX := range fnamesX {
		fx := strings.Split(fnameX.Name(), ".")[0]
		for _, fnameL := range fnamesL {
			fl := strings.Split(fnameL.Name(), ".")[0]
			if strings.ToLower(fx) == strings.ToLower(fl) {
				fnames[fx+".admx"] = fl + ".adml"
			}
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	//for _, fname := range fnames {
	//	if strings.Contains(fname.Name(), "admx") != true {
	//		continue
	//	}
	//	xmlFileX, err := os.Open("gpo/"+fname.Name())
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fl := strings.Split(fname.Name(), ".")[0]
	//	xmlFileL, err := os.Open("gpo/en-US/" + fl + ".adml")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	for key, value := range fnames {
		xmlFileX, err := os.Open("gpo/" + key)
		if err != nil {
			fmt.Println(err)
		}
		xmlFileL, err := os.Open("gpo/en-US/" + value)
		if err != nil {
			fmt.Println(err)
		}
		defer xmlFileL.Close()
		defer xmlFileX.Close()

		byteValueX, _ := ioutil.ReadAll(xmlFileX)
		byteValueL, _ := ioutil.ReadAll(xmlFileL)

		printOnly := func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}
		byteValueX = []byte(strings.Map(printOnly, string(byteValueX)))
		byteValueL = []byte(strings.Map(printOnly, string(byteValueL)))

		err = xml.Unmarshal(byteValueX, &n)
		if err != nil {
			fmt.Println("error: ", key, ": ", err)
		}

		err = xml.Unmarshal(byteValueL, &m)
		if err != nil {
			fmt.Println("error: ", value, ": ", err)
		}

		var rgx, _ = regexp.Compile(`..string.(\S+).`)
		var rgxS = regexp.MustCompile(`(SUPPORT\S+)`)
		if xmlFileL != nil {
			for _, str := range m.Resources.StringTable.StringV {
				for _, data := range str {
					lang[data.ID] = data.Text
				}
			}
		}
		clear(&res)
		catname := make(map[string]string)
		for _, category := range n.Categories.Category {
			catname[category.Name] = lang[rgx.FindStringSubmatch(category.DisplayName)[1]]
			if catname[category.Name] != "" {
				//fmt.Println("Category name: ", catname,"%%%%% Parent category: " ,category.ParentCategory.Ref)
			} else {
				//fmt.Println("Category name: ", category.DisplayName, "%%%%% Parent category: " ,category.ParentCategory.Ref)
			}
		}
		keyPath = categoryMap(n.Categories.Category)
		//fullPath = keyPath
		tmp := ""
		tmpArray := []string{}
		for key, value := range keyPath {
			tmp = categoryPath(key, value)
			tmpArray = strings.Split(tmp, "/")
			for i := 0; i < len(tmpArray); i++ {
				if strings.Contains(tmpArray[i], ":") {
					tmpArray[i] = strings.Split(tmpArray[i], ":")[1]
				}
				if catname[tmpArray[i]] != "" {
					tmpArray[i] = catname[tmpArray[i]]
				}
			}
			keyPath[key] = strings.Join(reverse(tmpArray), "/")
		}
		//rgx.FindStringSubmatch(item.DisplayName)[1]
		for _, policy := range n.Policies.Policy {
			var r AllPolicies
			/*
				it.Type
				it.ValueName
				it.DisplayName
				it.Key
				it.Required
				it.MaxValue
				it.MinValue
				it.Value
				it.DisabledValue
				it.EnabledValue
				it.TrueValue
				it.FalseValue
				it.ValuePrefix
			*/
			r.Name = policy.Name
			if len(rgxS.FindAllString(policy.SupportedOn.Ref, -1)) > 0 {
				r.SupportedOn = lang[rgxS.FindAllString(policy.SupportedOn.Ref, -1)[0]]
			} else {
				//fmt.Println(fname.Name(), " : ",rgxS.FindAllString(policy.SupportedOn.Ref,-1), " - ", policy.SupportedOn.Ref)
				if policy.SupportedOn.Ref == "" {
					r.SupportedOn = "All Windows versions"
				} else {
					r.SupportedOn = policy.SupportedOn.Ref
				}
			}
			r.Class = policy.Class
			/*
				if lang[policy.ParentCategory.Ref] != "" {
					r.Category = lang[policy.ParentCategory.Ref]
				} else {
					r.Category = policy.ParentCategory.Ref
				}
			*/
			tmp := policy.ParentCategory.Ref
			if strings.Contains(policy.ParentCategory.Ref, ":") {
				tmp = strings.Split(policy.ParentCategory.Ref, ":")[1]
			}
			if keyPath[tmp] != "" {
				r.Category = keyPath[tmp]
			} else {
				r.Category = tmp
			}
			if lang[rgx.FindStringSubmatch(policy.DisplayName)[1]] != "" {
				r.DisplayName = lang[rgx.FindStringSubmatch(policy.DisplayName)[1]]
			} else {
				r.DisplayName = policy.DisplayName
			}
			if lang[rgx.FindStringSubmatch(policy.ExplainText)[1]] != "" {
				r.ExplainText = lang[rgx.FindStringSubmatch(policy.ExplainText)[1]]
			} else {
				r.ExplainText = policy.ExplainText
			}

			//----values-----
			if policy.ValueName != "" && policy.EnabledValue == (EnabledValue{}) {
				it.Key = policy.Key
				it.ValueName = policy.ValueName
				it.Type = "REG_DWORD"
				it.EnabledValue = policy.EnabledValue.Decimal.Value
				it.DisabledValue = policy.DisabledValue.Decimal.Value
				if it != (Values{}) {
					r.Values = append(r.Values, it)
				}
				it = Values{}
			}
			if policy.Elements.Chardata != "" {
				if policy.Elements.Enum != nil {
					for _, item := range policy.Elements.Enum {
						for _, itm := range item.Item {
							it.Key = policy.Key
							if itm.ValueName != "" {
								it.ValueName = itm.ValueName
							} else if item.ValueName != "" {
								it.ValueName = item.ValueName
							} else {
								it.ValueName = policy.ValueName
							}

							it.DisplayName = lang[rgx.FindStringSubmatch(itm.DisplayName)[1]]
							it.Required = itm.Required
							if itm.Value.StringV != nil {
								it.Type = "REG_SZ"
							}
							if itm.Value.Decimal != (Decimal{}) {
								it.Type = "REG_DWORD"
								it.Value = itm.Value.Decimal.Value
							}
							if !structs.IsZero(itm.ValueList) {
								for _, i := range itm.ValueList.Itemvl {
									it.ValueName = i.ValueName
									it.Key = i.Key
									if i.Value.StringV != nil {
										it.Type = "REG_SZ"
									}
									if i.Value.Decimal.Value != "" {
										it.Type = "REG_DWORD"
										it.Value = i.Value.Decimal.Value
										it.Key = i.Value.Decimal.Key
										it.ValueName = i.Value.Decimal.ValueName
										it.Required = i.Value.Decimal.Required
									}
									if i.Value.Delete != "" {
										it.Type = ""
										it.Value = "DELETE"
									}
									it.Key = i.Key
									it.ValueName = i.ValueName
								}
							}
							if (Values{}) != it {
								r.Values = append(r.Values, it)
							}
						}
						for _, itm := range item.Textv {
							it.Key = policy.Key
							it.ValueName = itm.ValueName
							it.Required = itm.Required
							it.Type = "REG_SZ"
							if (Values{}) != it {
								r.Values = append(r.Values, it)
							}
						}
					}
					it = Values{}
				}
				if policy.Elements.Textv != nil {
					for _, item := range policy.Elements.Textv {
						it.Key = item.Key
						it.ValueName = item.ValueName
						it.Required = item.Required
						it.Type = "REG_SZ"
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
					}
					it = Values{}
				}
				if policy.Elements.Boolean != nil {
					for _, item := range policy.Elements.Boolean {
						if item.Key != "" {
							it.Key = item.Key
						} else {
							it.Key = policy.Key
						}
						it.ValueName = item.ValueName
						it.Type = "REG_DWORD"
						it.TrueValue = item.TrueValue.Decimal.Value
						it.FalseValue = item.FalseValue.Decimal.Value
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
					}
					it = Values{}
				}
				if policy.Elements.MultiText != nil {
					for _, item := range policy.Elements.MultiText {
						it.Key = policy.Key
						it.ValueName = item.ValueName
						it.Required = item.Required
						it.Type = "REG_MULTISZ"
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
					}
					it = Values{}
				}
				if policy.Elements.Decimal != nil {
					for _, item := range policy.Elements.Decimal {
						it.Key = policy.Key
						it.ValueName = item.ValueName
						it.Required = item.Required
						it.Type = "REG_DWORD"
						it.Value = item.Value
						it.MaxValue = item.MaxValue
						it.MinValue = item.MinValue
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
					}
					it = Values{}
				}
				if policy.Elements.LongDecimal != nil {
					for _, item := range policy.Elements.LongDecimal {
						it.Key = policy.Key
						it.ValueName = item.ValueName
						it.Required = item.Required
						it.Type = "REG_QWORD"
						it.Value = item.Value
						it.MaxValue = item.MaxValue
						it.MinValue = item.MinValue
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
					}
					it = Values{}
				}
				if policy.Elements.List != nil {
					for _, item := range policy.Elements.List {
						it.Key = item.Key
						it.Type = "REG_SZ"
						it.ValueName = "manual " + item.ValuePrefix + ""
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
						it = Values{}
					}
				}
			}
			if policy.EnabledValue != (EnabledValue{}) {
				it.EnabledValue = policy.EnabledValue.Decimal.Value
				it.DisabledValue = policy.DisabledValue.Decimal.Value
				it.Key = policy.Key
				it.ValueName = policy.ValueName
				it.Type = "REG_DWORD"
				if (Values{}) != it {
					r.Values = append(r.Values, it)
				}
				it = Values{}
			}

			res = append(res, r)
			clear(&r)
		}
		for _, pol := range res {
			admx = append(admx, pol)
		}

	}
	/*
		for key, value := range keyPath {
			fmt.Println(key + " = ", value)
		}
	*/
	//fmt.Println(categoryPath(n.Categories.Category))
	jsonres, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fil := "gpo.json"
	ioutil.WriteFile(fil, jsonres, 0777)
	//fmt.Println(string(jsonres))
}
