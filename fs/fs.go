package httpextfs

import (
	"errors"
	"io/fs"
)

var (
	errNilFileSystem = errors.New("httpext: nil file-system")
)

// FS returns a fs.FS that wraps 'filesystem' making it so that it can handle paths from an HTTP request.
//
// If 'defaultFileName' is "index.html", then behind the scenes, it does these conversions:
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
// If 'defaultFileName' is "img.png", then behind the scenes, it does these conversions:
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
func FS(filesystem fs.FS, defaultFileName string) fs.FS {
	return internalFS{
		filesystem:filesystem,
		defaultFileName:defaultFileName,
	}
}

type internalFS struct {
	filesystem fs.FS
	defaultFileName string
}

func (receiver internalFS) Open(name string) (fs.File, error) {
	var filesystem fs.FS = receiver.filesystem
	if nil == filesystem {
		return nil, errNilFileSystem
	}

	var defaultFileName string = receiver.defaultFileName

	var fsname string = fsName(name, defaultFileName)

	return filesystem.Open(fsname)
}
