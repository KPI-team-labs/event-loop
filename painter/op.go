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
	StartPoint image.Point
	EndPoint   image.Point
}

func (op *Rectangle) Do(t screen.Texture) bool {
	c := color.Black
	t.Fill(image.Rect(op.StartPoint.X, op.StartPoint.Y, op.EndPoint.X, op.EndPoint.Y), c, screen.Src)
	return false
}

// TFigure structure
type TFigure struct {
	PointCenter image.Point
}

func (op *TFigure) Do(t screen.Texture) bool {
	c := color.RGBA{R: 0, G: 0, B: 255, A: 1}
	t.Fill(image.Rect(op.PointCenter.X-150, op.PointCenter.Y-100, op.PointCenter.X+150, op.PointCenter.Y), c, draw.Src)
	t.Fill(image.Rect(op.PointCenter.X-50, op.PointCenter.Y, op.PointCenter.X+50, op.PointCenter.Y+100), c, draw.Src)
	return false
}

type MoveFigures struct {
	X             int
	Y             int
	TFiguresArray []*TFigure
}

func (op *MoveFigures) Do(t screen.Texture) bool {
	for i := range op.TFiguresArray {
		op.TFiguresArray[i].PointCenter.X += op.X
		op.TFiguresArray[i].PointCenter.Y += op.Y
	}
	return false
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}
