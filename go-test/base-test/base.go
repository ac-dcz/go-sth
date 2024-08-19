package base

import (
	"strings"
)

func Split(s, seq string) []string {
	var temp []string
	if seq == "" {
		return explode(s)
	}

	for {
		index := strings.Index(s, seq)
		if index < 0 {
			break
		}
		temp = append(temp, s[:index])
		s = s[index+len(seq):]
	}
	if len(s) > 0 {
		temp = append(temp, s)
	}
	return temp
}

func explode(s string) []string {
	var temp []string
	for i := range s {
		temp = append(temp, s[i:i+1])
	}
	return temp
}
