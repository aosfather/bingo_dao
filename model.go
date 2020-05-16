package bingo_dao

import (
	"fmt"
	"log"
)

type BaseType byte

func (this *BaseType) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var text string
	unmarshal(&text)
	if text == "TEXT" {
		*this = BT_TEXT
	} else if text == "DICT" {
		*this = BT_DICT
	} else if text == "MONEY" {
		*this = BT_MONEY
	} else {
		*this = 0
		return fmt.Errorf("value is wrong! [ %s ]", text)
	}
	return nil
}

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

func (this *Validate) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var text string
	unmarshal(&text)
	v, ok := validates[text]
	if !ok {
		return fmt.Errorf("not found validate[%s]", text)
	}

	*this = v
	return nil
}

/*
Output Length：显示输出的长度；
Convers. Routine：定义数据转换程序名；
*/
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

func (this *DataType) GetDataType() *DataType {
	return this
}

type DT func() *DataType

func (this *DT) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var text string
	unmarshal(&text)
	v := types.GetType(text)
	if v != nil {
		*this = v.GetDataType
	}

	return nil
}

//数据元素
type DataElement struct {
	Name  string //名称
	Type  DT     //数据类型
	Short string //短描述
	Head  string //表头描述
	Media string //中描述
	Long  string //长描述
}

func (this *DataElement) GetDataElement() *DataElement {
	return this
}

//数据结构
type DataStruct struct {
	Name   string
	Label  string //描述
	Type   string //数据结构类型
	Fields []Field
}

type Element func() *DataElement

func (this *Element) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var text string
	unmarshal(&text)
	log.Println(text)
	v := types.GetElement(text)
	if v != nil {

		*this = v.GetDataElement
	}

	return nil
}

type Field struct {
	Name    string
	Element Element
}

//字典
type DictItem struct {
	Code  string
	Label string
	Media string //中描述
	Long  string //长描述
}
type DictCatalog struct {
	Code  string //名称
	Items []DictItem
}

//表类型
type Table struct {
	Name   string
	Label  string //描述
	Type   string //数据结构类型
	Code   string
	Pk     TableIndex
	Indexs []TableIndex //
	Fields []Field
}

//索引
type TableIndex struct {
	Name   string
	Type   string
	Fields []string
}
