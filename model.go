package bingo_dao

type BaseType byte

const (
	BT_TEXT  BaseType = 1 //文本
	BT_DICT  BaseType = 2 //字典
	BT_MONEY BaseType = 3 //金额

)

//域
type Domain struct {
	Code   string //域code
	Label  string //域名称
	Parent string //父域
}

type Validate func(value string, name string, option string) bool

//数据类型
type DataType struct {
	Base       BaseType //基本类型
	DomainCode string   //所属的域
	Catalog    string   //分类
	Name       string   //类型名称
	Label      string   //类型描述
	Length     int      //长度限制
	Option     string   //校验参数
	Validator  Validate //校验器
}

func (this *DataType) validate(v string) bool {
	//长度校验
	if this.Length > 0 {
		if len(v) > this.Length {
			return false
		}
	}
	//校验器校验
	if this.Validator != nil {
		return this.Validator(v, this.Name, this.Option)
	}
	return true
}

//数据元素
type DataElement struct {
	Type  *DataType //数据类型
	Name  string    //名称
	Field string    //字段名称
	Short string    //短描述
	Head  string    //表头描述
	Media string    //中描述
	Long  string    //长描述
}

//数据结构
type DataStruct struct {
	Name     string
	Label    string //描述
	Type     string //数据结构类型
	Elements []*DataElement
}

//字典
type DictItem struct {
	Code  string
	Label string
	Media string //中描述
	Long  string //长描述
}
type DictCatalog struct {
	Code   string //名称
	Values []DictItem
}
