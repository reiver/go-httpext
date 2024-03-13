# go-httpext

Package **httpext** provides an `http.Handler` file-server that makes file-extensions for HTML files optional, for the Go programming language.

For example, instead of "/apple/banana/cherry.html" you could use "/apple/banana/cherry".

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-httpext

[![GoDoc](https://godoc.org/github.com/reiver/go-httpext?status.svg)](https://godoc.org/github.com/reiver/go-httpext)

## Example

Here is a simple example of parsing a hex color code:

```golang
var handler http.Handler = httpext.HTTPHandler(filesystem)

// ...

err := http.ListenAndServe(addr, handler)
```

## Import

To import package **httpext** use `import` code like the follownig:
```
import "github.com/reiver/go-httpext"
```

## Installation

To install package **httpext** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-httpext
```

## Author

Package **httpext** was written by [Charles Iliya Krempeaux](http://reiver.link)
