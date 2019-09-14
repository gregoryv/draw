[![Build Status](https://travis-ci.org/gregoryv/go-design.svg?branch=master)](https://travis-ci.org/gregoryv/go-design)
[![codecov](https://codecov.io/gh/gregoryv/go-design/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/go-design)
[![Maintainability](https://api.codeclimate.com/v1/badges/b0001c5ba7cd098b183d/maintainability)](https://codeclimate.com/github/gregoryv/go-design/maintainability)

[go-design](https://godoc.org/github.com/gregoryv/go-design) - package for writing software design diagrams in code

Wouldn't it be nice if gofmt also updated all diagrams while you are
refactoring?  If the diagrams are code, it would. This package allows
you to easily create various diagrams as a program.

Diagrams and shapes are all in SVG which is good enough for most
purposes.

WIP - major refactorings still going on

## Class diagram

This diagram is rendered by
[example_test.go/ExampleClassDiagram](https://godoc.org/github.com/gregoryv/go-design/#example-ClassDiagram)

![](img/class_example.svg)


## Sequence diagram

From [example_test.go/ExampleSequenceDiagram](https://godoc.org/github.com/gregoryv/go-design/#example-SequenceDiagram)

![](img/sequence_example.svg)

## Generic diagram

It should be easy to just add any extra shapes to any diagram when explaining a design.
This diagram is rendered by
[example_test.go/ExampleDiagram](https://godoc.org/github.com/gregoryv/go-design/#example-Diagram)

![](img/diagram_example.svg)
