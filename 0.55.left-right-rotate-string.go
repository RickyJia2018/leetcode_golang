package main

import "fmt"

func main() {
	str := "abc"
	leftRotate(str, 2)
	leftRotate(str, 5)
	rightRotate(str, 2)
	rightRotate(str, 5)

}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

/*
字符串的右旋转操作是把字符串尾部的若干个字符转移到字符串的前面。给定一个字符串 s 和一个正整数 k，请编写一个函数，将字符串中的后面 k 个字符移到字符串的前面，实现字符串的右旋转操作。

例如，对于输入字符串 "abcdefg" 和整数 2，函数应该将其转换为 "fgabcde"。

输入：输入共包含两行，第一行为一个正整数 k，代表右旋转的位数。第二行为字符串 s，代表需要旋转的字符串。

输出：输出共一行，为进行了右旋转操作后的字符串。

样例输入：

2
abcdefg
样例输出：

fgabcde
*/

func rightRotate(s string, n int) {
	fmt.Printf("Right Rotate %d times: \n", n)

	if n <= 0 {
		return
	}
	n = n % len(s)
	b := []byte(s)
	reverse(b)
	reverse(b[:n])
	reverse(b[n:])
	fmt.Printf("From %v -> %v\n", s, string(b))
}

/*
2. The left rotation operation of a string is to transfer several characters in front of the
string to the end of the string. Please define a function to realize the function of string
left rotation operation. For example, if you input the string "abcdefg" and the number
2, the function will return the result "cdefgab" obtained by rotating left by two bits.
Example 1:
Input: s = "abcdefg", k = 2
Output: "cdefgab"
Example 2:
Input: s = "lrloseumgh", k = 6
Output: "umghlrlose"
public string reverseLeftWords(string s, int n) {
}
*/
func leftRotate(s string, n int) {
	fmt.Printf("Left Rotate %d times: \n", n)
	if n <= 0 {
		return
	}
	n = n % len(s)
	b := []byte(s)
	reverse(b)
	reverse(b[:len(b)-n])
	reverse(b[len(b)-n:])
	fmt.Printf("From %v -> %v\n", s, string(b))
}
