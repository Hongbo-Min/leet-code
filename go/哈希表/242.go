package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	for _, b := range s {
		sMap[b]++
	}
	for _, b := range t {
		if num, ok := sMap[b]; ok {
			if num < 1 {
				return false
			}
			sMap[b]--
		} else {
			return false
		}
	}
	return true
}

func main() {
}
