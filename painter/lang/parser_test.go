package lang

import (
	"reflect"
	"strings"
	"testing"

	"github.com/KPI-team-labs/event-loop/painter"
)

func TestParser_Parse(t *testing.T) {
	// Create a sample UI state for testing
	state := NewUistate()

	// Create a parser instance with the sample UI state
	parser := ParserWithState(state)

	// Define test cases
	tests := []struct {
		name     string
		input    string
		expected []painter.Operation
		wantErr  bool
	}{
		{
			name:  "ValidWhiteCommand",
			input: "white, update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.WhiteFill),
				painter.UpdateOp,
			},
			wantErr: false,
		},
		{
			name:  "ValidBgRectCommand",
			input: "bgrect 0.1 0.2 0.3 0.4, update",
			expected: []painter.Operation{
				painter.OperationFunc(painter.WhiteFill),
				painter.Rectangle(80, 160, 240, 320),
				painter.UpdateOp,
			},
			wantErr: false,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare input as io.Reader
			reader := strings.NewReader(tt.input)

			// Parse the input
			got, err := parser.Parse(reader)

			// Check for errors
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Compare parsed result with expected
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse() = %v, want %v", got, tt.expected)
			}
		})
	}
}
