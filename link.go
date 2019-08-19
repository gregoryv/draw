package design

import "fmt"

func (dia *SequenceDiagram) Link(from, to, text string) *Link {
	fromIndex := -1
	toIndex := -1
	for i, column := range dia.columns {
		if column == from {
			fromIndex = i
			break
		}
	}
	for i, column := range dia.columns {
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
	dia.links = append(dia.links, lnk)
	if fromIndex == -1 {
		panic(fmt.Sprintf("Missing %q column", from))
	}
	if toIndex == -1 {
		panic(fmt.Sprintf("Missing %q column", to))
	}
	return lnk
}

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
