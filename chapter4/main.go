package main

import (
	"fmt"
	"unicode"
)

//练习4.3
func reverse(ptr *[7]int) {
	for i, j := 0, len(*ptr)-1; i < j; i, j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}

}
func test4_3() {
	ints := [...]int{1, 2, 3, 4, 5, 6, 7}
	reverse(&ints)
	fmt.Println(ints)
}

//练习4.4,将index处及以前的元素循环左移
func rotate(s []int, index int) []int {
	r := s[index:]
	for i := index - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func test4_4() {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ints)
	fmt.Println(rotate(ints, 2))
}

//练习4.5
func unique(s []string) []string {
	if len(s) < 1 {
		return s
	}
	current := s[0]
	position := 0

	for i := 1; i < len(s); i++ {
		if s[i] != current {
			position++
			s[position] = s[i]
			current = s[i]
		}
	}
	return s[:position+1]
}

func test4_5() {
	str := []string{"abc", "abc", "abc", "def", "qqq"}
	fmt.Println(str)
	fmt.Println(unique(str))
}

//练习4.6 将utf的字节slice所有相邻unicode空白字符缩减为一个ASCII空白字符
func deleteEmpty(s []rune) []rune {
	position := 0
	tag := unicode.IsSpace(s[0])
	for i := 1; i < len(s); i++ {
		if unicode.IsSpace(s[i]) {

		}
	}
}

func test4_6() {

}

func main() {
	fmt.Println("vim-go")
	//test4_3()
	//test4_4()
	test4_5()
}
