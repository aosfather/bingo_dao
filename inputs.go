package bingo_dao

/**
  输入相关的定义
*/
import "fmt"

//参数策略类型
type PolicyType byte

const (
	Must   PolicyType = 1
	Option PolicyType = 2
)

func (this *PolicyType) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var text string
	unmarshal(&text)
	if text == "Must" {
		*this = Must
	} else if text == "Option" {
		*this = Option
	} else {
		*this = 0
		return fmt.Errorf("value is wrong! [ %s ]", text)
	}
	return nil
}

func (this PolicyType) MarshalYAML() (interface{}, error) {
	if this == Must {
		return "Must", nil
	} else if this == Option {
		return "Option", nil
	}
	return nil, fmt.Errorf("not surport %v", this)
}

type Paramter struct {
	Name   string      `yaml:"paramter"` //参数名称
	Type   DataElement `yaml:"type"`     //参数类型
	Policy PolicyType  `yaml:"policy"`
}
