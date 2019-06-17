/**
 * @author zhangyuehao
 * @date 2019-06-17 20:43
 */

package main

import "fmt"

func main() {
	slice := []int{3, 3}
	res := twoSum(slice, 6)
	fmt.Println(fmt.Sprintf("res is :%v", res))
}

func twoSum(nums []int, target int) []int {
	checkMap := make(map[int]int)

	for i, v := range nums {
		checkMap[v] = i
	}

	indexSlice := make([]int, 0, 2)

	for i1, v1 := range nums {
		if vMap, ok := checkMap[target-v1]; ok && vMap != i1 {
			indexSlice = append(indexSlice, i1)
			indexSlice = append(indexSlice, vMap)
			return indexSlice
		}

	}

	return indexSlice
}
