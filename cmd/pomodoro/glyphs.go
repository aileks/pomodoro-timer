package main

import "strings"

const (
	baseDigitWidth  = 5
	baseDigitHeight = 7
	baseColonWidth  = 1
)

func renderDigit(scale int, digit rune) []string {
	pattern := digitPattern(digit)
	return scalePattern(pattern, scale, '█')
}

func renderColon(scale int) []string {
	pattern := colonPattern()
	return scalePattern(pattern, scale, '•')
}

func digitPattern(digit rune) []string {
	switch digit {
	case '0':
		return []string{
			"█████",
			"█   █",
			"█   █",
			"█   █",
			"█   █",
			"█   █",
			"█████",
		}
	case '1':
		return []string{
			"  █  ",
			" ██  ",
			"  █  ",
			"  █  ",
			"  █  ",
			"  █  ",
			"█████",
		}
	case '2':
		return []string{
			"█████",
			"    █",
			"    █",
			"█████",
			"█    ",
			"█    ",
			"█████",
		}
	case '3':
		return []string{
			"█████",
			"    █",
			"    █",
			"█████",
			"    █",
			"    █",
			"█████",
		}
	case '4':
		return []string{
			"█   █",
			"█   █",
			"█   █",
			"█████",
			"    █",
			"    █",
			"    █",
		}
	case '5':
		return []string{
			"█████",
			"█    ",
			"█    ",
			"█████",
			"    █",
			"    █",
			"█████",
		}
	case '6':
		return []string{
			"█████",
			"█    ",
			"█    ",
			"█████",
			"█   █",
			"█   █",
			"█████",
		}
	case '7':
		return []string{
			"█████",
			"    █",
			"    █",
			"    █",
			"    █",
			"    █",
			"    █",
		}
	case '8':
		return []string{
			"█████",
			"█   █",
			"█   █",
			"█████",
			"█   █",
			"█   █",
			"█████",
		}
	case '9':
		return []string{
			"█████",
			"█   █",
			"█   █",
			"█████",
			"    █",
			"    █",
			"█████",
		}
	default:
		return blankPattern(baseDigitWidth, baseDigitHeight)
	}
}

func colonPattern() []string {
	return []string{
		" ",
		" ",
		"•",
		" ",
		"•",
		" ",
		" ",
	}
}

func blankPattern(width int, height int) []string {
	lines := make([]string, height)
	for i := range lines {
		lines[i] = strings.Repeat(" ", width)
	}
	return lines
}

func scalePattern(pattern []string, scale int, fill rune) []string {
	if scale < 1 {
		scale = 1
	}
	lines := make([]string, 0, len(pattern)*scale)
	for _, row := range pattern {
		var builder strings.Builder
		for _, ch := range row {
			cell := ' '
			if ch != ' ' {
				cell = fill
			}
			for i := 0; i < scale; i++ {
				builder.WriteRune(cell)
			}
		}
		scaled := builder.String()
		for i := 0; i < scale; i++ {
			lines = append(lines, scaled)
		}
	}
	return lines
}
