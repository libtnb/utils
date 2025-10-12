package debug

import (
	"io"

	"github.com/goforj/godump"
)

// Dump is used to display detailed information about variables
// And this is a wrapper around godump.Dump.
func Dump(v ...any) {
	godump.Dump(v...)
}

// FDump is used to display detailed information about variables to the specified io.Writer
// And this is a wrapper around godump.Fdump.
func FDump(w io.Writer, v ...any) {
	godump.Fdump(w, v...)
}

// SDump is used to display detailed information about variables as a string,
// And this is a wrapper around godump.DumpStr.
func SDump(v ...any) string {
	return godump.DumpStr(v...)
}
