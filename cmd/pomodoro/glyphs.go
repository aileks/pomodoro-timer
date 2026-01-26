package main

func renderDigit(scale int, digit rune) []string {
	width, height := digitSize(scale)
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	segments := digitSegments(digit)
	if segments[segmentA] {
		drawHorizontal(grid, 0, 1, width-2)
	}
	if segments[segmentB] {
		drawVertical(grid, width-1, 1, scale)
	}
	if segments[segmentC] {
		drawVertical(grid, width-1, scale+2, height-2)
	}
	if segments[segmentD] {
		drawHorizontal(grid, height-1, 1, width-2)
	}
	if segments[segmentE] {
		drawVertical(grid, 0, scale+2, height-2)
	}
	if segments[segmentF] {
		drawVertical(grid, 0, 1, scale)
	}
	if segments[segmentG] {
		drawHorizontal(grid, scale+1, 1, width-2)
	}

	lines := make([]string, height)
	for i := range grid {
		lines[i] = string(grid[i])
	}
	return lines
}

func renderColon(scale int) []string {
	_, height := digitSize(scale)
	lines := make([]string, height)
	for i := range lines {
		lines[i] = " "
	}
	upper := scale
	lower := scale + 2
	if upper >= 0 && upper < height {
		lines[upper] = "."
	}
	if lower >= 0 && lower < height {
		lines[lower] = "."
	}
	return lines
}

const (
	segmentA = iota
	segmentB
	segmentC
	segmentD
	segmentE
	segmentF
	segmentG
)

func digitSegments(digit rune) [7]bool {
	switch digit {
	case '0':
		return [7]bool{true, true, true, true, true, true, false}
	case '1':
		return [7]bool{false, true, true, false, false, false, false}
	case '2':
		return [7]bool{true, true, false, true, true, false, true}
	case '3':
		return [7]bool{true, true, true, true, false, false, true}
	case '4':
		return [7]bool{false, true, true, false, false, true, true}
	case '5':
		return [7]bool{true, false, true, true, false, true, true}
	case '6':
		return [7]bool{true, false, true, true, true, true, true}
	case '7':
		return [7]bool{true, true, true, false, false, false, false}
	case '8':
		return [7]bool{true, true, true, true, true, true, true}
	case '9':
		return [7]bool{true, true, true, true, false, true, true}
	default:
		return [7]bool{}
	}
}

func drawHorizontal(grid [][]rune, row int, start int, end int) {
	if row < 0 || row >= len(grid) {
		return
	}
	if start > end {
		return
	}
	if start < 0 {
		start = 0
	}
	if end >= len(grid[row]) {
		end = len(grid[row]) - 1
	}
	for col := start; col <= end; col++ {
		setRune(grid, row, col, '-')
	}
}

func drawVertical(grid [][]rune, col int, start int, end int) {
	if col < 0 || col >= len(grid[0]) {
		return
	}
	if start > end {
		return
	}
	if start < 0 {
		start = 0
	}
	if end >= len(grid) {
		end = len(grid) - 1
	}
	for row := start; row <= end; row++ {
		setRune(grid, row, col, '|')
	}
}

func setRune(grid [][]rune, row int, col int, value rune) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return
	}
	if grid[row][col] == ' ' {
		grid[row][col] = value
		return
	}
	if grid[row][col] != value {
		grid[row][col] = '+'
	}
}
