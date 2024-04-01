package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mark := map[string][]string{}
	for _, str := range strs {
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})
		sortedStr := string(strBytes)
		mark[sortedStr] = append(mark[sortedStr], str)
	}
	res := make([][]string, 0, len(mark))
	for _, val := range mark {
		res = append(res, val)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("groupAnagrams(strs): %v\n", groupAnagrams(strs))
}
