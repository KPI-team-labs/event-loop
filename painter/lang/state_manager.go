package lang

import (
	"github.com/KPI-team-labs/event-loop/painter"
	"image"
)

type Uistate struct {
	backgroundColor painter.Operation
	Rect            *painter.Rectangle
	Figures         []*painter.TFigure
	moveOperations  []painter.Operation
	updateOperation painter.Operation
}

func (u *Uistate) GetOperations() []painter.Operation {
	var ops []painter.Operation

	if u.backgroundColor != nil {
		ops = append(ops, u.backgroundColor)
	}
	if u.Rect != nil {
		ops = append(ops, u.Rect)
	}
	if len(u.moveOperations) != 0 {
		ops = append(ops, u.moveOperations...)
		u.moveOperations = nil
	}
	if len(u.Figures) != 0 {
		for _, figure := range u.Figures {
			ops = append(ops, figure)
		}
	}
	if u.updateOperation != nil {
		ops = append(ops, u.updateOperation)
	}

	return ops
}

func (u *Uistate) ResetOperations() {
	if u.backgroundColor == nil {
		u.backgroundColor = painter.OperationFunc(painter.Reset)
	}
	if u.updateOperation != nil {
		u.updateOperation = nil
	}
}

func (u *Uistate) GreenBack() {
	u.backgroundColor = painter.OperationFunc(painter.GreenFill)
}

func (u *Uistate) WhiteBack() {
	u.backgroundColor = painter.OperationFunc(painter.WhiteFill)
}

func (u *Uistate) BackgRectangle(firstPoint image.Point, secondPoint image.Point) {
	u.Rect = &painter.Rectangle{
		StartPoint: firstPoint,
		EndPoint:   secondPoint,
	}
}

func (u *Uistate) AddTFigure(centralPoint image.Point) {
	figure := painter.TFigure{
		PointCenter: centralPoint,
	}
	u.Figures = append(u.Figures, &figure)
}

func (u *Uistate) AddMoveFigures(x int, y int) {
	moveOp := painter.MoveFigures{X: x, Y: y, TFiguresArray: u.Figures}
	u.moveOperations = append(u.moveOperations, &moveOp)
}

func (u *Uistate) ResetStateAndBackground() {
	u.Reset()
	u.backgroundColor = painter.OperationFunc(painter.Reset)
}

func (u *Uistate) SetUpdateOperation() {
	u.updateOperation = painter.UpdateOp
}

func (u *Uistate) Reset() {
	u.backgroundColor = nil
	u.Rect = nil
	u.Figures = nil
	u.moveOperations = nil
	u.updateOperation = nil
}
