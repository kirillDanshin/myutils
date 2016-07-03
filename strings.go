package myutils

import "github.com/valyala/bytebufferpool"

// Concat two strings
func Concat(a string, b ...string) string {
	if len(b) == 0 {
		return a
	}
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	buf.SetString(a)
	for _, s := range b {
		buf.WriteString(s)
	}

	return string(buf.B[:])
}
