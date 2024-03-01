package main

import (
	"leetcode/go-helper"
	"strconv"
)

func main() {
	helper.PrintJSON(compareVersion("1.01", "1.001")) // 0
	helper.PrintJSON(compareVersion("1.0", "1.0.0")) // 0
	helper.PrintJSON(compareVersion("0.1", "1.1")) // -1
}

func compareVersion(version1 string, version2 string) int {
	s1, s2 := 0, 0
	for ; s1 < len(version1) && s2 < len(version2); {
		p1 := s1
		for p1 < len(version1) && version1[p1] != '.' {
			p1++
		}
		v1, _ := strconv.Atoi(version1[s1 : p1])
		p2 := s2
		for p2 < len(version2) && version2[p2] != '.' {
			p2++
		}
		v2, _ := strconv.Atoi(version2[s2 : p2])
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
		s1 = p1 + 1
		s2 = p2 + 1
	}
	if s1 < len(version1) {
		v1 := ""
		for s1 < len(version1) {
			if version1[s1] == '.' {
				if val, _ := strconv.Atoi(v1); val > 0 {
					return 1
				}
				v1 = ""
			} else {
				v1 += string(version1[s1])
			}
			s1++
		}
		if v1 != "" {
			if val, _ := strconv.Atoi(v1); val > 0 {
				return 1
			}
		}
	}
	if s2 < len(version2) {
		v2 := ""
		for s2 < len(version2) {
			if version2[s2] == '.' {
				if val, _ := strconv.Atoi(v2); val > 0 {
					return -1
				}
				v2 = ""
			} else {
				v2 += string(version2[s2])
			}
			s2++
		}
		if v2 != "" {
			if val, _ := strconv.Atoi(v2); val > 0 {
				return -1
			}
		}
	}
	return 0
}