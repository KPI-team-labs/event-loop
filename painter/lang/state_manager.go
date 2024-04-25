package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Uistate preserves UI elements.
type Uistate struct {
	BackgroundColor painter.OperationFunc
	Rect            painter.OperationFunc
	Figures         []*painter.TFigure
}

// Instance of NewUistata.
func NewUistate() *Uistate {
	return &Uistate{}
}

// Sets background color.
func (u *Uistate) SetBackgroundColor(op painter.OperationFunc) {
	u.BackgroundColor = op
}

// Draw operation.
func (u *Uistate) SetRect(op painter.OperationFunc) {
	u.Rect = op
}

// Reset all operations.
func (u *Uistate) Reset() {
	u.BackgroundColor = painter.OperationFunc(painter.BlackFill)
	u.Rect = painter.Rectangle(0, 0, 0, 0)
	u.Figures = nil
}

// Update .
func (u *Uistate) Update() []painter.Operation {
	var operations []painter.Operation

	if u.BackgroundColor != nil {
		operations = append(operations, u.BackgroundColor)
	}
	if u.Rect != nil {
		operations = append(operations, u.Rect)
	}
	for _, figure := range u.Figures {
		operations = append(operations, figure.Drawfigure())
	}
	operations = append(operations, painter.UpdateOp)

	return operations
}

// MoveFigures moves all figures to specified coordinates.
func (u *Uistate) MoveFigures(dx, dy int) {
	for _, figure := range u.Figures {
		figure.MoveFigure(dx, dy)
	}
}
