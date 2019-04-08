package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "ccmouse",
		"course": "golang",
		"site":   "imooc",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	courseName, ok = m["course22"]
	fmt.Println(courseName, ok)

	fmt.Println("--------------")
	if k, ok := m["course"]; ok {
		fmt.Println(k, ok)
	} else {
		fmt.Println("key does not exist!")
	}

	delete(m, "course")

	if k, ok := m["course"]; ok {
		fmt.Println(k, ok)
	} else {
		fmt.Println("key does not exist!")
	}
}
