package design

func NewActivityDiagram() *ActivityDiagram {
	return &ActivityDiagram{
		Diagram: NewDiagram(),
	}
}

type ActivityDiagram struct {
	Diagram
}
