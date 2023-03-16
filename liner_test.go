package liner_test

import (
	"testing"

	"github.com/isksss/go-liner"
)

func TestMakeLine(t *testing.T) {
	tests := []struct {
		name string
		len  int
		moji string
		want string
	}{
		{name: "dafault", len: 10, moji: "#", want: "##########"},
		{name: "length is 0", len: 0, moji: "#", want: ""},
		{name: "length is 20", len: 20, moji: "#", want: "####################"},
		{name: "moji is ##", len: 10, moji: "##", want: "####################"},
	}

	for _, test := range tests {
		got := liner.MakeLine(test.moji, test.len)
		if got != test.want {
			t.Errorf("%v: want %v, but %v:", test.name, test.want, got)
		}
	}
}
