package draw

import "os"

func ExampleNewSvg() {
	s := NewSVG()
	s.WriteSVG(os.Stdout)
	// output:
	// <svg
	//   xmlns="http://www.w3.org/2000/svg"
	//   xmlns:xlink="http://www.w3.org/1999/xlink"
	//   class="root" width="100" height="100"></svg>
}
