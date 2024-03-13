package httpext

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"net/url"

	"github.com/reiver/go-ext2mime"
	"github.com/reiver/go-path"

	"github.com/reiver/go-httpext/fs"
)

const (
	defaultContentType = "application/octet-stream"
	defaultWebPageFileName = "webpage.html"
)

var statusTextNotFound            string = http.StatusText(http.StatusNotFound)
var statusTextInternalServerError string = http.StatusText(http.StatusInternalServerError)

func HTTPHandler(filesystem fs.FS) http.Handler {
	if nil == filesystem {
		return nil
	}

	filesystem = httpextfs.FS(filesystem, defaultWebPageFileName)

	return internalHTTPHandler{
		filesystem:filesystem,
	}
}

type internalHTTPHandler struct {
	filesystem fs.FS
}

func (receiver internalHTTPHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
		return
	}

	var filesystem fs.FS = receiver.filesystem
	if nil == filesystem {
		http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
		return
	}

	var httpRequestPath string
	{
		var httpRequestURL *url.URL = request.URL
		if nil == httpRequestURL {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		httpRequestPath = httpRequestURL.Path
	}

	var file fs.File
	{
		var err error

		file, err = filesystem.Open(httpRequestPath)
		switch {
		case errors.Is(err, fs.ErrNotExist):
			http.Error(responseWriter, statusTextNotFound, http.StatusNotFound)
			return
		default:
			if nil != err {
				http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
				return
			}
		}
		defer file.Close()
	}

	var filename string
	{
		var fileinfo fs.FileInfo
		var err error

		fileinfo, err = file.Stat()
		if nil != err {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}
		if nil == fileinfo {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		filename = fileinfo.Name()
	}

	var fileextension string = path.Ext(filename)

	{
		var contenttype string
		var found bool

		contenttype, found = ext2mime.Get(fileextension)
		if !found {
			contenttype = defaultContentType
		}

		var header http.Header = responseWriter.Header()
		if nil == header {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		header.Set("Content-Type", contenttype)
	}

	io.Copy(responseWriter, file)
}
