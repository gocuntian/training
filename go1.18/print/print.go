package print

import (
	"fmt"
	"io"
)

// PrintAnything ...
func PrintAnything[T any](w io.Writer, p T) {
	fmt.Fprintln(w, p)
}
