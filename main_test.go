package main

import "testing"

func TestMaskedStr(t *testing.T) {
	testCases := []struct {
		name     string
		origin   string
		hs, ts   int
		expected string
	}{
		{
			name:   "empty origin",
			origin: "", hs: 0, ts: 0,
			expected: "",
		},
		{
			name:   "full masked",
			origin: "ab", hs: 0, ts: 0,
			expected: "**",
		},
		{
			name:   "head size within length",
			origin: "ab", hs: 1, ts: 0,
			expected: "a*",
		},
		{
			name:   "tail size within length",
			origin: "ab", hs: 0, ts: 1,
			expected: "*b",
		},
		{
			name:   "head size equal length",
			origin: "ab", hs: 2, ts: 0,
			expected: "ab",
		},
		{
			name:   "tail size equal length",
			origin: "ab", hs: 0, ts: 2,
			expected: "ab",
		},
		{
			name:   "head size over length",
			origin: "ab", hs: 3, ts: 0,
			expected: "ab",
		},
		{
			name:   "tail size over length",
			origin: "ab", hs: 0, ts: 3,
			expected: "ab",
		},
	}
	for _, tc := range testCases {
		actual := MaskedStr(tc.origin, tc.hs, tc.ts)
		if actual != tc.expected {
			t.Errorf("MaskedStr failed. case: %s, actual: %s", tc.name, actual)
		}
	}
}
