package lang

import (
	"fmt"
	"github.com/roman-mazur/architecture-lab-3/painter"
	"golang.org/x/exp/shiny/screen"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

var figures []*painter.TFigure

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	argsBytes, err := ioutil.ReadAll(in)
	var res []painter.Operation

	if err != nil {
		return nil, err
	}

	argsStr := string(argsBytes)
	args := strings.Split(argsStr, ",")

	for _, arg := range args {
		op := CommandHandler(strings.TrimSpace(arg))
		if op != nil {
			res = append(res, op)
		}
	}

	return res, nil
}

func CommandHandler(command string) painter.Operation {
	commands := map[string]func(command string) painter.Operation{
		"white": func(command string) painter.Operation {
			return painter.OperationFunc(painter.WhiteFill)
		},
		"green": func(command string) painter.Operation {
			return painter.OperationFunc(painter.GreenFill)
		},
		"update": func(command string) painter.Operation {
			return painter.UpdateOp
		},
		"bgrect": BgRectHandler,
		"figure": TwoArgsHandler,
		"move":   TwoArgsHandler,
		"reset": func(command string) painter.Operation {
			figures = nil
			return painter.OperationFunc(painter.ResetScreen)
		},
	}

	for commandStr := range commands {
		if strings.HasPrefix(command, commandStr) {
			return commands[commandStr](command)
		}
	}

	fmt.Printf("Command %s does not exist\n", command)
	return nil
}

func BgRectHandler(command string) painter.Operation {
	args := strings.Split(command, " ")
	startX := args[1]
	startXPoint, _ := strconv.Atoi(startX)
	startY := args[2]
	startYPoint, _ := strconv.Atoi(startY)
	endX := args[3]
	endXPoint, _ := strconv.Atoi(endX)
	endY := args[4]
	endYPoint, _ := strconv.Atoi(endY)
	rect := &painter.Rectangle{
		X1: startXPoint,
		Y1: startYPoint,
		X2: endXPoint,
		Y2: endYPoint,
	}
	return painter.OperationFunc(func(t screen.Texture) {
		rect.Do(t)
	})
}

func TwoArgsHandler(command string) painter.Operation {
	args := strings.Split(command, " ")
	commandType := args[0]
	xCord := args[1]
	xCordPoint, _ := strconv.Atoi(xCord)
	yCord := args[2]
	yCordPoint, _ := strconv.Atoi(yCord)
	var tCommand painter.Operation
	if commandType == "figure" {
		tCommand = &painter.TFigure{
			X: xCordPoint,
			Y: yCordPoint,
		}
		figures = append(figures, tCommand.(*painter.TFigure))
	} else if commandType == "move" {
		tCommand = &painter.Move{
			X:       xCordPoint,
			Y:       yCordPoint,
			Figures: figures,
		}
	}

	return painter.OperationFunc(func(t screen.Texture) {
		tCommand.Do(t)
	})
}
