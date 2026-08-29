package main

import "testing"

// TestFloatToIntTruncates locks in the "truncate toward zero, never
// round" behavior described in the README, for both positive and
// negative inputs.
func TestFloatToIntTruncates(t *testing.T) {
	cases := []struct {
		name string
		in   float64
		want int
	}{
		{"just below a whole number", 3.99, 3},
		{"just above a whole number", 3.01, 3},
		{"negative, truncates toward zero", -3.99, -3},
		{"exactly whole", 4.0, 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := int(tc.in)
			if got != tc.want {
				t.Errorf("int(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

// TestNarrowingConversionWraps locks in the exact wrapped values a
// narrowing conversion produces, so a future change to this lesson (or
// a misunderstanding while editing it) is caught immediately.
func TestNarrowingConversionWraps(t *testing.T) {
	var big int32 = 300
	if got := int8(big); got != 44 {
		t.Errorf("int8(int32(300)) = %d, want 44", got)
	}

	var big2 int64 = 40000
	if got := int16(big2); got != -25536 {
		t.Errorf("int16(int64(40000)) = %d, want -25536", got)
	}
}
