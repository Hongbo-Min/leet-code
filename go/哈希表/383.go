package main

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, b := range magazine {
		magazineMap[b]++
	}
	for _, b := range ransomNote {
		if num, ok := magazineMap[b]; ok {
			if num < 1 {
				return false
			}
			magazineMap[b]--
		} else {
			return false
		}
	}
	return true
}

func main() {
}
