# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

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
