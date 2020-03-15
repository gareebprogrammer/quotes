package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type qdata struct {
	quote  string
	author string
}

func parseJSON(path string) []qdata {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var results []map[string]interface{}
	json.Unmarshal(file, &results)
	var data []qdata
	temp := qdata{}
	for _, result := range results {
		temp.author = fmt.Sprintf("%v", result["author"])
		temp.quote = fmt.Sprintf("%v", result["quote"])
		data = append(data, temp)
	}
	return data
}
