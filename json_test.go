package jsonotg

import (
	"testing"
)

type someStruct struct {
	SomeField string `json:"someField"`
}

const jsonStr = `{
	"level0": null,
	"level1": "This is the first level",
	"level2": {
		"key1": 123,
		"key2": "Another string value",
		"key3": false,
		"nestedArray": [
			{
				"subkey1": true,
				"subkey2": 4.56,
				"evenMoreNested": {
				"finalKey": "This is the deepest level"
				}
			},
			"justAString",
			100,
			{
				"someField": "An object within an array"
			}
	  	]
	}
}`

func TestParseJSON(t *testing.T) {
	j, err := CreateJSON([]byte(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	tcs := []struct {
		Name   string
		F      func() (interface{}, bool)
		Expect any
	}{
		{
			Name:   "level0",
			F:      func() (interface{}, bool) { return j.GetField("level0").IsNull(), true },
			Expect: true,
		},
		{
			Name:   "level1",
			F:      func() (interface{}, bool) { return j.GetField("level1").AsString() },
			Expect: "This is the first level",
		},
		{
			Name: "level2.key1",
			F: func() (interface{}, bool) {
				return j.GetField("level2").GetField("key1").AsInt64()
			},
			Expect: int64(123),
		},
		{
			Name: "level2.key3",
			F: func() (interface{}, bool) {
				return j.GetField("level2").GetField("key3").AsBool()
			},
			Expect: false,
		},
		{
			Name: "level2.nestedArray",
			F: func() (interface{}, bool) {
				arr, ok := j.GetField("level2").GetField("nestedArray").AsArray()
				return len(arr), ok
			},
			Expect: 4,
		},
		{
			Name: "resp.level2.nestedArray[1]",
			F: func() (interface{}, bool) {
				arr, _ := j.GetField("level2").GetField("nestedArray").AsArray()
				return arr[1].AsString()
			},
			Expect: "justAString",
		},
		{
			Name: "resp.level2.nestedArray[2]",
			F: func() (interface{}, bool) {
				arr, _ := j.GetField("level2").GetField("nestedArray").AsArray()
				return arr[2].AsInt64()
			},
			Expect: int64(100),
		},
		{
			Name: "resp.level2.nestedArray[3]",
			F: func() (interface{}, bool) {
				arr, _ := j.GetField("level2").GetField("nestedArray").AsArray()
				var something someStruct
				if err = arr[3].Unmarshal(&something); err != nil {
					t.Fatal(err)
				}

				return something, true
			},
			Expect: someStruct{SomeField: "An object within an array"},
		},
	}

	for _, tc := range tcs {
		out, ok := tc.F()
		if !ok {
			t.Errorf("%s: expected ok, got !ok", tc.Name)
		}

		if out != tc.Expect {
			t.Errorf("%s: expected %v, got %v", tc.Name, tc.Expect, out)
		}
	}

	// // Get struct from array
	// // resp.level2.nestedArray[3].someField
	// type someStruct struct {
	// 	SomeField string `json:"someField"`
	// }
	// var something someStruct
	// if err = arrVal[3].Unmarshal(&something); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(something.SomeField)
}
