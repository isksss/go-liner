package liner_test

import "testing"

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
}
