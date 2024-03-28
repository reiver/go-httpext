package httpextfs

import (
	"errors"
	"io/fs"

	"github.com/reiver/go-path"
)

var (
	errDefaultExtensionsTooShort = errors.New("httpext: default extensions too short")
	errNilFileSystem = errors.New("httpext: nil file-system")
)

// FS returns an fs.FS that wraps 'filesystem' making it so that it can handle paths from an HTTP request.
//
// If 'defaultStem' is "index" and 'defaultExtensions' is []string{".html"}, then behind the scenes, it does these conversions:
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
// If 'defaultStem' is "img" and 'defaultExtensions' is []string{".png"}, then behind the scenes, it does these conversions:
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
func FS(filesystem fs.FS, defaultStem string, defaultExtensions ...string) fs.FS {
	if len(defaultExtensions) < 1 {
		return nil
	}

	return internalFS{
		filesystem:filesystem,
		defaultStem:defaultStem,
		defaultExtensions:defaultExtensions,
	}
}

type internalFS struct {
	filesystem fs.FS
	defaultStem string
	defaultExtensions []string
}

func (receiver internalFS) Open(name string) (fs.File, error) {
	var filesystem fs.FS = receiver.filesystem
	if nil == filesystem {
		return nil, errNilFileSystem
	}

	if "" != path.Ext(name) {
		var fsname string = fsName(name, "", "")
		return filesystem.Open(fsname)
	}

	var defaultExtensions []string = receiver.defaultExtensions
	if len(defaultExtensions) < 1 {
		return nil, errDefaultExtensionsTooShort
	}

	{
		var defaultStem string  = receiver.defaultStem

		for _, defaultExtension := range defaultExtensions  {
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
