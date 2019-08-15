package shape

import "io"

type StyledWriter interface {
	io.WriterTo
	Style() string
}
