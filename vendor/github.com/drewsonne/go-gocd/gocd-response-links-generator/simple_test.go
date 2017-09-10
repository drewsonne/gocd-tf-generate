// Adapted from `stringer`:
// - https://blog.golang.org/generate
// - http://godoc.org/golang.org/x/tools/cmd/stringer

package main

import (
	"strings"
	"testing"
	//"fmt"
	"go/format"
)

const testInput = `type BuildDetails_ResponseLinks struct {
	Doc    *url.URL
	Find   *url.URL
}
`

const testOutput = `package test
import "encoding/json"
func (l BuildDetails_ResponseLinks) MarshalJSON() ([]byte, error) {
	type h struct {
		h string ` + "`json:\"href\"`" + `
	}
	ls := struct {
		Doc  *h ` + "`json:\"doc,omitempty\"`" + `
		Find *h ` + "`json:\"find,omitempty\"`" + `
	}{}
	if l.Doc != nil {
		ls.Doc = &h{h: l.Doc.String()}
	}
	if l.Find != nil {
		ls.Find = &h{h: l.Find.String()}
	}
	j, e := json.Marshal(ls)
	if e != nil {
		return nil, e
	}
	return j, nil
}

func (l *BuildDetails_ResponseLinks) UnmarshalJSON(j []byte) error {
	var d map[string]map[string]string
	e := json.Unmarshal(j, &d)
	if e != nil {
		return e
	}

	if d["doc"]["href"] != "" {
		l.Doc, e = url.Parse(d["doc"]["href"])
		if e != nil {
			return e
		}
	}
	if d["find"]["href"] != "" {
		l.Find, e = url.Parse(d["find"]["href"])
		if e != nil {
			return e
		}
	}
	return nil
}
`

func TestGenerator(t *testing.T) {
	var g Generator
	name := "test"
	file := name + ".go"
	input := "package " + name + "\nimport \"net/url\"\n" + testInput
	g.parsePackage(".", []string{file}, input)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(testInput, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", name)
	}
	g.generate(tokens[1])
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		t.Error(err)
	}

	got := "package " + name + "\nimport \"encoding/json\"\n" + string(src)

	if got != testOutput {
		t.Errorf("%s: got\n====\n%s\n====\nexpected\n====\n%s\n====", name, got, testOutput)
	}
}
