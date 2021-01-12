package palindrome

/*
 * 125. 验证回文串
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 */

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1

	for start < end {
		for start < end && !isalnum(s[start]) {
			start += 1
		}
		for start < end && !isalnum(s[end]) {
			end -= 1
		}

		if s[start] != s[end] {
			return false
		}
		start +=  1
		end -= 1
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
