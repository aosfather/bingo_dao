package bingo_dao

import "testing"

func TestConcatFileds_BuildTransform(t *testing.T) {
	value := make(map[string]interface{})
	value["id"] = "1234"
	value["j1"] = "abc"
	value["j2"] = "efg"
	c := ConcatFileds{NewField: "New", JoinString: "_", Fields: []string{"j1", "j2"}}
	t.Log(c.BuildTransform()(value))
}

func TestFieldMapping_BuildTransform(t *testing.T) {
	value := make(map[string]interface{})
	value["id"] = "1234"
	value["j1"] = "abc"
	value["status"] = "1"
	c := FieldMapping{NewField: "", Field: "status", DefaultValue: "null", Mapping: map[string]string{"1": "TRUE"}}
	t.Log(c.BuildTransform()(value))
	c2 := FieldMapping{NewField: "Nstatus", Field: "status", DefaultValue: "null", Mapping: map[string]string{"1": "TRUE"}}
	t.Log(c2.BuildTransform()(value))
}

func TestCutField_BuildTransform(t *testing.T) {
	value := make(map[string]interface{})
	value["id"] = "1234"
	value["j1"] = "abc"
	value["status"] = "1"
	c := CutField{NewField: "New", Field: "j1", Start: 0, End: 1}
	t.Log(c.BuildTransform()(value))
}

func TestReplaceField_BuildTransform(t *testing.T) {
	value := make(map[string]interface{})
	value["id"] = "1234"
	value["j1"] = "abcdefghijk"
	c := ReplaceField{NewField: "New", Field: "j1", Match: "abc", Replace: "替换", EnableRegexp: false}
	t.Log(c.BuildTransform()(value))
	c1 := ReplaceField{NewField: "New1", Field: "j1", Match: "babc", Replace: "替换", EnableRegexp: false}
	t.Log(c1.BuildTransform()(value))
	//使用正则表达式
	c2 := ReplaceField{NewField: "New2", Field: "j1", Match: "bc?", Replace: "替换", EnableRegexp: true}
	t.Log(c2.BuildTransform()(value))
}

func TestTrimField_BuildTransform(t *testing.T) {
	value := make(map[string]interface{})
	value["id"] = " 1234 "
	c := TrimField{NewField: "New", Field: "id", Type: TT_Left}
	t.Log(c.BuildTransform()(value))
	c1 := TrimField{NewField: "New1", Field: "id", Type: TT_Right}
	t.Log(c1.BuildTransform()(value))
	c2 := TrimField{NewField: "New2", Field: "id", Type: TT_Both}
	t.Log(c2.BuildTransform()(value))
}

func TestCaptionField_BuildTransform(t *testing.T) {
	value := make(map[string]interface{})
	value["id"] = "abCDefGH"
	c := CaptionField{NewField: "New", Field: "id", Type: CT_Title}
	t.Log(c.BuildTransform()(value))
	c1 := CaptionField{NewField: "New1", Field: "id", Type: CT_Upper}
	t.Log(c1.BuildTransform()(value))
	c2 := CaptionField{NewField: "New2", Field: "id", Type: CT_Lowcase}
	t.Log(c2.BuildTransform()(value))

	c3 := CaptionField{NewField: "New3", Field: "id", Type: CT_First}
	t.Log(c3.BuildTransform()(value))
}
