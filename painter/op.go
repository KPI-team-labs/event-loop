package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

// Rectangle структура прямокутника
type Rectangle struct {
	startPoint image.Point
	endPoint   image.Point
}

// Draws rectangle
func (op *Rectangle) Do(t screen.Texture) bool {
	rect_color := color.Black
	t.Fill(image.Rect(op.startPoint.X, op.startPoint.Y, op.endPoint.X, op.endPoint.Y), rect_color, screen.Src)
	return false
}

// TFigure structure
type TFigure struct {
	X, Y  int
	COLOR color.RGBA
}

// Draws TFigure on the screen
func (op *TFigure) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.X-150, op.Y+100, op.X+150, op.Y), op.COLOR, draw.Src)
	t.Fill(image.Rect(op.X-45, op.Y+45, op.X+45, op.Y+200), op.COLOR, draw.Src)
	return false
}

// Structure for moving pictures
type Move struct {
	X, Y    int
	Figures []*TFigure
}

func (op *Move) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].X += op.X
		op.Figures[i].Y += op.Y
	}
	return false
}

func ResetScreen(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, draw.Src)
}
