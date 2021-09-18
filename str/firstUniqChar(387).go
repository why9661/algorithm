package str

/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
*/

func firstUniqChar(s string) int {
label:
	for i, char1 := range s {
		for _, char2 := range s[i+1:] {
			if char1 == char2 {
				continue label
			}
		}
		return i
	}

	return -1
}
