package bingo_dao

type Types struct {
	types    map[string]*DataType
	elements map[string]*DataElement
	dicts    *DictionaryValidator
	regex    *RegexValidator
}

func (this *Types) init() {
	this.types = make(map[string]*DataType)
	this.elements = make(map[string]*DataElement)
	this.dicts = new(DictionaryValidator)
	this.dicts.Init()
	SetValidate("dict", this.dicts.Validate)

	this.regex = new(RegexValidator)
	this.regex.Init()
	SetValidate("regex", this.regex.Validate)
}

func (this *Types) GetType(name string) *DataType {
	return this.types[name]
}
func (this *Types) GetElement(name string) *DataElement {
	return this.elements[name]
}

func (this *Types) AddType(t *DataType) {
	if _, ok := this.types[t.Name]; !ok {
		this.types[t.Name] = t
	}
}

var validates map[string]Validate = make(map[string]Validate)
var types Types

func SetValidate(key string, v Validate) {
	if key != "" && v != nil {
		validates[key] = v
	}

}
func GetTypes() *Types {
	return &types
}
func init() {
	types = Types{}
	types.init()
}
