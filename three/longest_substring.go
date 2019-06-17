/**
 * @author zhangyuehao
 * @date 2019-06-17 21:23
 */

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	runeMap := make(map[rune]int)
	start := -1
	maxLenth := 0

	for i, v := range s {
		last, ok := runeMap[v]
		if ok && start < last {
			start = last
		}
		if i-start > maxLenth {
			maxLenth = i - start
		}
		runeMap[v] = i
	}

	return maxLenth
}
