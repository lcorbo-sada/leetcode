package main

import (
	"reflect"
	"testing"
)

func TwoSumTest(t *testing.T) {
	t.Run("Test 1 in Parallel", func(t *testing.T) {
		t.Parallel()
		input := []int{2, 7, 11, 15}
		target := 6
		output := twoSum(input, target)
		expectedResult := []int{0, 1}
		testResult := reflect.DeepEqual(output, expectedResult)
		if !testResult {
			t.Fatalf(`Wanted:" %v, "Got: %v`, testResult, expectedResult)
		}
	})
	t.Run("Test 2 in Parallel", func(t *testing.T) {
		t.Parallel()
		input := []int{3, 2, 4}
		target := 6
		output := twoSum(input, target)
		expectedResult := []int{1, 2}
		testResult := reflect.DeepEqual(output, expectedResult)
		if !testResult {
			t.Fatalf(`Wanted:" %v, "Got: %v`, testResult, expectedResult)
		}
	})
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		input := []int{3, 3}
		target := 4
		output := twoSum(input, target)
		expectedResult := []int{0, 1}
		testResult := reflect.DeepEqual(output, expectedResult)
		if !testResult {
			t.Fatalf(`Wanted:" %v, "Got: %v`, testResult, expectedResult)
		}
	})
}
