package bingo_dao

import "strings"

/*
  数据转换组件
*/
//数据数组
type DataArray *[]Data
type Data map[string]interface{}

//单条处理函数
type TransformFunction func(input Data) Data
type ArrayTransform func(input DataArray)

//转换组件
type TransformComponse interface {
	IsArray() bool
	BuildTransform() TransformFunction
	BuildArrayTransform() ArrayTransform
}

type singleComponse struct {
}

func (this *singleComponse) IsArray() bool {
	return false
}

func (this *singleComponse) BuildArrayTransform() ArrayTransform {
	return nil
}

/*
  1、Concat fields
    新字段
    连接符号
    连接的字段
*/
type ConcatFileds struct {
	singleComponse
	NewField   string
	JoinString string
	Fields     []string
}

func (this *ConcatFileds) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		var values []string
		for _, f := range this.Fields {
			values = append(values, input[f].(string))
		}
		v := strings.Join(values, this.JoinString)
		input[this.NewField] = v
		return input
	}
	return t
}

/*
  2、值映射
  * 需要映射的字段
  * 映射生成的字段
  * 没有匹配的默认值
  * 映射规则（key-->value）
*/
type FieldMapping struct {
	singleComponse
	NewField     string
	Field        string
	DefaultValue string
	Mapping      map[string]string
}

func (this *FieldMapping) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		v := input[this.Field]
		//获取映射值
		target, ok := this.Mapping[v.(string)]
		if !ok {
			target = this.DefaultValue
		}
		//设置映射后的结果值
		targetField := this.NewField
		if targetField == "" {
			targetField = this.Field
		}
		input[targetField] = target
		return input
	}
	return t
}
