package httpextfs

import (
	"errors"
	"io/fs"
)

var (
	errDefaultsTooShort = errors.New("httpext: defaults too short")
	errNilFileSystem    = errors.New("httpext: nil file-system")
)

// FS returns a fs.FS that wraps 'filesystem' making it so that it can handle paths from an HTTP request.
//
// If 'defaults' is "index.html", then behind the scenes, it does these conversions:
//
//	"/" -> "index.html"
//
//	"/apple/banana/cherry/" -> "apple/banana/cherry/index.html"
//
//	"/item" -> "item.html"
//
//	"/dir/item" -> "dir/item.html"
//
//	"/img/photo.jpeg" -> "img/photo.jpeg"
//
// If 'defaults' is "img.png", then behind the scenes, it does these conversions:
//
//	"/" -> "img.png"
//
//	"/apple/banana/cherry/" -> "apple/banana/cherry/img.png"
//
//	"/item" -> "item.html"
//
//	"/dir/item" -> "dir/item.html"
//
//	"/img/photo.jpeg" -> "img/photo.jpeg"
//
// If 'defaultFileName' is the empty string (i.e., "") then it defaults to "webpage.html".
func FS(filesystem fs.FS, defaults ...string) fs.FS {
	if len(defaults) < 2 {
		return nil
	}

	return internalFS{
		filesystem:filesystem,
		defaults:defaults,
	}
}

type internalFS struct {
	filesystem fs.FS
	defaults []string
}

func (receiver internalFS) Open(name string) (fs.File, error) {
	var filesystem fs.FS = receiver.filesystem
	if nil == filesystem {
		return nil, errNilFileSystem
	}

	var defaults []string = receiver.defaults
	if len(defaults) < 2 {
		return nil, errDefaultsTooShort
	}

	{
		var defaultStem string  = defaults[0]

		for _, defaultExtension := range defaults[1:]  {
			var fsname string = fsName(name, defaultStem, defaultExtension)
			file, err := filesystem.Open(fsname)
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return file, err
		}
		return nil, fs.ErrNotExist
	}
}
