package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets90ba780e7ce0fde2c751eedf5fb441ba62ac6407 = "<!doctype html>\r\n<body>\r\n    <p>Hello, {{.Foo}}</p>\r\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"html"}, "/templates/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1623404589, 1623404589285154400),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1623404123, 1623404123785871600),
		Data:     nil,
	}, "/templates/html": &assets.File{
		Path:     "/templates/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1623404642, 1623404642268160400),
		Data:     nil,
	}, "/templates/html/index.tmpl": &assets.File{
		Path:     "/templates/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1623404642, 1623404642267999800),
		Data:     []byte(_Assets90ba780e7ce0fde2c751eedf5fb441ba62ac6407),
	}}, "")
