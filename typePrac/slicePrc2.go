package main

import "fmt"


func printSlcie(x []int, name string){
	fmt.Printf("%v len=%d cap=%d slice=%v\n", name, len(x), cap(x), x)
}

//使用append()将两个切片追加到一起
func AppendThreePoint(){
	sl1 := []int{1, 2}
	printSlcie(sl1, "sl1")
	sl2 := []int{3, 4}
	sl3 := append(sl1, sl2...)
	printSlcie(sl3, "sl3")
}

//切片迭代 for index, value := range slice{...}
func loopSlice(){
	slice1 := []int{1, 2, 3, 4, 5}
	for index, value := range slice1{
		fmt.Println(index, value)
	}
}

var Slice = make([]int, 5)
func loopSlicePoint(){
	slice1 := []int{1, 2, 3, 4, 5}
	for index, value := range slice1{
		fmt.Printf("Value: %d, ValueAddr: %X ElemAddr: %X\n",value, &value, &slice1[index])
	}
}

func loopBlankSlice(){
	Slice := []int{1, 2, 3, 4, 5}
	for _, value := range Slice{
		fmt.Printf("Value: %d\n", value)
	}
}

func traditionalLoopSlice(){
	slice2 := []int{1, 2, 3, 4, 5}
	for index:=2; index<len(slice2); index++{
		fmt.Printf("Index: %d Value: %d\n", index, slice2[index])
	}
}

func dimenSlice(){
	slice3 := [][]int{{10}, {100, 200}}
	slice3[0] = append(slice3[0], 20)
	//printSlcie(slice3, "slice3")
	fmt.Print(slice3)
}

func main(){
	//AppendThreePoint()
	//loopSlice()
	//loopSlicePoint()
	loopBlankSlice()
	traditionalLoopSlice()
	dimenSlice()
}