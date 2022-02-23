package stringy

import (
	"fmt"
	"io"
)

func Stringify[T fmt.Stringer](w io.Writer, p T) {
	fmt.Fprintln(w, p.String())
}
