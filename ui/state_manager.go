package ui

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"image/color"
)

type Uistate struct {
	backgroundColor     painter.Operation
	rectangleBackground *painter.Rectangle
	shapeArray          []*painter.TFigure
	moveOperations      []painter.Operation
	updateOperation     painter.Operation
}

func (u *Uistate) GetOperations() []painter.Operation {
	var ops []painter.Operation
	if u.backgroundColor != nil {
		ops = append(ops, u.backgroundColor)
	}

	if u.rectangleBackground != nil {
		ops = append(ops, u.rectangleBackground)
	}
	ops = append(ops, u.moveOperations...)
	u.moveOperations = nil
	for _, figure := range u.shapeArray {
		ops = append(ops, figure)
	}
	u.shapeArray = nil

	if u.updateOperation != nil {
		ops = append(ops, u.updateOperation)
	}

	return ops
}

func (u *Uistate) ResetOperations() {
	u.backgroundColor = painter.OperationFunc(painter.ResetScreen)
	if u.updateOperation != nil {
		u.updateOperation = nil
	}
}

func (u *Uistate) GreenBackground() {
	u.backgroundColor = painter.OperationFunc(painter.GreenFill)
}

func (u *Uistate) WhiteBackground() {
	u.backgroundColor = painter.OperationFunc(painter.WhiteFill)
}

func (u *Uistate) BackgroundRectangle(x1, y1, x2, y2 int) {
	u.rectangleBackground = &painter.Rectangle{
		X1: x1, Y1: y1, X2: x2, Y2: y2,
	}
}

func (u *Uistate) AddFigure(x, y int, rgba color.RGBA) {
	figure := &painter.TFigure{
		X: x, Y: y, COLOR: rgba,
	}
	u.shapeArray = append(u.shapeArray, figure)
}

func (u *Uistate) AddMoveOperation(x, y int) {
	moveOp := &painter.Move{
		X: x, Y: y, Figures: u.shapeArray,
	}
	u.moveOperations = append(u.moveOperations, moveOp)
}

func (u *Uistate) ResetStateAndBackground() {
	u.Reset()
	u.backgroundColor = painter.OperationFunc(painter.ResetScreen)
}

func (u *Uistate) SetUpdateOperation() {
	u.updateOperation = painter.UpdateOp
}

func (u *Uistate) Reset() {
	u.backgroundColor = nil
	u.rectangleBackground = nil
	u.shapeArray = nil
	u.moveOperations = nil
	u.updateOperation = nil
}
