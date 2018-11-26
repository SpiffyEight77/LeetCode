package main

//v1 O(N) O(N)
func minAddToMakeValid(S string) int {
	top := 0
	res := make([]string, 0, len(S))
	if len(S) == 0 {
		return 0
	}

	res = append(res, string(S[0]))
	for i := 1; i < len(S); i++ {
		if top != -1 && res[top] == "(" && string(S[i]) == ")" {
			res = append(res[:top])
			top--
		} else {
			res = append(res, string(S[i]))
			top++
		}
	}
	return len(res)
}

//v2 O(N) O(1)
