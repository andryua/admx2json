package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

var dataCat = map[string]string{}

func categoryMap(v []Category) map[string]string {
	for _, c := range v {
		if _, ok := dataCat[c.Name]; ok {
			continue
		}
		if c.ParentCategory.Ref != "" {
			dataCat[c.Name] = c.ParentCategory.Ref
		} else {
			dataCat[c.Name] = c.Name
		}
	}
	return dataCat
}

func ParseFiles() ([]PolicyDefinitions, map[string]string, map[string]string, map[string]string) {
	var data []PolicyDefinitions
	var n PolicyDefinitions
	var m PolicyDefinitionResources
	lang := make(map[string]string)
	root := "/home/DN301081KAI/go/src/admx/gpo"
	enUs := "/home/DN301081KAI/go/src/admx/gpo/en-US"
	fnamesX, err := ioutil.ReadDir(root)
	fnamesL, err := ioutil.ReadDir(enUs)
	fnames := make(map[string]string)
	var catalogname = map[string]string{}
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
	var rgx, _ = regexp.Compile(`..string.(\S+).`)

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

		if xmlFileL != nil {
			for _, str := range m.Resources.StringTable.StringV {
				for _, data := range str {
					if _, ok := lang[data.ID]; ok {
						continue
					}
					lang[data.ID] = data.Text
				}
			}
		}
		if xmlFileX != nil {
			for _, category := range n.Categories.Category {
				if _, ok := catalogname[category.Name]; ok {
					continue
				}
				catalogname[category.Name] = lang[rgx.FindStringSubmatch(category.DisplayName)[1]]
			}
		}
		dataCat = categoryMap(n.Categories.Category)
		data = append(data, n)
		clear(&n)
	}
	return data, lang, dataCat, catalogname
}
