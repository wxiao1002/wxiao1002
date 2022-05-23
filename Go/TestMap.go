package main

import "fmt"

func main() {

	// 创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	// 插入值
	countryCapitalMap["cc"] = "cc"
	countryCapitalMap["aa"] = "aa"
	countryCapitalMap["bb"] = "bb"
	// 便利
	for countrt := range countryCapitalMap {
		fmt.Println(countrt, countryCapitalMap[countrt])
	}

	// 查看元素是否在集合中存在
	capital, ok := countryCapitalMap["123"]
	if ok {
		fmt.Println("yes", capital)
	} else {
		fmt.Println("no")
	}

	/**
	创建集合
	*/

	countryMap := map[string]string{"aa": "1", "bb": "2"}
	fmt.Println(countryMap)
	// 便利
	for key_ := range countryMap {
		fmt.Println(key_, countryMap[key_])
	}
	// 删除
	delete(countryMap, "aa")
	// 便利
	for key_ := range countryMap {
		fmt.Println(key_, countryMap[key_])
	}
}
