[![Build Status](https://travis-ci.org/gregoryv/go-design.svg?branch=master)](https://travis-ci.org/gregoryv/go-design)
[![codecov](https://codecov.io/gh/gregoryv/go-design/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/go-design)
[![Maintainability](https://api.codeclimate.com/v1/badges/b0001c5ba7cd098b183d/maintainability)](https://codeclimate.com/github/gregoryv/go-design/maintainability)

[go-design](https://godoc.org/github.com/gregoryv/go-design) - package for writing software design diagrams

- Cross platform
- No dependencies
- SVG output

Program your diagrams and refactoring automatically updates them.
Take a look at the below examples and then browse the [showcase](./showcase/README.md) of golang standard packages.

## Class diagram

This diagram is rendered by
[example_test.go/ExampleClassDiagram](https://godoc.org/github.com/gregoryv/go-design/#example-ClassDiagram)

<img src="img/class_example.svg" style="width: 500"/>


## Sequence diagram

From [example_test.go/ExampleSequenceDiagram](https://godoc.org/github.com/gregoryv/go-design/#example-SequenceDiagram)

![](img/sequence_example.svg)

## Generic diagram

It should be easy to just add any extra shapes to any diagram when explaining a design.
This diagram is rendered by
[example_test.go/ExampleDiagram](https://godoc.org/github.com/gregoryv/go-design/#example-Diagram)

![](img/diagram_example.svg)


## TODO

- Labeled arrows
- Link to optional godoc service
- More shapes

WIP - major rewrites still going on
