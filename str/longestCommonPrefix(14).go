package str

import "strings"

/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	result := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], result) {
			if len(result) == 0 {
				return ""
			}
			result = result[:len(result)-1]
		}
	}

	return result
}
