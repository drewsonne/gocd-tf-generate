// Adapted from `stringer`:
// - https://blog.golang.org/generate
// - http://godoc.org/golang.org/x/tools/cmd/stringer

package main

import (
	"go/importer"
	"go/types"
)

func defaultImporter() types.Importer {
	return importer.Default()
}
