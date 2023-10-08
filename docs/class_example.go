package docs

import (
	"github.com/gregoryv/draw"
	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
)

func ExampleClassDiagram() *design.ClassDiagram {
	var (
		d        = design.NewClassDiagram()
		record   = d.Struct(shape.Record{})
		circle   = d.Struct(shape.Circle{})
		diamond  = d.Struct(shape.Diamond{})
		triangle = d.Struct(shape.Triangle{})
		shapE    = d.Interface((*shape.Shape)(nil))
	)
	d.HideRealizations()

	var (
		fnt      = d.Struct(draw.Font{})
		style    = d.Struct(draw.Style{})
		seqdia   = d.Struct(design.SequenceDiagram{})
		classdia = d.Struct(design.ClassDiagram{})
		dia      = d.Struct(design.Diagram{})
		aligner  = d.Struct(shape.Aligner{})
		adj      = d.Struct(shape.Adjuster{})
		rel      = d.Struct(design.Relation{})
	)
	d.HideRealizations()

	d.Place(shapE).At(280, 20)
	d.Place(record).At(20, 160)

	d.Place(circle).RightOf(shapE, 200)
	d.Place(diamond, triangle, adj).Below(circle, 20)
	shape.Move(adj, -50, 0)

	d.Place(fnt).Below(record, 120)
	d.Place(style).RightOf(fnt, 90)
	d.VAlignCenter(record, fnt)

	d.Place(rel).Below(shapE, 20)
	d.Place(dia).RightOf(style, 70)
	d.Place(aligner).RightOf(dia, 60)
	d.HAlignCenter(fnt, style, dia, aligner)
	d.HAlignCenter(record, rel)

	d.Place(seqdia).Below(aligner, 90)
	d.Place(classdia).Below(dia, 90)
	d.VAlignCenter(dia, classdia)
	d.HAlignBottom(classdia, seqdia)

	d.SetCaption("Figure 1. Class diagram of design and design.shape packages")
	return d
}
