package builder

import (
	"bytes"
	"fmt"
	"io"
)

// Class is a CSS class
type Class struct {
	name  string
	attrs map[string]string
}

func (c *Class) Style(t TagBuilder) Builder {
	t.Tag("class", c.name)
	return t
}

func CSSClass(name string, attrs map[string]string) *Class {
	return &Class{
		name:  name,
		attrs: attrs,
	}
}

// Style is a group of CSS
func Style(classes ...*Class) Builder {
	return &node{
		name: style,
		body: Join(func() (out []Builder) {
			for _, c := range classes {
				out = append(out, c)
			}
			return
		}()...),
	}
}

func (c *Class) Build(w io.Writer) {
	var b bytes.Buffer
	for k, v := range c.attrs {
		fmt.Fprintf(&b, "%s:%s;", k, v)
	}
	fmt.Fprintf(w, `.%s {%s}`, c.name, b.String())
}
