package httpext

import (
	"io/fs"
)

type Interpreter interface {
	InterpreterExtension() string
	InterpretFile(fs.File) fs.File
}
