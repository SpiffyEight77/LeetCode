package main

func main() {
	stack := make([]int, 0)

	stack = append(stack, 10)

	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	println(v)

	if len(stack) == 0 {
		println("stack is empty")
	}
}
