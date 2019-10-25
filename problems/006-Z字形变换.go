package main

import "fmt"

/*
006 Z字形变换 https://leetcode-cn.com/problems/zigzag-conversion/

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/
func main() {
	fmt.Println(convert("LEETCODEISHIRING", 4))
}

// 执行用时 4 ms 内存消耗 4.1 MB
func convert(s string, numRows int) string {
	if len(s) < 2 {
		return s
	}
	if numRows == 1 {
		return s
	}
	var z []byte
	for i := 0; i < numRows; i++ {
		j := i
		if j < len(s) {
			z = append(z, s[j])
		}
		for j < len(s) {
			if 2*(numRows-i-1) != 0 {
				j = j + 2*(numRows-i-1)
				if j < len(s) {
					z = append(z, s[j])
				}
			}
			if 2*i > 0 {
				j = j + 2*i
				if j < len(s) {
					z = append(z, s[j])
				}
			}
		}
	}
	return string(z)
}
