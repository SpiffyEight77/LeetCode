package main

func main() {
	m := make(map[string]int)

	m["hello"] = 1
	println(m["hello"])

	delete(m, "hello")

	for k, v := range m {
		println(k, v)
	}
}
