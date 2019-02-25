package main

import (
	"fmt"
	"sort"
	"strings"
)

func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice) + len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])

	return result
}

// func InsertStringSlice(slice, insertion []string, index int) []string {
// 	return append(slice[:index], append(insertion, slice[index:]...)...)
// }

func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice) - (end - start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

func SortFoldedStrings(slice []string) {
	sort.Sort(FoldedStrings(slice))
}

type FoldedStrings []string

func (slice FoldedStrings) Len() int {
	return len(slice)
}

func (slice FoldedStrings) Less(i, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedStrings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := []string{"K", "L", "M", "N"}
	u := []string{"m", "n", "o", "p", "q", "r"}

	s = append(s, "h", "i", "j")
	s = append(s, t...)
	s = append(s, u[2:5]...)
	fmt.Print(s, "\n")

	k := InsertStringSliceCopy(s, []string{"1", "2", "3"}, 0)
	fmt.Print(k, "\n")

	str := []string{"A", "B", "C", "D", "E", "F", "G"}
	j := RemoveStringSliceCopy(str, 1, 5)
	fmt.Print("j: ", j, "\n")

	str2 := []int{1, 2, 4, 5, 6, 7, 8}
	fmt.Print(str2[:2], "\n")
	fmt.Print(str2, "\n")
	fmt.Print(str2[4:], "\n")
	fmt.Print(str2, "\n")

	str2_slice := append(str2[:2], str2[4:]...)
	fmt.Print("test apend:", str2_slice, "\n")
	fmt.Print(str2, "\n")

	files := []string{"Test.conf", "util.go", "MakeFile", "misc.go", "main.go"}
	fmt.Printf("Unsorted: %q\n", files)
	sort.Strings(files)
	fmt.Printf("Underlying bytes: %q\n", files)
	SortFoldedStrings(files)
	fmt.Printf("Case insensitive: %q\n", files)
}