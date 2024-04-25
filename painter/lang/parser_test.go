package lang

import (
	"bytes"
	"testing"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []painter.Operation
		wantErr  bool
	}{
		{
			name:     "Parse White Fill",
			input:    `white`,
			expected: []painter.Operation{painter.OperationFunc(painter.WhiteFill)},
			wantErr:  false,
		},
		{
			name:     "Parse Green Fill",
			input:    `green`,
			expected: []painter.Operation{painter.OperationFunc(painter.GreenFill)},
			wantErr:  false,
		},
		{
			name:     "Parse Black Rectangle",
			input:    `bgrect 10 20 30 40`,
			expected: []painter.Operation{&painter.BlackRectangle{X1: 10, X2: 20, Y1: 30, Y2: 40}},
			wantErr:  false,
		},
		{
			name:     "Parse Cross Figure",
			input:    `figure 5 5`,
			expected: []painter.Operation{&painter.CrossFigure{X: 5, Y: 5}},
			wantErr:  false,
		},
		{
			name:     "Parse Move Operation",
			input:    `move 10 20`,
			expected: []painter.Operation{&painter.MoveOperation{X: 10, Y: 20, Crosses: []*painter.CrossFigure{}}},
			wantErr:  false,
		},
		{
			name:     "Parse Reset",
			input:    `reset`,
			expected: []painter.Operation{painter.OperationFunc(painter.Reset)},
			wantErr:  false,
		},
		{
			name:     "Parse Update Operation",
			input:    `update`,
			expected: []painter.Operation{painter.UpdateOp},
			wantErr:  false,
		},
		{
			name:     "Invalid Command",
			input:    `invalid`,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := Parser{}
			output, err := parser.Parse(bytes.NewReader([]byte(tt.input)))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b := output == nil
			if b != (tt.expected == nil) {
				t.Errorf("Unexpected nil")
			}
		})
	}
}
