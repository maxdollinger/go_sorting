package data

import (
	"strings"
)

func (f *Formater) TableString() string {
	if len(f.evals) == 0 {
		return "No evaluation data available"
	}

	rows := f.prepareTableData()
	colWidths := f.calculateColumnWidths(rows)

	return f.buildTable(rows, colWidths)
}

func (f *Formater) prepareTableData() [][]string {
	headers := []string{"Method", "Dist", "N", "Q1", "Median", "3Q", "Min", "Max", "Memory"}
	rows := [][]string{headers}

	for _, eval := range f.evals {
		row := []string{
			eval.Method,
			eval.Distribution,
			numShort(eval.Size),
			timeFormat(eval.ExectimeP(0.25)),
			timeFormat(eval.ExectimeMedian()),
			timeFormat(eval.ExectimeP(0.75)),
			timeFormat(eval.ExectimeFastest()),
			timeFormat(eval.ExectimeSlowest()),
			eval.MemoryMeanStr(),
		}
		rows = append(rows, row)
	}

	return rows
}

func (f *Formater) calculateColumnWidths(rows [][]string) []int {
	if len(rows) == 0 {
		return nil
	}

	colWidths := make([]int, len(rows[0]))
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) {
				cellWidth := runeWidth(cell)
				if cellWidth > colWidths[i] {
					colWidths[i] = cellWidth
				}
			}
		}
	}

	return colWidths
}

func (f *Formater) buildTable(rows [][]string, colWidths []int) string {
	var result strings.Builder

	// Top border
	result.WriteString(buildBorder(colWidths, "┌", "┬", "┐"))
	result.WriteString("\n")

	// Header row
	result.WriteString(buildRow(rows[0], colWidths))
	result.WriteString("\n")

	// Header separator
	result.WriteString(buildBorder(colWidths, "├", "┼", "┤"))
	result.WriteString("\n")

	// Data rows
	for i := 1; i < len(rows); i++ {
		// Separator for each input size
		if i > 1 && rows[i][2] != rows[i-1][2] {
			result.WriteString(buildBorder(colWidths, "├", "┼", "┤"))
			result.WriteString("\n")
		}

		result.WriteString(buildRow(rows[i], colWidths))
		result.WriteString("\n")
	}

	// Bottom border
	result.WriteString(buildBorder(colWidths, "└", "┴", "┘"))

	return result.String()
}

func runeWidth(s string) int {
	return len([]rune(s))
}

func buildRow(row []string, colWidths []int) string {
	var result strings.Builder
	result.WriteString("│")

	for i, cell := range row {
		if i >= len(colWidths) {
			continue
		}

		cellWidth := runeWidth(cell)
		padding := colWidths[i] - cellWidth

		if padding < 0 {
			padding = 0
		}

		result.WriteString(" ")
		result.WriteString(cell)
		if padding > 0 {
			result.WriteString(strings.Repeat(" ", padding))
		}
		result.WriteString(" │")
	}

	return result.String()
}

func buildBorder(colWidths []int, left, middle, right string) string {
	var result strings.Builder
	result.WriteString(left)

	for i, width := range colWidths {
		result.WriteString(strings.Repeat("─", width+2))
		if i < len(colWidths)-1 {
			result.WriteString(middle)
		}
	}

	result.WriteString(right)
	return result.String()
}
