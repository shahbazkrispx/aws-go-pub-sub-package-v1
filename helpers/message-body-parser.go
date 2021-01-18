package helpers

import (
	"encoding/json"
	"fmt"
)


func ParseBody(d string) ([]string, error) {
	var data map[string]interface{}
	var res []string

	err := json.Unmarshal([]byte(d), &data)
	if err != nil {
		return nil, err
	}

	for _, d2 := range data["MessageAttributes"].(map[string]interface{}) {
		//fmt.Println(d2.(map[string]interface{}))
		res = append(res, fmt.Sprintf("%v",d2.(map[string]interface{})["Value"]))

	}
	return res, nil
}