package pkg

import "sort"

// 求两个 Slice 的并集
func Union(a, b []interface{}) []interface{} {
	m := make(map[interface{}]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		m[item] = true
	}
	result := make([]interface{}, 0)
	if len(a) == 0 && len(b) == 0 {
		return result
	}

	for item := range m {
		result = append(result, item)
	}

	var t interface{}
	if len(a) != 0 {
		t = a[0]
	}
	if len(b) != 0 {
		t = b[0]
	}
	switch t.(type) {
	case int:
		sort.Slice(result, func(i, j int) bool {
			return result[i].(int) < result[j].(int)
		})
	case string:
		sort.Slice(result, func(i, j int) bool {
			return result[i].(string) < result[j].(string)
		})
	}
	return result
}

// 求两个 Slice 的交集
func Intersect(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}
	var result []int
	for _, item := range b {
		if m[item] {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// 求两个 Slice 的差集
func Difference(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range b {
		m[item] = true
	}
	var result []int
	for _, item := range a {
		if !m[item] {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
