package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/open2b/scriggo/templates"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {

	start := time.Now()

	dstDir, err := os.MkdirTemp(".", "public-temp-*")
	if err != nil {
		log.Fatal(err)
	}
	dstBase := filepath.Base(dstDir)
	defer func() {
		err = os.RemoveAll(dstDir)
		if err != nil {
			log.Fatal(err)
		}
	}()

	md := goldmark.New(
		goldmark.WithRendererOptions(html.WithUnsafe()),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()))

	buildOptions := &templates.BuildOptions{
		Globals: globals,
		MarkdownConverter: func(src []byte, out io.Writer) error {
			return md.Convert(src, out)
		},
	}

	srcFS := os.DirFS("./")

	err = fs.WalkDir(srcFS, ".", func(path string, d fs.DirEntry, err error) error {
		if path == "public" || path == dstBase || strings.HasPrefix(path, ".git/") || strings.HasPrefix(path, "cmd/") {
			return fs.SkipDir
		}
		if d.IsDir() || path[0] == '.' || path == "README.md" {
			return nil
		}
		ext := filepath.Ext(path)
		switch ext {
		case ".md":
			template, err := templates.Build(srcFS, path, buildOptions)
			if err != nil {
				log.Fatal(err)
			}
			name := filepath.Join(dstDir, path[:len(path)-3]) + ".html"
			err = os.MkdirAll(filepath.Dir(name), 0500)
			if err != nil {
				log.Fatal(err)
			}
			fi, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				log.Fatal(err)
			}
			err = template.Run(fi, nil, nil)
			if err == nil {
				err = fi.Close()
			}
			if err != nil {
				log.Fatal(err)
			}
		case ".css", ".js", ".png", ".ico", ".txt":
			src, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			name := filepath.Join(dstDir, path)
			err = os.MkdirAll(filepath.Dir(name), 0700)
			if err != nil {
				log.Fatal(err)
			}
			dst, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(dst, src)
			if err == nil {
				_ = src.Close()
				err = dst.Close()
			}
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("public/")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(dstDir, "public")
	if err != nil {
		log.Fatal(err)
	}
	buildTime := time.Since(start)
	_, _ = fmt.Fprintf(os.Stderr, "Public site generated in %s\n", buildTime)
	return
}

var globals = templates.Declarations{}
