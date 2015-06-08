/* github.com/rubicks/pngcrunch/pngcrunch_test.go */

package main

import (
	"testing"
)

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFixargs(t *testing.T) {

	cases := []struct {
		in, want []string
	}{
		{[]string{}, []string{"-", "-"}},
		{[]string{"foo.png"}, []string{"foo.png", "-"}},
		{[]string{"foo.png", "bar.png"}, []string{"foo.png", "bar.png"}},
		{[]string{"-", "bar.png"}, []string{"-", "bar.png"}},
		{[]string{"foo.png", "-"}, []string{"foo.png", "-"}},
	}
	for _, c := range cases {
		got := fixargs(c.in)
		if !eq(got, c.want) {
			t.Errorf("fixargs(%#v) == %#v, want %#v\n", c.in, got, c.want)
		}
	}
}
