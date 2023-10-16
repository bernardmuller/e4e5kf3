package main

import (
	"testing"
)

func TestIndexOfFile(t *testing.T) {
	var number_failed int = 0
	var found_indexes []int
	test_files := []File{"a", "b", "c", "x", "e", "f", "g", "u"}
	t.Run("IndexOfFile returns index of File", func(t *testing.T) {
		for _, file := range test_files {
			index, err := IndexOfFile(file, FILES)
			if err != nil {
				number_failed++
				continue
			} else {
				found_indexes = append(found_indexes, index)
			}
		}
		if len(found_indexes) != 6 {
			t.Errorf("IndexOfFile returned incorrect index")
		}
		if number_failed != 2 {
			t.Errorf("%d tests failed", number_failed)
		}
	})
}
