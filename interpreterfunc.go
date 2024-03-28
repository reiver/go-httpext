package httpext

import (
	"io/fs"
)

func InterpreterFunc(extension string, fn func(fs.File)fs.File) Interpreter {
	if nil == fn {
		return nil
	}
	if "" == extension {
		return nil
	}

	return internalInterpreterFunc{
		extension:extension,
		fn:fn,
	}
}

type internalInterpreterFunc struct {
	extension string
	fn func(fs.File)fs.File
}

func (receiver internalInterpreterFunc) InterpreterExtension() string {
	return receiver.extension
}

func (receiver internalInterpreterFunc) InterpretFile(file fs.File) fs.File {
	var fn func(fs.File)fs.File = receiver.fn
	if nil == fn {
		return nil
	}

	return fn(file)
}
