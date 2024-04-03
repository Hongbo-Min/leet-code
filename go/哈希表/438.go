package main

import "fmt"

func findAnagrams(s, p string) (res []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		res = append(res, 0)
	}
	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			res = append(res, i+1)
		}
	}
	return
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Printf("findAnagrams(s, p): %v\n", findAnagrams(s, p))
}
