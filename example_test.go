package draw

import "os"

func ExampleNewSvg() {
	s := NewSVG()
	s.WriteSVG(os.Stdout)
	// output:
	// <svg
	//   xmlns="http://www.w3.org/2000/svg"
	//   xmlns:xlink="http://www.w3.org/1999/xlink"
	//   width="100" height="100" font-family="Arial, Helvetica, sans-serif"></svg>
}
