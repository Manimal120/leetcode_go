package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	var freq [256]int

	res, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//var freq [256]int
	s := "1bc"
	fmt.Println(s[0], 'a', s[0]-'a')
	fmt.Printf("%T, %T, %T", s[0], 'a', s[0]-'a')
}
