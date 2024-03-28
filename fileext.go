package httpext

import (
	"errors"
	"io/fs"

	"github.com/reiver/go-path"
)

var (
	errNilFileInfo = errors.New("httpext: nil file-info")
)

func fileExt(file fs.File) (string, error) {

	var filename string
	{
		var fileinfo fs.FileInfo
		var err error

		fileinfo, err = file.Stat()
		if nil != err {
			return "", err
		}
		if nil == fileinfo {
			return "", errNilFileInfo
		}

		filename = fileinfo.Name()
	}

	return path.Ext(filename), nil
}
