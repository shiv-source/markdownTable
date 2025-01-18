package markdownTable

import (
	"testing"
)

func TestCreateMarkdownTable(t *testing.T) {
	tests := []struct {
		name     string
		head     []string
		body     [][]string
		expected string
	}{
		{
			name: "Basic table",
			head: []string{"ID", "Name", "Age"},
			body: [][]string{
				{"1", "John", "26"},
				{"2", "Bob", "25"},
				{"3", "Alice", "27"},
			},
			expected: "| ID | Name  | Age |\n| -- | ----- | --- |\n| 1  | John  | 26  |\n| 2  | Bob   | 25  |\n| 3  | Alice | 27  |\n",
		},
		{
			name: "Uneven rows",
			head: []string{"ID", "Name", "Age"},
			body: [][]string{
				{"1", "John", "26"},
				{"2", "Bob", "25"},
				{"3", "Alice", "27", "ExtraColumn"},
			},
			expected: "| ID | Name  | Age |\n| -- | ----- | --- |\n| 1  | John  | 26  |\n| 2  | Bob   | 25  |\n| 3  | Alice | 27  |\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateMarkdownTable(tt.head, tt.body)
			if got != tt.expected {
				t.Errorf("CreateMarkdownTable() = %v, want %v", got, tt.expected)
			}
		})
	}
}
