package builder

import (
	"fmt"
	"io"
)

const (
	a      = "a"
	p      = "p"
	b      = "b"
	h1     = "h1"
	h2     = "h2"
	h3     = "h3"
	h4     = "h4"
	body   = "body"
	title  = "title"
	input  = "input"
	form   = "form"
	style  = "style"
	script = "script"
)

type funcBuilder func(w io.Writer)

func (f funcBuilder) Build(w io.Writer) {
	f(w)
}

// Each will call f until it returns nil
func Each(f func() Builder) Builder {
	return funcBuilder(func(w io.Writer) {
		next := f()
		for next != nil {
			next.Build(w)
			next = f()
		}
	})
}

// Fmt takes a format string
func Fmt(s string, args ...interface{}) Builder {
	return Raw(fmt.Sprintf(s, args...))
}

// A is a link
func A(s ...Builder) TagBuilder {
	return &node{
		name: a,
		body: Join(s...),
	}
}

// Form input
func Form(action string, s ...Builder) TagBuilder {
	n := &node{
		name: form,
		body: Join(s...),
	}
	return n.Tag("action", action)
}

// Input in a form
func Input(inputType, name string) TagBuilder {
	n := &node{
		name: input,
	}

	return n.Tag("type", inputType).Tag("name", name)
}

// P is a paragraph
func P(s ...Builder) TagBuilder {
	return &node{
		name: p,
		body: Join(s...),
	}
}

// B is a bold
func B(s ...Builder) TagBuilder {
	return &node{
		name: b,
		body: Join(s...),
	}
}

// H1 is a h1 heading
func H1(s ...Builder) TagBuilder {
	return &node{
		name: h1,
		body: Join(s...),
	}
}

// H2 is a h2 heading
func H2(s ...Builder) TagBuilder {
	return &node{
		name: h2,
		body: Join(s...),
	}
}

// H3 is a h3 heading
func H3(s ...Builder) TagBuilder {
	return &node{
		name: h3,
		body: Join(s...),
	}
}

// H4 is a h4 heading
func H4(s ...Builder) TagBuilder {
	return &node{
		name: h4,
		body: Join(s...),
	}
}

// Body is a body
func Body(s ...Builder) TagBuilder {
	return &node{
		name: body,
		body: Join(s...),
	}
}

func Title(s ...Builder) TagBuilder {
	return &node{
		name: title,
		body: Join(s...),
	}
}

type pair struct {
	k string
	v string
}

func (p *pair) Build(w io.Writer) {
	if p.v == "" {
		fmt.Fprintf(w, ` %s`, p.k)
		return
	}
	fmt.Fprintf(w, ` %s="%s"`, p.k, p.v)
}

type joiner []Builder

func (j joiner) Build(w io.Writer) {
	for _, b := range j {
		b.Build(w)
	}
}

// Join multiple builders into a single builder
func Join(b ...Builder) Builder {
	j := joiner(b)
	return j
}

type node struct {
	name  string
	body  Builder
	pairs []pair
}

func (n *node) Tag(k, v string) TagBuilder {
	n.pairs = append(n.pairs, pair{k, v})
	return n
}

var (
	startTag    = []byte{'<'}
	endTag      = []byte{'>'}
	startEndTag = []byte{'<', '/'}
)

func (n *node) Build(w io.Writer) {
	tmp := sToB(n.name)

	w.Write(startTag)
	w.Write(tmp)
	for _, p := range n.pairs {
		p.Build(w)
	}
	w.Write(endTag)
	if n.body == nil {
		return
	}

	if n.body != nil {
		n.body.Build(w)
	}
	w.Write(startEndTag)
	w.Write(tmp)
	w.Write(endTag)
}
