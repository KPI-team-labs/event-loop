package lang

import (
	"bufio"
	"errors"
	"github.com/roman-mazur/architecture-lab-3/painter"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	uistate *Uistate
}

func ParserWithState(state *Uistate) *Parser {
	return &Parser{uistate: state}
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), ",")
		if len(args) == 0 {
			continue
		}

		for _, arg := range args {
			commands := strings.Fields(arg)

			switch commands[0] {
			case "white":
				p.uistate.SetBgColor(painter.OperationFunc(painter.WhiteFill))
			case "green":
				p.uistate.SetBgColor(painter.OperationFunc(painter.GreenFill))
			case "bgrect":
				if len(commands) != 5 {
					return nil, errors.New("Invalid number of arguments for bgrect instruction.")
				}

				startXPoint, errXstart := strconv.ParseFloat(commands[1], 64)
				startYPoint, errYstart := strconv.ParseFloat(commands[2], 64)
				endXPoint, errXend := strconv.ParseFloat(commands[3], 64)
				endYPoint, errYend := strconv.ParseFloat(commands[4], 64)

				if errXstart != nil || errYstart != nil || errXend != nil || errYend != nil {
					return nil, errors.New("Invalid arguments for bgrect instruction.")
				}

				x1 := int(startXPoint * 800)
				y1 := int(startYPoint * 800)
				x2 := int(endXPoint * 800)
				y2 := int(endYPoint * 800)

				p.uistate.Rect = painter.Rectangle(x1, y1, x2, y2)
			case "figure":
				if len(commands) != 3 {
					return nil, errors.New("Invalid number of arguments for figure instruction.")
				}

				xCord, errXpoint := strconv.ParseFloat(commands[1], 64)
				yCord, errYpoint := strconv.ParseFloat(commands[2], 64)

				if errXpoint != nil || errYpoint != nil {
					return nil, errors.New("Invalid arguments for figure instruction.")
				}

				xCordPoint := int(xCord * 800)
				yCordPoint := int(yCord * 800)

				p.uistate.AddTFigure(&painter.TFigure{
					X: xCordPoint,
					Y: yCordPoint,
				})
			case "move":
				if len(commands) != 3 {
					return nil, errors.New("Invalid number of arguments for move instruction.")
				}

				xMove, errXmove := strconv.ParseFloat(commands[1], 64)
				yMove, errYmove := strconv.ParseFloat(commands[2], 64)

				if errXmove != nil || errYmove != nil {
					return nil, errors.New("Invalid arguments for move instruction.")
				}

				xMoveInt := int(xMove * 800)
				yMoveInt := int(yMove * 800)

				p.uistate.MoveFigures(xMoveInt, yMoveInt)
			case "update":
				res = append(res, p.uistate.Update()...)
			case "reset":
				p.uistate.Reset()
			default:
				return nil, errors.New("Invalid command.")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
