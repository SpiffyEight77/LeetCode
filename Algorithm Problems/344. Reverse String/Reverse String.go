package main

//v1 O(N)
func reverseString(s string) string {
	res := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		res[i], res[len(s)-i-1] = res[len(s)-i-1], res[i]
	}
	return string(res)
}
