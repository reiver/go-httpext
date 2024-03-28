package httpext

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"net/url"

	"github.com/reiver/go-ext2mime"

	"github.com/reiver/go-httpext/fs"
)

const (
	defaultContentType = "application/octet-stream"
	defaultWebPageStem = "webpage"
	defaultWebPageExtension = ".html"
)

var statusTextNotFound            string = http.StatusText(http.StatusNotFound)
var statusTextInternalServerError string = http.StatusText(http.StatusInternalServerError)

func HTTPHandler(filesystem fs.FS, interpreters ...Interpreter) http.Handler {
	if nil == filesystem {
		return nil
	}

	var extensions []string = []string{defaultWebPageExtension}

	var interpretersMap map[string]Interpreter = make(map[string]Interpreter)
	for _, interpreter := range interpreters {
		if nil == interpreter {
			continue
		}

		var ext string = interpreter.InterpreterExtension()
		if "" == ext {
			continue
		}

		interpretersMap[ext] = interpreter

		extensions = append(extensions, ext)
	}

	filesystem = httpextfs.FS(filesystem, defaultWebPageStem, extensions...)

	return internalHTTPHandler{
		filesystem:filesystem,
		interpreters:interpretersMap,
	}
}

type internalHTTPHandler struct {
	filesystem fs.FS
	interpreters map[string]Interpreter
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

	var fileextension string
	{
		var err error

		fileextension, err = fileExt(file)
		if nil != err {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}
	}

	{
		var interpreters map[string]Interpreter = receiver.interpreters
		if nil != interpreters {

			var interpreter Interpreter
			var found bool

			interpreter, found = interpreters[fileextension]
			if found {
				var newfile fs.File = interpreter.InterpretFile(file)
				if nil == newfile {
					http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
					return
				}

				file = newfile
			}
		}

		{
			var err error

			fileextension, err = fileExt(file)
			if nil != err {
				http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
				return
			}
		}
	}

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
