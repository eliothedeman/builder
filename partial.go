package builder

import (
	"bytes"
	"io"
)

type partial []byte

func (p partial) Build(w io.Writer) {
	w.Write([]byte(p))
}

// Raw string
func Raw(s string) Builder {
	return partial(s)
}

// Compile joins multiple builders into a single static builder
func Compile(b ...Builder) Builder {
	var buff bytes.Buffer

	for _, x := range b {
		x.Build(&buff)
	}

	return partial(buff.Bytes())
}
