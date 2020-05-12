package bingo_dao

import "regexp"

//正则表达式
type RegexValidator struct {
	patterns map[string]*regexp.Regexp
}

func (this *RegexValidator) Init() {
	this.patterns = make(map[string]*regexp.Regexp)
}
func (this *RegexValidator) Validate(value string, name string, option string) bool {
	pattern := this.patterns[name]
	if pattern == nil && option != "" {
		pattern, err := regexp.Compile(option)
		if err == nil {
			this.patterns[name] = pattern
		}
	}
	if pattern != nil {
		return pattern.Match([]byte(value))
	}

	return false

}

//字典校验器
type DictionaryValidator struct {
	dictionary map[string]DictCatalog
}

func (this *DictionaryValidator) Init() {
	this.dictionary = make(map[string]DictCatalog)
}

func (this *DictionaryValidator) Validate(value string, name string, dict string) bool {
	catalog := this.dictionary[dict]
	//轮询code看是否属于取值范围内的值
	if catalog.Code != "" {
		for _, v := range catalog.Values {
			if v.Code == value {
				return true
			}
		}
	}

	return false
}

func (this *DictionaryValidator) AddCatalog(catalog DictCatalog) {
	this.dictionary[catalog.Code] = catalog
}
func (this *DictionaryValidator) AddCatalogByItem(c string, item ...DictItem) {
	this.dictionary[c] = DictCatalog{c, item}
}
