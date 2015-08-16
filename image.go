// Package image check if given path is image
package image // import "github.com/shamsher31/goisimg"

import (
	"github.com/shamsher31/goimgext"
	"path"
	"strings"
)

// Call Get from imgext package
var extensions = imgext.Get()

// Extensions is the extensions for different image types
var Extensions map[string]bool

func init() {
	Extensions = map[string]bool{}
	for _, ext := range extensions {
		Extensions[ext] = true
	}
}

// Is checks if a path is a image
func Is(p string) bool {
	ext := strings.TrimLeft(path.Ext(p), ".")
	return Extensions[strings.ToLower(ext)]
}
