package main

import (
	"encoding/json"
	
)

func addToSlice(sl *[]Serve) {
	str := `{
	"name": "fileCustome",
	"class": "org.cofax.cds.FileServlet.Custome"
  }`
	serve := Serve{}
	json.Unmarshal([]byte(str), &serve)
	*sl = append(*sl, serve)
}
