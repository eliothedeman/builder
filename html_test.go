package builder

import (
	"bytes"
	"testing"
)

func toString(b Builder) string {
	var x bytes.Buffer
	b.Build(&x)
	return x.String()
}

func TestBasic(t *testing.T) {

	bb := Body(
		Raw("this is a test"),
	)

	if toString(bb) != "<body>this is a test</body>" {
		t.Fatal(toString(bb))
	}

	bb = Body(
		P(
			Raw("whats up?"),
		).Tag("class", "test").Tag("xxx", "xxxx"),
	)
	if toString(bb) != `<body><p class="test" xxx="xxxx">whats up?</p></body>` {
		t.Fatal(toString(bb))
	}

}

func BenchmarkBuildVsCompile(b *testing.B) {
	bb := Body(
		Raw(";laksdjflkj"),
		P(
			Raw("alksjdf;lkjadsf"),
			H1(Raw(";lkajsdflkajdf")),
		),
	)
	b.Run("dynamic", func(b *testing.B) {
		buff := bytes.NewBuffer(nil)
		for i := 0; i < b.N; i++ {
			bb.Build(buff)
			b.SetBytes(int64(buff.Len()))
			buff.Reset()
		}
	})
	b.Run("compiled", func(b *testing.B) {
		buff := bytes.NewBuffer(nil)
		c := Compile(bb)
		for i := 0; i < b.N; i++ {
			c.Build(buff)
			b.SetBytes(int64(buff.Len()))
			buff.Reset()
		}
	})
}
