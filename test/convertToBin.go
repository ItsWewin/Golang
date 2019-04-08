package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(123312121),
		convertToBin(0),
	)

	printFile("test.txt")

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6}
	var arr4 [4][5]int

	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	for i := range arr2 {
		fmt.Println(arr2[i])
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	for i, v := range arr4 {
		fmt.Println(i, v)
	}

	var arr5 = [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr5[2:5]
	s2 := s1[1:4]
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

	fmt.Println("-------------")

	fmt.Println("len =", len(s2), "cap =", cap(s2))
	fmt.Println(s2[0], s2[1], s2[2])

	s3 := s2[:]
	fmt.Println("s3 =", s3)

	s4 := s2[:4]
	fmt.Println("s4 =", s4)

	// s5 := s2[:5]
	// fmt.Println("s5 =", s5)

	fmt.Println("--------------")
	fmt.Println("s2 =", s2, "len =", len(s2), "cap =", cap(s2))
	s6 := append(s2, 10)
	fmt.Println("s6 =", s6, "len =", len(s6), "cap =", cap(s6))
	s7 := append(s2, 10, 11, 12, 13, 14, 15, 16, 17)
	fmt.Println("s7 =", s7, "len =", len(s7), "cap =", cap(s7))

	fmt.Println("s2 =", s2, "len =", len(s2), "cap =", cap(s2))
}
