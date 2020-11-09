package main

import (
	"testing"
)

func TestGetIntegersFromSlice(t *testing.T) {
	testCases := []struct {
		input []string
		output []int
		isValid bool
	}{
		// Normal use case, all valid inputs
		{
			[]string{"2", "5", "10"},
			[]int{2, 5, 10},
			true,
		},
		// Some extra invalid stuff in input
		{
			[]string{"2", "5", "foo", "10", "bar"},
			[]int{2, 5, 10},
			true,
		},
		// Both empty
		{
			[]string{},
			[]int{},
			true,
		},
		// Mismatch in input vs. output
		{
			[]string{"2", "10"},
			[]int{2, 5, 10},
			false,
		},
		// Empty input, non-empty output
		{
			[]string{},
			[]int{2},
			false,
		},
		// Non-empty input, empty output
		{
			[]string{"10"},
			[]int{},
			false,
		},
		// Input with invalid values, empty output
		{
			[]string{""},
			[]int{},
			true,
		},
	}

	for _, tc := range testCases {
		_, out := GetIntegersFromSlice(&tc.input)
		if SlicesEqual(tc.output, out) != tc.isValid {
			t.Errorf("GetIntegersFromSlice fails: in %v, out: %v, want: %v", tc.input, out, tc.output)
		}
	}
}

// SlicesEqual tells whether a and b contain the same elements. A nil argument is equivalent to an empty slice.
func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGetRandomValuesFromSlices(t *testing.T) {
	testCases := []struct {
		inM []int
		inT []int
		skipM int
		skipT int
		outM []int
		outT []int
		isValid bool
	}{
		// Normal case - valid inputs, no skipping
		{
			[]int{2, 5, 10},
			[]int{2, 3, 4, 5, 6, 7, 8, 9},
			0,
			0,
			[]int{2, 5, 10},
			[]int{2, 3, 4, 5, 6, 7, 8, 9},
			true,
		},
		// Only one possible value from both possible
		{
			[]int{2},
			[]int{2},
			0,
			0,
			[]int{2},
			[]int{2},
			true,
		},
		// Skip specific values
		{
			[]int{2, 5},
			[]int{2, 5},
			2,
			2,
			[]int{2, 5},
			[]int{2, 5},
			true,
		},
		// Impossible situation, only one possible value and those need to be skipped
		{
			[]int{10},
			[]int{ 9},
			10,
			9,
			[]int{},
			[]int{},
			false,
		},
	}

	for _, tc := range testCases {
		_, rm, rt := GetRandomValuesFromSlices(&tc.inM, &tc.inT, tc.skipM, tc.skipT)
		if tc.isValid && !SliceContains(&tc.outM, rm) && (tc.skipM != rm || tc.skipT != rt) {
			t.Errorf("GetRandomValuesFromSlices fails: Test %v, rm: %d, rt: %d", tc, rm, rt)
		}
	}
}

func SliceContains(s *[]int, e int) bool {
	for _, a := range *s {
		if a == e {
			return true
		}
	}
	return false
}