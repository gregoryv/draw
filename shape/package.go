package shape

import (
	"io"
)

type Stringer interface {
	String() string
}

type svg interface {
	WriteSvg(io.Writer) error
}
