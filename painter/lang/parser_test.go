package lang

import (
	"image"
	"strings"
	"testing"

	"github.com/KPI-team-labs/event-loop/painter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	// Array of structs containing test cases.
	tests := []struct {
		name     string
		input    string
		expected painter.Operation
	}{
		{
			name:     "valid bgrect instruction",
			input:    "bgrect 0.3 0.3 0.6 0.6",
			expected: &painter.Rectangle{StartPoint: image.Point{X: 240, Y: 240}, EndPoint: image.Point{X: 480, Y: 480}},
		},
		{
			name:     "valid figure instruction",
			input:    "figure 0 0.55",
			expected: &painter.TFigure{PointCenter: image.Point{X: 0, Y: 440}},
		},
		{
			name:     "valid move instruction",
			input:    "move 0.05 0.05",
			expected: &painter.MoveFigures{X: 40, Y: 40},
		},
		{
			name:     "update instruction",
			input:    "update",
			expected: painter.UpdateOp,
		},
		{
			name:     "invalid instruction",
			input:    "hello",
			expected: nil,
		},
		{
			name:     "bgrect instruction with fewer arguments",
			input:    "bgrect 0.15 0.15 0.85",
			expected: nil,
		},
		{
			name:     "bgrect instruction with more arguments",
			input:    "bgrect 0.15 0.15 0.85 0.85 0.65",
			expected: nil,
		},
		{
			name:     "bgrect instruction with invalid arguments",
			input:    "bgrect h e l o",
			expected: nil,
		},
		{
			name:     "bgrect instruction with invalid argument (just a word)",
			input:    "bgrect 0.1 0.1 0.9 o",
			expected: nil,
		},
		{
			name:     "bgrect instruction with invalid argument (word and number)",
			input:    "bgrect 0.1 0.1 0.9 0.9o",
			expected: nil,
		},
		{
			name:     "figure instruction with fewer arguments",
			input:    "figure 0.95",
			expected: nil,
		},
		{
			name:     "figure instruction with more arguments",
			input:    "figure 0.95 0.6 0.2",
			expected: nil,
		},
		{
			name:     "figure instruction with invalid arguments",
			input:    "figure h i",
			expected: nil,
		},
		{
			name:     "figure instruction with invalid argument (just a word)",
			input:    "figure h 0.7",
			expected: nil,
		},
		{
			name:     "figure instruction with invalid argument (word and number)",
			input:    "figure 0.7 0.7i",
			expected: nil,
		},
		{
			name:     "move instruction with fewer arguments",
			input:    "move 0.35",
			expected: nil,
		},
		{
			name:     "move instruction with more arguments",
			input:    "move 0.35 0 0.15",
			expected: nil,
		},
		{
			name:     "move instruction with invalid arguments",
			input:    "move g g",
			expected: nil,
		},
		{
			name:     "move instruction with invalid argument (just a word)",
			input:    "move 0.1 q",
			expected: nil,
		},
		{
			name:     "move instruction with invalid argument (word and number)",
			input:    "move 0.05 i0.05",
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare input as io.Reader
			parser := &Parser{}

			// Parse the input
			got, err := parser.Parse(strings.NewReader(tt.input))

			// Check for errors
			if tt.expected == nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				// Asserting that the type of the parsed operation is the same as the expected operation.
				assert.IsType(t, tt.expected, got[1])
				// Asserting that the parsed operation is equal to the expected operation.
				assert.Equal(t, tt.expected, got[1])
			}
		})
	}
}

func TestParser_Fill(t *testing.T) {

	// A list of test cases consisting of commands to be parsed by the parser and their expected operations
	tests := []struct {
		name     string
		input    string
		expected painter.Operation
	}{
		{
			name:     "white instruction",
			input:    "white",
			expected: painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:     "green instruction",
			input:    "green",
			expected: painter.OperationFunc(painter.GreenFill),
		},
		{
			name:     "reset instruction",
			input:    "reset",
			expected: painter.OperationFunc(painter.Reset),
		},
	}

	// Create a new parser
	parser := &Parser{}

	// Iterate through the list of test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the command using the parser
			got, err := parser.Parse(strings.NewReader(tt.input))

			// Assert that there are no errors
			require.NoError(t, err)

			// Assert that the number of operations returned is 1
			require.Len(t, got, 1)

			// Assert that the type of the first operation is the expected type
			assert.IsType(t, tt.expected, got[0])
		})
	}
}
