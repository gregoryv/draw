package design

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/shape"
)

func BenchmarkClassDiagram_WriteSVG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := exampleClassDiagram()
		d.WriteSVG(ioutil.Discard)
	}
}

func exampleClassDiagram() *ClassDiagram {
	var (
		d        = NewClassDiagram()
		record   = d.Struct(shape.Record{})
		arrow    = d.Struct(shape.Arrow{})
		line     = d.Struct(shape.Line{})
		circle   = d.Struct(shape.Circle{})
		diaarrow = d.Struct(shape.Diamond{})
		triangle = d.Struct(shape.Triangle{})
		shapE    = d.Interface((*shape.Shape)(nil))
	)
	d.HideRealizations()

	var (
		fnt      = d.Struct(shape.Font{})
		style    = d.Struct(draw.Style{})
		seqdia   = d.Struct(SequenceDiagram{})
		classdia = d.Struct(ClassDiagram{})
		dia      = d.Struct(Diagram{})
		aligner  = d.Struct(shape.Aligner{})
		adj      = d.Struct(shape.Adjuster{})
		rel      = d.Struct(Relation{})
	)
	d.HideRealizations()

	d.Place(shapE).At(220, 20)
	d.Place(record).At(20, 120)
	d.Place(line).Below(shapE, 90)
	d.VAlignCenter(shapE, line)

	d.Place(arrow).RightOf(line, 90)
	d.Place(circle).RightOf(shapE, 280)
	d.Place(diaarrow).Below(circle)
	d.Place(triangle).Below(diaarrow)
	d.HAlignBottom(record, arrow, line)
	shape.Move(line, 30, 30)

	d.Place(fnt).Below(record, 170)
	d.Place(style).RightOf(fnt, 90)
	d.VAlignCenter(shapE, line, style)
	d.VAlignCenter(record, fnt)

	d.Place(rel).Below(line, 80)
	d.Place(dia).RightOf(style, 90)
	d.Place(aligner).RightOf(dia, 80)
	d.HAlignCenter(fnt, style, dia, aligner)

	d.Place(adj).Below(fnt, 70)
	d.Place(seqdia).Below(aligner, 90)
	d.Place(classdia).Below(dia, 90)
	d.VAlignCenter(dia, classdia)
	d.HAlignBottom(classdia, seqdia)

	d.SetCaption("Figure 1. Class diagram of design and design.shape packages")
	return d
}

func TestClassDiagram_OtherType(t *testing.T) {
	newClassDiagram().SaveAs("testdata/classdia_test.svg")
	data := newClassDiagram().String()
	ioutil.WriteFile("testdata/classdia_inlined.svg", []byte(data), 0644)
}

func TestClassDiagram_Inline(t *testing.T) {
	got := newClassDiagram().Inline()
	if strings.Contains(got, "class") {
		t.Error("found class attributes\n", got)
	}
}

func TestClassDiagram_String(t *testing.T) {
	got := newClassDiagram().String()
	if !strings.Contains(got, "class") {
		t.Error("missing class attributes\n", got)
	}
}

func newClassDiagram() *ClassDiagram {
	var (
		d        = NewClassDiagram()
		car      = d.Struct(Car{})
		wheel    = d.Slice(make(Wheels, 4))
		stringer = d.Interface((*fmt.Stringer)(nil))
		driver   = d.Struct(Driver{})
	)
	d.Place(car).At(10, 10)
	d.Place(wheel).Below(car)
	d.Place(stringer).RightOf(wheel)
	d.Place(driver).Below(stringer, 100)
	d.Link(stringer, driver, "labeled")
	return d
}

type Car struct {
	Model string
	Wheels
}

// String
func (me Car) String() string { return me.Model }

// running
func (me *Car) Running() bool { return false }

type Wheel struct {
	Make string
}

type Wheels []Wheel

// String
func (me *Wheel) String() string { return me.Make }

type Driver struct{}
