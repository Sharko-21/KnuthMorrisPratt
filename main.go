package KnuthMorrisPratt

import (
	"strings"
)

func getPi(substr string) []int {
	var pi = make([]int, len(substr))
	pi[0] = 0

	j := 0
	i := 1

	for i < len(substr) {
		if substr[i] != substr[j] {
			if j == 0 {
				pi[i] = 0
				i++
				continue
			} else {
				j = pi[j - 1]
				continue
			}
		} else {
			pi[i] = j + 1
			i++
			j++
			continue
		}
	}

	return pi
}

func searchImageIndexInStr(str string, substr string, pi []int) int {
	k := 0
	i := 0
	for k < len(str) && i < len(substr) {
		if str[k] == substr[i] {
			if i == len(substr) - 1 {
				return len(str) - len(substr)
			}
			k++
			i++
			continue
		} else {
			if i != 0 {
				i = pi[i - 1]
				continue
			} else {
				k++
			}
		}
	}
	return - 1
}

func Index(str string, substr string) int {
	substrLen := len(substr)

	switch {
	case substrLen == 0:
		return 0
	case substrLen == 1:
		return strings.IndexByte(str, substr[0])
	case substrLen == len(str):
		if substr == str {
			return 0
		}
		return -1
	case substrLen > len(str):
		return -1
	case true /*here must be max str length*/:
		pi := getPi(substr)
		return searchImageIndexInStr(str, substr, pi)
	}
	return 0
}

func Contains(str string, substr string) bool {
	return Index(str, substr) >= 0
}
