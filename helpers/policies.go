package helpers

import (
	//"github.com/fatih/structs"
	"reflect"
	"regexp"
	"strings"
)

//--------Structs for result JSON---------------
type AllPolicies struct {
	ID          int      `json:"id"`
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

func PoliciesParse(data []Policy, lang map[string]string, keyPath map[string]string) []AllPolicies {
	var res []AllPolicies
	var rgx, _ = regexp.Compile(`..string.(\S+).`)
	var rgxS = regexp.MustCompile(`(SUPPORT\S+)`)
	//rgx.FindStringSubmatch(item.DisplayName)[1]
	k := -1
	for _, policy := range data {
		var r = AllPolicies{}
		var it = Values{}
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
		if policy.Name == "" {
			continue
		}
		k++
		r.ID = k
		r.Name = policy.Name
		if len(rgxS.FindAllString(policy.SupportedOn.Ref, -1)) == 0 {
			r.SupportedOn = lang[rgxS.FindAllString("SUPPORTED_Windows7ToVistaAndWindows10", -1)[0]]
		} else {
			r.SupportedOn = lang[rgxS.FindAllString(policy.SupportedOn.Ref, -1)[0]]
		}

		r.Class = policy.Class

		tmp := policy.ParentCategory.Ref
		if strings.Contains(policy.ParentCategory.Ref, ":") {
			tmp = strings.Split(policy.ParentCategory.Ref, ":")[1]
		}
		tmp = strings.TrimSpace(tmp)
		//fmt.Println(tmp,":",keyPath[tmp])
		//if keyPath[tmp] != "" {
		r.Category = keyPath[tmp]
		//} else {
		//	r.Category = tmp
		//}
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

		if policy.EnabledList.Item != nil {
			for _, en := range policy.EnabledList.Item {
				if en.Key == "" {
					it.Key = policy.Key
				} else {
					it.Key = en.Key
				}
				it.Value = en.Value.Decimal.Value
				it.ValueName = en.ValueName
				it.Type = "REG_DWORD"
				if (Values{}) != it {
					r.Values = append(r.Values, it)
				}
			}
		}
		if policy.DisabledList.Item != nil {
			for _, di := range policy.DisabledList.Item {
				if di.Key == "" {
					it.Key = policy.Key
				} else {
					it.Key = di.Key
				}

				it.Value = di.Value.Decimal.Value
				it.ValueName = di.ValueName
				it.Type = "REG_DWORD"
				if (Values{}) != it {
					r.Values = append(r.Values, it)
				}
			}
		}

		if policy.Elements.Chardata != "" {
			if policy.Elements.Enum != nil {
				for _, item := range policy.Elements.Enum {
					for _, itm := range item.Item {
						//if itm.Key == "" {
						//	it.Key = itm.Key
						//} else {
						it.Key = policy.Key
						//}
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
						//if !structs.IsZero(itm.ValueList) {
						//if itm.ValueList != {} {
						for _, i := range itm.ValueList.Itemvl {
							it.ValueName = i.ValueName
							if i.Key == "" {
								it.Key = policy.Key
							} else {
								it.Key = i.Key
							}
							if i.Value.StringV != nil {
								it.Type = "REG_SZ"
							}
							if i.Value.Decimal.Value != "" {
								it.Type = "REG_DWORD"
								it.Value = i.Value.Decimal.Value
								//fmt.Println(i.Value.Decimal.Key)
								//if i.Value.Decimal.Key == "" {
								//} else {
								//	it.Key = i.Value.Decimal.Key
								//}
								it.ValueName = i.Value.Decimal.ValueName
								it.Required = i.Value.Decimal.Required
							}
							if i.Value.Delete != "" {
								it.Type = ""
								it.Value = "DELETE"
							}
							//it.Key = i.Key
							it.ValueName = i.ValueName
						}
						//it.Key = policy.Key
						//}
						if (Values{}) != it {
							r.Values = append(r.Values, it)
						}
					}
					for _, itm := range item.Textv {
						if itm.Key == "" {
							it.Key = policy.Key
						} else {
							it.Key = itm.Key
						}
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
					if item.Key == "" {
						it.Key = policy.Key
					} else {
						it.Key = item.Key
					}
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
					if item.Key == "" {
						it.Key = policy.Key
					} else {
						it.Key = item.Key
					}
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
					if item.Key == "" {
						it.Key = policy.Key
					} else {
						it.Key = item.Key
					}
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
	}
	return res
}
