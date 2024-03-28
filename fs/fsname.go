package httpextfs

import (
	"github.com/reiver/go-path"
)

const  defaultWebPageStem      string = "webpage"
const  defaultWebPageExtension string = ".html"

func fsName(name string, defaultStem string, defaultExtension string) string {
	name = path.Canonical("/" + name)

	if "" == defaultStem {
		defaultStem = defaultWebPageStem
	}
	if "" == defaultExtension {
		defaultExtension = defaultWebPageExtension
	}

	var defaultFileName string = defaultStem + defaultExtension

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
			name += defaultExtension
		}
	}

	return name
}
