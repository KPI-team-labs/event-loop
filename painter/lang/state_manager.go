package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
)

type Uistate struct {
	backgroundColor painter.OperationFunc
	Rect            painter.OperationFunc
	Figures         []*painter.TFigure
}

func NewUistate() *Uistate {
	return &Uistate{}
}

func (f *Uistate) SetBgColor(op painter.OperationFunc) {
	f.backgroundColor = op
}

func (f *Uistate) SetRect(op painter.OperationFunc) {
	f.Rect = op
}

func (f *Uistate) Reset() {
	f.backgroundColor = painter.OperationFunc(painter.BlackFill)
	f.Rect = painter.Rectangle(0, 0, 0, 0)
	f.Figures = nil
}

func (f *Uistate) Update() []painter.Operation {
	var result []painter.Operation

	if f.backgroundColor != nil {
		result = append(result, f.backgroundColor)
	}

	if f.Rect != nil {
		result = append(result, f.Rect)

	}

	for _, figure := range f.Figures {
		result = append(result, figure.Drawfigure())
	}

	result = append(result, painter.UpdateOp)
	return result

}

func (f *Uistate) MoveFigures(dx, dy int) {
	for _, figure := range f.Figures {
		figure.MoveFigure(dx, dy)
	}
}
