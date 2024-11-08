package main

import "fmt"


func main(){
	items := []int{1, 10, 3, 24, 40, 32, 98, 10, 4, 3}
	fmt.Print(mergeSort(items))
}

func mergeSort(items []int)[]int {
	if len(items) < 2 {
		return items
	}
	first 	:= mergeSort(items[:len(items)/2])
	second 	:= mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a, b []int)[]int{
	merged := []int{}

	i:= 0
	j:= 0

	for i < len(a) && j <len(b){
		if a[i]<b[j]{
			merged = append(merged, a[i])
			i ++
		} else {
			merged = append(merged, b[j])
			j ++
		}
	}
	for ; i < len(a); i++{
		merged = append(merged, a[i])
	}
	for ; j < len(b); j++{
		merged = append(merged, b[j])
	}
	return merged
}