package shape

type Font struct {
	Height     int
	LineHeight int

	charWidths map[rune]float32
}

func (f Font) TextWidth(txt string) int {
	var width float32
	for _, r := range txt {
		w, found := f.charWidths[r]
		if !found {
			w = 8.0
		}
		width += float32(w) * float32(f.Height) / float32(12)
	}
	return int(width)
}

// font-size: 12px
var arial = map[rune]float32{
	'A': 8,
	'B': 8,
	'C': 8,
	'D': 8,
	'E': 8,
	'F': 7,
	'G': 9,
	'H': 8,
	'I': 3,
	'J': 6,
	'K': 8,
	'L': 7,
	'M': 9,
	'N': 8,
	'O': 9,
	'P': 8,
	'Q': 9,
	'R': 8,
	'S': 8,
	'T': 7,
	'U': 8,
	'V': 8,
	'X': 8,
	'Y': 8,
	'Z': 7,
	'a': 7,
	'b': 7,
	'c': 6,
	'd': 7,
	'e': 7,
	'f': 3,
	'g': 7,
	'h': 7,
	'i': 2,
	'j': 2,
	'k': 6,
	'l': 2,
	'm': 9,
	'n': 7,
	'o': 7,
	'p': 7,
	'q': 7,
	'r': 3,
	's': 6,
	't': 3,
	'u': 7,
	'v': 6,
	'x': 6,
	'y': 6,
	'z': 6,
	'.': 3,
	'_': 7,
	'1': 7,
	'2': 7,
	'3': 7,
	'4': 7,
	'5': 7,
	'6': 7,
	'7': 7,
	'8': 7,
	'9': 7,
	'0': 7,
	' ': 3,
}
