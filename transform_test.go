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
