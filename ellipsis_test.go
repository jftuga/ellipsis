package ellipsis

import "testing"

func TestShorten(t *testing.T) {
	tests := []struct {
		name string
		s    string
		w    int
		want string
	}{
		{"empty", "", 4, ""},
		{"already fits", "abc", 5, "abc"},
		{"exact width", "abcdef", 6, "abcdef"},
		{"negative short", "abc", -1, ""},
		{"negative long", "abcdef", -1, ""},
		{"zero short", "abc", 0, ""},
		{"zero long", "abcdef", 0, ""},
		{"one", "abc", 1, "c"},
		{"two", "abc", 2, "bc"},
		{"three", "abcdef", 3, "def"},
		{"four", "abcde", 4, "...e"},
		{"odd width", "abcdefgh", 5, "a...h"},
		{"even width", "abcdefgh", 6, "a...gh"},
		{"unicode fits", "éééé", 5, "éééé"},
		{"unicode suffix", "abcdeé", 1, "é"},
		{"unicode middle", "你好世界再见", 5, "你...见"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shorten(tt.s, tt.w); got != tt.want {
				t.Errorf("Shorten(%q, %d) = %q; want %q", tt.s, tt.w, got, tt.want)
			}
		})
	}
}
