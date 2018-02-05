package builder

import "io"

// Builder provides
type Builder interface {
	Build(w io.Writer)
}

// TagBuilder provides
type TagBuilder interface {
	Builder
	Tag(k, v string) TagBuilder
}
