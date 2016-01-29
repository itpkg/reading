package utils

import (
	"encoding/json"
)

func ToJson(o interface{}) ([]byte, error) {
	return json.Marshal(o)

}

func FromJson(j []byte, o interface{}) error {
	return json.Unmarshal(j, o)
}
