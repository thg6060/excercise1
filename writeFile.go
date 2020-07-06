package main

import (
	"encoding/json"
	"io/ioutil"
)

func writeFile() error{
	slServe :=[]Serve{
		Serve{ Name: "cofaxAdmin",Class: "org.cofax.cds.AdminServlet" },
		Serve{ Name: "fileServlet",Class: "fileServletfileServlet"  },
	}
	result, err := json.MarshalIndent(slServe, "", "    ")
	ioutil.WriteFile("serve.json", result, 0644)
	if err != nil {
		return err
	}
	return nil

}
