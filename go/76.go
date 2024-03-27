package main

func judge(runeAndCount map[rune]int) bool {
	for _, value := range runeAndCount {
		if value > 0 {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	var res string
	mark := make(map[rune]struct{})
	runeAndCount := make(map[rune]int)
	for _, r := range t {
		mark[r] = struct{}{}
		runeAndCount[r]++
	}
	left := 0
	for right, r := range s {
		if _, ok := mark[r]; ok {
			runeAndCount[r]--
		}
		for judge(runeAndCount) {
			subStr := s[left : right+1]
			if len(res) != 0 {
				if len(res) > len(subStr) {
					res = subStr
				}
			} else {
				res = subStr
			}
			if _, ok := mark[rune(s[left])]; ok {
				runeAndCount[rune(s[left])]++
			}
			left++
		}
	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	println(minWindow(s, t))
}
