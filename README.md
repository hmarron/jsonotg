# jsonotg

### Description

A small library for getting values out of a json object without needing to define a load of structs.


### Usage

Load a `[]byte` containing json into a `jsonotg.JSON`
```
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

j, err := jsonotg.CreateJSON([]byte(jsonStr))
if err != nil {
    panic(err)
}
```


Fields can be indexed by chaining `GetField`, followed by an `AsType`
```
boolVal, ok := j.
	GetField("level2").
	GetField("key3").
	AsBool()

fmt.Println(boolVal, ok)
```


Slices work as follows:
```
arrVal, ok := j.
	GetField("level2").
	GetField("nestedArray").
	AsArray()

fmt.Println(arrVal, ok)

// Get a string from the array
stringVal, ok = arrVal[1].AsString()
fmt.Println(stringVal, ok)
```

Getting a struct from the array 
(useful if you want to get on deeply nested struct out of a json object, but don't want to define structs for all the levels above it that we don't care about)
```
type someStruct struct {
	SomeField string `json:"someField"`
}
var something someStruct
if err = arrVal[3].Unmarshal(&something); err != nil {
    panic(err)
}
fmt.Println(something.SomeField)
```
