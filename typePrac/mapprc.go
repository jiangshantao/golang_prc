package main

import "fmt"

func initMap(){
	dict1 := make(map[string]int)
	dict1["hehe"] = 2
	dict2 := map[string]string{
		"red": "jiangyanhong",
		"yellow": "topbuzz",
		"blue": "jiangshantao"}
	//dict3 := map[[]string]int{} trigger "invalid map key error"
	fmt.Println(dict1)
	fmt.Println(dict2)
}

func assignValueMap(){
	colors := map[string]string{
		"red": "1",
		"yellow": "2",
		"blue": "3",
		"balck": "4"}
	fmt.Println(colors)
	colors["brown"] = "5"
	fmt.Println(colors)
	value, exists := colors["red"]
	a, b := colors["red"]
	c := colors["red"]
	fmt.Println(value, exists)
	fmt.Println(a, b)
	fmt.Println(c)
	if exists{
		fmt.Println("red key has value")
	}
	if value != ""{
		fmt.Println(value)
	}
}

func rangeMap(){
	colors3 := map[string]string{
		"red": "1234",
		"yellow": "4321",
		"blue": "5432",
		"pink": "8765",
	}

	for key, value := range colors3{
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func deleteMapVlaue(){
	colors4 := map[string]string{
		"red": "12345",
		"yellow": "54321",
		"blue": "54321",
		"pink": "87654",
	}
	loopMap(colors4)
	removeMapValue(colors4, "red")
	loopMap(colors4)
}

func removeMapValue(color map[string]string, key string){
	delete(color, key)
}

func loopMap(color map[string]string){
	for key, value := range color{
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
	fmt.Println("\n")
}

func main(){
	//initMap()
	//assignValueMap()
	//rangeMap()
	deleteMapVlaue()

}