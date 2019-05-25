package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ServiceMap []struct {
	Service struct {
		StringPrefix  string   `json:"StringPrefix"`
		Actions       []string `json:"Actions"`
		ARNFormat     string   `json:"ARNFormat"`
		ARNRegex      string   `json:"ARNRegex"`
		ConditionKeys []string `json:"conditionKeys"`
		HasResource   bool     `json:"HasResource"`
	} `json:"Amazon Comprehend"`
}

func main()  {
	// read file into string
	b, err := ioutil.ReadFile("policy.js")
	if err != nil {
		fmt.Print(err)
	}
	// chop off heading js code; 'app.PolicyEditorConfig='
	str := string(b)[len("app.PolicyEditorConfig="):]
	//fmt.Println(str)

	//var f interface{}

	var results map[string]interface{}

	err = json.Unmarshal([]byte(str), &results)
	if err != nil {
		fmt.Print(err)
	}
	serviceMap := results["serviceMap"].(map[string]interface{})
	for key, result := range serviceMap {
		println(key, result)
		prefix := result.(map[string]interface{})["StringPrefix"]
		actions := result.(map[string]interface{})["Actions"]
		for _, action := range actions.([]interface{}) {
			println(action.(string), prefix.(string))
		}



	}
	//for key, result := range results {
	//	address := result["serviceMap"].(map[string]interface{})
	//	fmt.Println("Reading Value for Key :", key)
	//	//Reading each value by its key
	//	fmt.Println("Id :", result["id"],
	//		"- Name :", result["name"],
	//		"- Department :", result["department"],
	//		"- Designation :", result["designation"])
	//	fmt.Println("Address :", address["city"], address["state"], address["country"])
	//}



}

