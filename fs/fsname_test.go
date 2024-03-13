package httpextfs

import (
	"testing"
)

func TestFSName(t *testing.T) {

	tests := []struct{
		HTTPPath string
		DefaultFileName string
		Expected string
	}{
		{
			HTTPPath: "",
			DefaultFileName: "",
			Expected: "webpage.html",
		},
		{
			HTTPPath: "/",
			DefaultFileName: "",
			Expected: "webpage.html",
		},
		{
			HTTPPath: ".",
			DefaultFileName: "",
			Expected: "webpage.html",
		},
		{
			HTTPPath: "./",
			DefaultFileName: "",
			Expected: "webpage.html",
		},
		{
			HTTPPath: "../",
			DefaultFileName: "",
			Expected: "webpage.html",
		},
		{
			HTTPPath: "../../",
			DefaultFileName: "",
			Expected: "webpage.html",
		},



		{
			HTTPPath: "robots.txt",
			DefaultFileName: "",
			Expected: "robots.txt",
		},
		{
			HTTPPath: "/robots.txt",
			DefaultFileName: "",
			Expected: "robots.txt",
		},
		{
			HTTPPath: "./robots.txt",
			DefaultFileName: "",
			Expected: "robots.txt",
		},
		{
			HTTPPath: "../robots.txt",
			DefaultFileName: "",
			Expected: "robots.txt",
		},
		{
			HTTPPath: "../../robots.txt",
			DefaultFileName: "",
			Expected: "robots.txt",
		},



		{
			HTTPPath: "2024/03/13/post",
			DefaultFileName: "",
			Expected: "2024/03/13/post.html",
		},
		{
			HTTPPath: "/2024/03/13/post",
			DefaultFileName: "",
			Expected: "2024/03/13/post.html",
		},
		{
			HTTPPath: "./2024/03/13/post",
			DefaultFileName: "",
			Expected: "2024/03/13/post.html",
		},
		{
			HTTPPath: "../2024/03/13/post",
			DefaultFileName: "",
			Expected: "2024/03/13/post.html",
		},
		{
			HTTPPath: "../../2024/03/13/post",
			DefaultFileName: "",
			Expected: "2024/03/13/post.html",
		},



		{
			HTTPPath: "2024/03/13/photo.jpeg",
			DefaultFileName: "",
			Expected: "2024/03/13/photo.jpeg",
		},
		{
			HTTPPath: "/2024/03/13/photo.jpeg",
			DefaultFileName: "",
			Expected: "2024/03/13/photo.jpeg",
		},
		{
			HTTPPath: "./2024/03/13/photo.jpeg",
			DefaultFileName: "",
			Expected: "2024/03/13/photo.jpeg",
		},
		{
			HTTPPath: "../2024/03/13/photo.jpeg",
			DefaultFileName: "",
			Expected: "2024/03/13/photo.jpeg",
		},
		{
			HTTPPath: "../../2024/03/13/photo.jpeg",
			DefaultFileName: "",
			Expected: "2024/03/13/photo.jpeg",
		},



		{
			HTTPPath: "2024/03/13/",
			DefaultFileName: "",
			Expected: "2024/03/13/webpage.html",
		},
		{
			HTTPPath: "/2024/03/13/",
			DefaultFileName: "",
			Expected: "2024/03/13/webpage.html",
		},
		{
			HTTPPath: "./2024/03/13/",
			DefaultFileName: "",
			Expected: "2024/03/13/webpage.html",
		},
		{
			HTTPPath: "../2024/03/13/",
			DefaultFileName: "",
			Expected: "2024/03/13/webpage.html",
		},
		{
			HTTPPath: "../../2024/03/13/",
			DefaultFileName: "",
			Expected: "2024/03/13/webpage.html",
		},



		{
			HTTPPath: "2024/03/13/",
			DefaultFileName: "index.html",
			Expected: "2024/03/13/index.html",
		},
		{
			HTTPPath: "/2024/03/13/",
			DefaultFileName: "index.html",
			Expected: "2024/03/13/index.html",
		},
		{
			HTTPPath: "./2024/03/13/",
			DefaultFileName: "index.html",
			Expected: "2024/03/13/index.html",
		},
		{
			HTTPPath: "../2024/03/13/",
			DefaultFileName: "index.html",
			Expected: "2024/03/13/index.html",
		},
		{
			HTTPPath: "../../2024/03/13/",
			DefaultFileName: "index.html",
			Expected: "2024/03/13/index.html",
		},



		{
			HTTPPath: "2024/03/13/",
			DefaultFileName: "meme.gif",
			Expected: "2024/03/13/meme.gif",
		},
		{
			HTTPPath: "/2024/03/13/",
			DefaultFileName: "meme.gif",
			Expected: "2024/03/13/meme.gif",
		},
		{
			HTTPPath: "./2024/03/13/",
			DefaultFileName: "meme.gif",
			Expected: "2024/03/13/meme.gif",
		},
		{
			HTTPPath: "../2024/03/13/",
			DefaultFileName: "meme.gif",
			Expected: "2024/03/13/meme.gif",
		},
		{
			HTTPPath: "../../2024/03/13/",
			DefaultFileName: "meme.gif",
			Expected: "2024/03/13/meme.gif",
		},
	}

	for testNumber, test := range tests {

		actual := fsName(test.HTTPPath, test.DefaultFileName)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'file-system name' is not what was expected", testNumber)
			t.Logf("EXPECTED (FS-PATH): %q", expected)
			t.Logf("ACTUAL   (FS-PATH): %q", actual)
			t.Logf("HTTP-PATH: %q", test.HTTPPath)
			t.Logf("DEFAULT-FILE-NAME: %q", test.DefaultFileName)
			continue
		}
	}
}
