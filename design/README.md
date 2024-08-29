[design](https://godoc.org/github.com/gregoryv/draw/design) - package for writing software design diagrams

## Sequence diagram

<img src="img/app_sequence_diagram.svg">

    var (
        d   = design.NewSequenceDiagram()
        cli = d.AddStruct(app.Client{})
        srv = d.AddStruct(app.Server{})
        db  = d.AddStruct(sql.DB{})
    )
    d.Link(cli, srv, "connect()")
    d.Link(srv, db, "SELECT").Class = "highlight"
    d.Link(db, srv, "Rows")
    d.Link(srv, srv, "Transform to view model").Class = "highlight"
    d.Skip()
	d.Return(srv, cli, "Send HTML")

## Activity diagram

<img src="img/activity_diagram.svg">

Rendered by
[ExampleActivityDiagram](https://godoc.org/github.com/gregoryv/draw/design/#example-ActivityDiagram)

## Class diagram

Class diagrams show relations between structs and
interfaces. Reflection includes fields and methods.

This diagram is rendered by
[ExampleClassDiagram](https://godoc.org/github.com/gregoryv/draw/design/#example-ClassDiagram)

<img src="img/class_example.svg">

## Generic diagram

It should be easy to just add any extra shapes to any diagram when explaining a design.
This diagram is rendered by
[ExampleDiagram](https://godoc.org/github.com/gregoryv/draw/design/#example-Diagram)

![](img/diagram_example.svg)


## C4 diagram

Read more about the [C4 model](https://c4model.com/).

[Example C4 diagrams](https://gregoryv.github.io/draw/#c4diagrams)

![](img/c4_example.svg)

## Grid layout

Simplifying placing shapes in a grid layout aligning different sizes of shapes.

![](img/grid_layout.svg)


## Gantt chart

![](img/gantt_chart.svg)

[ExampleGanttChart](https://godoc.org/github.com/gregoryv/draw/design/#example-GanttChart)

## Showcase

You can find more examples in the [showcase](showcase) folder.
