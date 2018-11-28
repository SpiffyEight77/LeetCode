package main

func validateStackSequences(pushed []int, popped []int) bool {
	index := 0
	res := make([]int, 0, len(popped))
	for i := 0; i < len(popped); i++ {
		res = append(res, pushed[i])
		if len(res) > 0 && index < len(popped) && popped[index] == res[len(res)-1] {
			for {
				if len(res) <= 0 || index >= len(popped) || popped[index] != res[len(res)-1] {
					break
				}
				res = append(res[:len(res)-1])
				index++
			}
		}
	}
	if len(res) == 0 {
		return true
	}
	return false
}
