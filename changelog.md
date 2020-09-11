# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [unreleased]

- Renamed NewSvg to NewSVG
- Moved type shape.Style to package draw

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
