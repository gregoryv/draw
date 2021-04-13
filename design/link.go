package design

import "fmt"

func (d *SequenceDiagram) Link(from, to, text string) *Link {
	fromIndex := -1
	toIndex := -1
	for i, column := range d.columns {
		if column == from {
			fromIndex = i
			break
		}
	}
	for i, column := range d.columns {
		if column == to {
			toIndex = i
			break
		}
	}
	lnk := &Link{
		fromIndex: fromIndex,
		toIndex:   toIndex,
		text:      text,
	}
	d.links = append(d.links, lnk)
	if fromIndex == -1 {
		panic(fmt.Sprintf("Missing %q column", from))
	}
	if toIndex == -1 {
		panic(fmt.Sprintf("Missing %q column", to))
	}
	return lnk
}

// Skip adds dotted distance on all columns
func (d *SequenceDiagram) Skip() {
	d.links = append(d.links, skip)
}

func (d *SequenceDiagram) ClearLinks() {
	d.links = make([]*Link, 0)
}

var skip *Link = &Link{}

// Link represents an arrow in a sequence diagram
type Link struct {
	fromIndex, toIndex int
	text               string
	Class              string
	TextClass          string
}

func (l *Link) toSelf() bool {
	return l.fromIndex == l.toIndex
}

func (l *Link) class() string {
	if l.Class == "" {
		return "arrow"
	}
	return l.Class
}
