# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [0.21.1] 2021-04-21
## [0.21.0] 2021-04-21

- Line and arrow Height and Width methods consider head and tail shapes
- Diamond.Position returns top left
- Triangle points upwards and Position returns top left
- Replace ActivityDiagram.If with Or when switching decision
- Replace ActivityDiagram.Then with Trans and add TransRight methods
- Activity diagram default spacing increased from 40 to 60

## [0.20.0] 2021-04-16

- Add Label.SetHref and Component.SetHref
- Rename type xy.Position to xy.Point
- Show relations between slices and structs in ClassDiagram
- Add SequenceDiagram.Skip method for adding a dashed spacing
- Replace draw.TagWriter with nexus.Printer

## [0.19.0] 2021-04-05

- Add labeled hexagon shape
- Fix label alignment
- Adjuster respects diagram style.Spacing
- NewDiagram returns a pointer to diagram
- Diagram variations embed reference to diagram

## [0.18.0] 2021-01-23

- Include slice methods in VRecord
- Remove double padding for record titles
- Improve label alignment in Diagram.Link method

## [0.17.0] 2020-12-16

- Fix Class diagram to show pointer receiver methods
- Add SequenceDiagram.Group method for grouping columns with a labeled text area

## [0.16.0] 2020-10-30

- Diagrams Stringer implementations no longer inline css, use Inline method
- Add VRecord.Slice method for slice composition and aggregates

## [0.15.1] 2020-10-27

- Fix missing relations in inlined class diagrams

## [0.15.0] 2020-10-27

- Diagrams implement stringer interface
- AdaptSize makes width and height +1 pixel
- Add Inline method to diagrams
- Add type ClassAttributes with CSS method for easy inlining in HTML

## [0.14.0] 2020-10-06

- Fix VRecord to show pointer methods if pointer to struct given

## [0.13.0] 2020-09-12

- Add shape.Internet
- Add default spacing to shape.Adjuster
- Move type shape.Style to package draw
- Rename NewSvg to NewSVG

## [0.12.0] 2020-09-09
### Changed

- Moved draw/shape/design package to draw/design
- Renamed type Svg to SVG
- Renamed type SvgWriter to SVGWriter
- Renamed WriteSvg to WriteSVG

## [0.11.0] 2020-03-17
### Changed

- Generic design.NewVRecord for both structs and interfaces

### Removed

- Specific methods for creating struct and interface methods.

## [0.10.0] 2020-02-03
### Added

- Ganttchart weekly view
- Cylinder and database shapes
- Func shape.SetClass for setting class of many shapes

### Fixed

- Component and rect labels use class + "-title"

## [0.9.0] 2020-01-02
### Added

- ActivityDiagram helper methods, e.g Start, Then, If and Exit
- Mark current day by default in gantt chart

### Changed

- Renamed Direction constants LR and RL to RightDir and LeftDir
- NewGanttChart constructor uses date.String
- Exposed TagWriter
- shape.Svg moved to draw.Svg

### Fixed

- Label positioned correctly in activity diagram for vertical arrows

## [0.8.0] 2019-12-18
### Added

- Aggregate relation when pointer to specific type is used
- Gantt chart
- Actor shape

### Changed

- Label is optional when creating links with func Link
- Hide Height, Width attributes in Diagram
- shape.NewDot defaults to radius 10

## [0.7.0] - 2019-12-15
### Changed

- Namespace from gregoryv/go-design to gregoryv/draw. Fix your imports
  import "github.com/gregoryv/draw/shape/design"

## [0.6.0] - 2019-12-15
### Added

- Activity diagram examples
- Labeled links in plain diagrams
- State shape
- PlaceGrid to quickly place many shapes into a grid layout

## [0.5.0] - 2019-12-13
### Added

- Link method to generic diagram
- Rect, Dot, ExitDot and Component shapes

### Changed

- Label and circle has an edge
- Circle shape is stroked by default

## [0.4.0] - 2019-12-04
### Fixed

- Caption in sequence diagrams

### Changed

- Default color of records is white with pale lines


## [0.3.0] - 2019-10-12
### Fixed

- Arrows point to edge irrelevant of angle

### Added

- Diagrams have optional bottom centered caption
- Class diagram shows composition using diamond for tail
- Note shape with multiline support

### Changed

- Font size decoupled from class attributes
- Arrow for realizing interface is dashed
- Label and circle position is top left corner
- Hide methods realized by visible interfaces and structs


## [0.2.0] - 2019-10-03
### Changed

- All shapes can have a class
- Arrows in class diagrams attach to edge and point to center

### Fixed

- RightOf and LeftOf adjustments sets matching y position
- Use class for non self arrows

### Removed

- SvgWriterShape interface


## [0.1.0] - 2019-09-21
### Added

- Sequence diagram
- Class diagram
- Record shape with fields and methods
- Arrow, line and label shapes
