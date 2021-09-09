package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/native"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {

	start := time.Now()

	err := os.Mkdir("public", 0755)
	if err != nil {
		log.Fatal(err)
	}

	md := goldmark.New(
		goldmark.WithRendererOptions(html.WithUnsafe()),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()))

	buildOptions := &scriggo.BuildOptions{
		Globals: globals,
		MarkdownConverter: func(src []byte, out io.Writer) error {
			return md.Convert(src, out)
		},
	}

	srcFS := os.DirFS("site")

	err = fs.WalkDir(srcFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path[0] == '.' {
			return nil
		}
		if d.IsDir() {
			return os.MkdirAll(filepath.Join("public", path), 0700)
		}
		ext := filepath.Ext(path)
		switch ext {
		case ".html":
		case ".md":
			template, err := scriggo.BuildTemplate(srcFS, path, buildOptions)
			if err != nil {
				return err
			}
			name := filepath.Join("public", path[:len(path)-3]) + ".html"
			fi, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				return err
			}
			err = template.Run(fi, nil, nil)
			if err == nil {
				err = fi.Close()
			}
		default:
			src, err := srcFS.Open(path)
			if err != nil {
				return err
			}
			name := filepath.Join("public", path)
			err = os.MkdirAll(filepath.Dir(name), 0700)
			if err != nil {
				return err
			}
			dst, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				return err
			}
			_, err = io.Copy(dst, src)
			if err == nil {
				_ = src.Close()
				err = dst.Close()
			}
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	buildTime := time.Since(start)
	_, _ = fmt.Fprintf(os.Stderr, "Public site generated in %s\n", buildTime)
	return
}

var globals = native.Declarations{}
