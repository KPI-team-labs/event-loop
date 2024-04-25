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

// Rectangle структура прямокутника
func Rectangle(x1, y1, x2, y2 int) OperationFunc {
	return func(t screen.Texture) {
		t.Fill(image.Rect(x1, y1, x2, y2), color.Black, screen.Src)
	}
}

// TFigure structure
type TFigure struct {
	X int
	Y int
}

// Draws TFigure on the screen
func (f *TFigure) Drawfigure() OperationFunc {
	c := color.RGBA{0, 0, 255, 1}
	return func(t screen.Texture) {
		t.Fill(image.Rect(f.X-150, f.Y-100, f.X+150, f.Y), c, screen.Src)
		t.Fill(image.Rect(f.X-50, f.Y, f.X+50, f.Y+100), c, screen.Src)
	}

}

// Structure for moving pictures

func (f *TFigure) MoveFigure(x, y int) {
	f.X += x
	f.Y += y
}
