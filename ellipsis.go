/*
ellipses.go
-John Taylor
Dec-29-2020

Insert an ellipses into the middle of a long string to shorten it

*/

package ellipsis

// Shorten will insert "..." into the middle of a long string
// to shorten it to at most w Unicode code points (runes).
// Nonpositive widths return an empty string; widths 1–3 keep the last w runes.
func Shorten(s string, w int) string {
	if w <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= w {
		return s
	}
	if w <= 3 {
		return string(runes[len(runes)-w:])
	}

	maxsz := (w - 3) / 2
	extra := 0
	if w%2 == 0 {
		extra = 1
	}

	a := string(runes[0:maxsz])
	z := string(runes[len(runes)-maxsz-extra:])

	return a + "..." + z
}
