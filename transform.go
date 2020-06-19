package bingo_dao

import (
	"regexp"
	"strings"
	"unicode"
)

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

/**
  3、增加常亮作为新的字段
*/
type AddConstant struct {
	singleComponse
	NewField string
	Type     string
	Value    string
}

func (this *AddConstant) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		//只有不存在新字段的时候才加入，如果存在该字段则不作处理
		if _, ok := input[this.NewField]; !ok {
			input[this.NewField] = this.Value
		}
		return input
	}
	return t
}

/**
  4、字段截取
*/
type CutField struct {
	singleComponse
	NewField string
	Field    string
	Start    int
	End      int
}

func (this *CutField) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		//只有不存在新字段的时候才加入，如果存在该字段则不作处理
		if v, ok := input[this.Field]; ok {
			target := v.(string)
			target = target[this.Start:this.End]
			input[this.NewField] = target
		}
		return input
	}
	return t
}

/*
  5、字符串替换
  * 输入字段
  * 输出字段
  * 匹配内容
  * 替换内容
  * 启用正则表达式、大小写敏感、整个单词匹配
  * 设为空串
*/
type ReplaceField struct {
	singleComponse
	NewField     string
	Field        string
	Match        string
	EnableRegexp bool
	Replace      string
	regex        *regexp.Regexp
}

func (this *ReplaceField) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		//只有不存在新字段的时候才加入，如果存在该字段则不作处理
		if v, ok := input[this.Field]; ok {
			target := v.(string)
			if this.EnableRegexp {
				if this.regex == nil {
					this.regex, _ = regexp.Compile(this.Match)
				}
				target = this.regex.ReplaceAllString(target, this.Replace)
			} else {
				target = strings.Replace(target, this.Match, this.Replace, -1)
			}
			input[this.NewField] = target
		}
		return input
	}
	return t
}

type TrimType byte

const (
	TT_Left  TrimType = 1 //左边空格
	TT_Right TrimType = 2 //右边空格
	TT_Both  TrimType = 3 //两端空格
)

/**
  6、字符串去空格
*/
type TrimField struct {
	singleComponse
	NewField string
	Field    string
	Type     TrimType
}

func (this *TrimField) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		if v, ok := input[this.Field]; ok {
			target := v.(string)
			switch this.Type {
			case TT_Left:
				target = strings.TrimLeftFunc(target, unicode.IsSpace)
			case TT_Right:
				target = strings.TrimRightFunc(target, unicode.IsSpace)
			case TT_Both:
				target = strings.TrimSpace(target)

			}
			input[this.NewField] = target
		}
		return input
	}
	return t
}

//大小写类型
type UpperLowType byte

const (
	CT_Title   UpperLowType = 3 //标题
	CT_First   UpperLowType = 4 //首字母大写
	CT_Upper   UpperLowType = 1 //转大写
	CT_Lowcase UpperLowType = 2 //转小写
)

type CaptionField struct {
	singleComponse
	NewField string
	Field    string
	Type     UpperLowType
}

func (this *CaptionField) BuildTransform() TransformFunction {
	t := func(input Data) Data {
		if v, ok := input[this.Field]; ok {
			target := v.(string)
			switch this.Type {
			case CT_Upper:
				target = strings.ToUpper(target)
			case CT_Lowcase:
				target = strings.ToLower(target)
			case CT_Title:
				target = strings.ToTitle(target)
			case CT_First:
				strArry := []rune(target)
				if strArry[0] >= 97 && strArry[0] <= 122 {
					strArry[0] -= 32
				}
				target = string(strArry)

			}
			input[this.NewField] = target
		}
		return input
	}
	return t
}
