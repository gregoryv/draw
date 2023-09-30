/*
Package goviz provides means to generate Go diagrams.

To generate a simple sequence diagram from within your code use

	func x() {
	  goviz.SaveFrameSeq("x_sequence.svg")
	}

or if you want to tweek the diagram before saving

	funx x() {
	  d, _ := goviz.FrameSequence()
	  // modify diagram d
	  _ = d.SaveAs("x_sequence.svg")
	}
*/
package goviz

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/gregoryv/draw/design"
)

// SaveFrameSeq write a sequence diagram to the given file based on
// the caller runtime frames.
func SaveFrameSeq(filename string) error {
	skip := 3 // this func + frameSeq + runtime.Callers
	d, err := frameSeq(skip)
	if err != nil {
		return err
	}
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = fh.Write([]byte(d.Inline()))
	return err

}

// FrameSequence returns a sequence diagram based on the caller
// runtime frames.
func FrameSequence() (*design.SequenceDiagram, error) {
	return frameSeq(3)
}

// frameSeq returns a sequence diagram to the given writer
func frameSeq(skip int) (*design.SequenceDiagram, error) {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No PCs available. This can happen if the first argument to
		// runtime.Callers is large.
		//
		// Return now to avoid processing the zero Frame that would
		// otherwise be returned by frames.Next below.
		return nil, fmt.Errorf("no frames")
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	cframes := runtime.CallersFrames(pc)

	// store frames once for multipass processing when building the
	// diagram
	var frames []runtime.Frame
	for {
		frame, more := cframes.Next()
		skip--
		if skip >= 0 {
			continue
		}
		frames = append(frames, frame)
		if !more {
			break
		}
	}
	// remove last (runtime.goexit)
	frames = frames[:len(frames)-1]
	// sequence starts in reverse order
	reverse(frames)

	// build diagram
	d := design.NewSequenceDiagram()
	d.SetCaption(frames[0].Function)

	// so we know which columns have been added
	cache := make(map[string]bool)

	// add columns
	for _, frame := range frames {
		col := column(frame.Function)
		if _, found := cache[col]; found {
			continue
		}
		cache[col] = true
		d.Add(col)
	}
	// link frames, the first one is used as caption
	for i := 1; i < len(frames); i++ {
		from := column(frames[i-1].Function)
		_, _, fn := ffsplit(frames[i].Function)
		to := column(frames[i].Function)
		d.Link(from, to, fn)
	}
	return d, nil
}

func column(v string) string {
	pkg, rcv, _ := ffsplit(v)
	if rcv != "" {
		if rcv[0] == '*' {
			return "*" + path.Base(pkg) + "." + rcv[1:]
		}
		return path.Base(pkg) + "." + rcv
	}
	return path.Base(pkg)
}

func ffsplit(v string) (pkg, receiver, fn string) {
	pkg = fqp(v)
	parts := strings.Split(v[len(pkg)+1:], ".")

	switch len(parts) {
	case 1:
		fn = parts[0]
	case 2:
		receiver = parts[0]
		fn = parts[1]
	}
	receiver = strings.Trim(receiver, "()")
	return

}

// pstr returns fully qualified package string
func fqp(s string) string {
	i := strings.LastIndex(s, "/")
	if i == -1 {
		i = 1 // builtin, eg. runtime
	}
	j := strings.Index(s[i:], ".")
	if j == -1 {
		return ""
	}
	return s[:i+j]
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
