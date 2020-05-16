package logalizer

import (
	"sort"
)

type kv struct {
	Key   string
	Value int
}

// SortMapByValue ... sort map by value
func SortMapByValue(values map[string]int) []kv {
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	if len(ss) < 10 {
		return ss[:len(ss)]
	} else {
		return ss[:10]
	}
}
