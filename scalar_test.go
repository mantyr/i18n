package i18n

import "testing"

func TestScalarToString(t *testing.T) {
	cases := []struct {
		in   any
		want string
		ok   bool
	}{
		{"hello", "hello", true},
		{42, "42", true},
		{int64(42), "42", true},
		{3.14, "3.14", true},
		{true, "true", true},
		{[]int{1}, "", false},
		{nil, "", false},
	}
	for _, tc := range cases {
		got, err := scalarToString(tc.in)
		if tc.ok && err != nil {
			t.Errorf("%v: unexpected error %v", tc.in, err)
		}
		if !tc.ok && err == nil {
			t.Errorf("%v: expected error, got %q", tc.in, got)
		}
		if tc.ok && got != tc.want {
			t.Errorf("%v: got %q want %q", tc.in, got, tc.want)
		}
	}
}
