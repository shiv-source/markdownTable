# markdownTable

A Go package to generate Markdown tables from header and body data. This package automatically adjusts column widths based on the longest entry in each column, ensuring that the table is aligned and visually clear when rendered in Markdown.

## Features

- Converts a given header and body data into a well-formatted Markdown table.
- Automatically adjusts column widths based on the longest entry (including headers).
- Truncates rows with fewer columns than the header to maintain consistency.
- Provides a clean and simple way to create Markdown tables programmatically in Go.

## Installation

To install the `markdownTable` package, use the following Go command:

```bash
go get github.com/shiv-source/markdownTable
```

## Example Usage

Here’s an example of how to use the `CreateMarkdownTable` function:

```go
package main

import (
	"fmt"
	"github.com/shiv-source/markdownTable"
)

func main() {
	head := []string{"ID", "Name", "Age"}
	body := [][]string{
		{"1", "John", "26"},
		{"2", "Bob", "25"},
		{"3", "Alice", "27"},
	}

	// Create the Markdown table
	markdownTable := markdownTable.CreateMarkdownTable(head, body)

	// Print the Markdown table
	fmt.Println(markdownTable)
}
```

### Expected Output:

```markdown
| ID  | Name         | Age |
|-----|--------------|-----|
| 1   | John         | 26  |
| 2   | Bob          | 25  |
| 3   | Alice        | 27  |
```

## Function Signature

```go
func CreateMarkdownTable(head []string, body [][]string) string
```

### Parameters:
- `head`: A slice of strings representing the column headers of the table.
- `body`: A slice of slices, where each inner slice represents a row of data in the table.

### Returns:
- A string containing the Markdown-formatted table.

## Contributing

If you would like to contribute to this project, feel free to fork the repository, make changes, and create a pull request. All contributions are welcome!

### Steps to contribute:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.