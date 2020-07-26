package lengthOfLongestSubstring

func LengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	dictIndex := make(map[string]int)
	head := 0
	longestSubstring := 0
	var index int
	var ok bool
	for i := 0; i < length; i++ {
		strValue := string(s[i])
		index, ok = dictIndex[strValue]
		if !ok {
			dictIndex[strValue] = i
			if (i - head + 1) > longestSubstring {
				longestSubstring = i - head + 1
			}
		} else {
			for j := head; j < index; j++ {
				delete(dictIndex, string(s[j]))
			}
			dictIndex[string(s[index])] = i
			head = index + 1
		}
	}
	return longestSubstring
}