package main

func main() {
	queue := make([]int, 0)

	queue = append(queue, 10)

	v := queue[0]
	queue = queue[1:]
	println(v)

	if len(queue) == 0 {
		println("queue is empty")
	}
}
