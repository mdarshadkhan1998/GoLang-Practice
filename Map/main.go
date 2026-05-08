package main

import "fmt"

func main() {
	map1 := map[string]string{"b": "boy", "c": "cat"}
	map1["a"] = "Arshad"
	fmt.Println(map1)
	delete(map1, "b")
	fmt.Println(map1)

	map2 := make(map[string]string)
	map2["a"] = "antman"
	map2["b"] = "batman"
	map2["c"] = "catman"
	fmt.Print(map2)
	print(map2)
	fmt.Print(map2)
}

func print(m map[string]string) {
	m["c"] = "car"
}
