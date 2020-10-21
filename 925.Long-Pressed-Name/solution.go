package leetcode

func isLongPressedName(name string, typed string) bool {
	var i, j int
	for i < len(typed) {
		if j < len(name) && typed[i] == name[j] {
			i++
			j++
		} else if i > 0 && typed[i] == typed[i-1] {
			i++
		} else {
			return false
		}
	}

	return j == len(name)
}
