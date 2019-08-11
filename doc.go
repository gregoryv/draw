package design

import (
	"io"
	"os"
	"text/template"
)

type DesignDoc struct {
	Parts Stringers
}

func NewDesignDoc() *DesignDoc {
	return &DesignDoc{}
}

func (doc *DesignDoc) Editor() Editor {
	return doc.edit
}

type Editor func(...interface{})

func (doc *DesignDoc) edit(arguments ...interface{}) {
	for _, arg := range arguments {
		doc.appendByType(arg)
	}
}

func (doc *DesignDoc) appendByType(arg interface{}) {
	var valid Stringer
	switch arg := arg.(type) {
	case string:
		valid = plain(arg)
	case *Graph:
		doc.Parts = append(doc.Parts, plain("\n"))
		valid = arg
	case Stringer:
		valid = arg
	default:
		panic(arg)
	}
	doc.Parts = append(doc.Parts, valid)
}

func (doc *DesignDoc) SaveAs(filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = doc.WriteTo(fh)
	return err
}

func (doc *DesignDoc) WriteTo(w io.Writer) (int, error) {
	r, writer := io.Pipe()
	go func() {
		tpl := template.Must(template.New("html").Parse(htmlSource))
		tpl.Execute(writer, doc)
		writer.Close()
	}()
	n, err := io.Copy(w, r)
	// Safe for most cases, we're dealing with small documents
	return int(n), err
}

const htmlSource = `<!DOCTYPE html>

<html>
  <head>
    <style>
      .component, .smallbox {
        fill:#ffffcc;
        stroke:black;
        stroke-width:1;
      }
      line {
        stroke:black;
        stroke-width:1;
      }

    </style>
  </head>
<body>
{{range .Parts}}{{.}}{{end}}
</body>
</html>`

type Stringers []Stringer

type Stringer interface {
	String() string
}

type StringerFunc func() string

func (fn StringerFunc) String() string {
	return fn()
}

type plain string

func (p plain) String() string {
	return string(p)
}
