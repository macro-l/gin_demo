/***
 * 冒泡排序
 ***/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	//num = 100000
	//rand.Seed(time.Now().UnixNano())
	//sourcesPath     = "./source.txt"
	sourcesJsonPath = "./source_json.txt"
)

func main() {

	data := data(sourcesJsonPath)
	fmt.Println(data)
}

func data(sourcesJsonPath string) []int {
	sources_json, err := ioutil.ReadFile(sourcesJsonPath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var sources []int
	//var sources []struct{}
	json.Unmarshal(sources_json, &sources)
	//fmt.Println(sources)

	return sources
}
