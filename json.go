package jsonotg

import (
	"encoding/json"
)

type JSON struct {
	raw interface{}
}

func CreateJSON(data []byte) (*JSON, error) {
	var j JSON
	if err := json.Unmarshal(data, &j.raw); err != nil {
		return nil, err
	}
	return &j, nil
}

func (j *JSON) GetField(fieldName string) *JSON {
	return &JSON{j.raw.(map[string]interface{})[fieldName]}
}

func (j *JSON) AsArray() ([]*JSON, bool) {
	is, ok := j.raw.([]interface{})
	if !ok {
		return nil, false
	}

	val := make([]*JSON, 0, len(is))

	for _, i := range is {
		val = append(val, &JSON{i})
	}

	return val, true
}

func (j *JSON) AsString() (string, bool) {
	val, ok := j.raw.(string)
	return val, ok
}

func (j *JSON) AsInt64() (int64, bool) {
	f, ok := j.raw.(float64)
	if !ok {
		return 0, false
	}

	return int64(f), true
}

func (j *JSON) AsBool() (bool, bool) {
	val, ok := j.raw.(bool)
	return val, ok
}

func (j *JSON) Unmarshal(v any) error {
	// Convert map back to json string
	jsonString, err := json.Marshal(j.raw)
	if err != nil {
		return err
	}

	// Convert json string to struct
	return json.Unmarshal(jsonString, &v)
}

func (j *JSON) IsNull() bool {
	return j.raw == nil
}
