package httpextfs

import (
	"github.com/reiver/go-path"
)

const  defaultWebPageFileName string = "webpage.html"

func fsName(name string, defaultFileName string) string {
	name = path.Canonical("/" + name)

	if "" == defaultFileName {
		defaultFileName = defaultWebPageFileName
	}

	{
		var last int = len(name) - 1

		if '/' == name[last] {
			name  = path.Join(name, defaultFileName)
		}

		if '/' == name[0] {
			name = name[1:]
		}
	}

	{
		var fileextension string = path.Ext(name)
		if "" == fileextension {
			name += ".html"
		}
	}

	return name
}
