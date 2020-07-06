package main

import (
	"encoding/json"
	"io/ioutil"
)

type Serve struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func readFile() ([]byte, error) {
	result, err := ioutil.ReadFile("serve.json")
	if err != nil {
		return nil, err
	}
	return result, nil
}
func toSlice(data []byte) ([]Serve, error) {
	sl := make([]Serve, 2)
	err := json.Unmarshal(data, &sl)
	return sl, err
}
