package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

var figures []*TFigure

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
func BlackFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

// TFigure structure .
type TFigure struct {
	X int
	Y int
}

// Rectangle drawing operation.
func Rectangle(x1, y1, x2, y2 int) OperationFunc {
	return func(t screen.Texture) {
		rect := image.Rect(x1, y1, x2, y2)
		t.Fill(rect, color.Black, screen.Src)
	}
}

// Drawfigure returns function that draws.
func (f *TFigure) Drawfigure() OperationFunc {
	blueColor := color.RGBA{0, 0, 255, 255} // Adjusted alpha value to 255 for proper color representation
	return func(t screen.Texture) {
		upperRect := image.Rect(f.X-150, f.Y-100, f.X+150, f.Y)
		t.Fill(upperRect, blueColor, screen.Src)
		lowerRect := image.Rect(f.X-50, f.Y, f.X+50, f.Y+100)
		t.Fill(lowerRect, blueColor, screen.Src)
	}
}

// Moves figure.
func (f *TFigure) MoveFigure(x, y int) {
	f.X += x
	f.Y += y
}
