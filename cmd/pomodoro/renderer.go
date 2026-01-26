package main

import "strings"

type Renderer struct {
	MinScale     int
	MaxScale     int
	DigitSpacing int
	ColonSpacing int
}

type RenderedBlock struct {
	Lines  []string
	Width  int
	Height int
	Scale  int
}

func NewRenderer() Renderer {
	return Renderer{
		MinScale:     1,
		MaxScale:     4,
		DigitSpacing: 1,
		ColonSpacing: 1,
	}
}

func (r Renderer) ScaleForWidth(timeStr string, maxWidth int) int {
	if maxWidth <= 0 {
		return clampInt(2, r.MinScale, r.MaxScale)
	}
	for scale := r.MaxScale; scale >= r.MinScale; scale-- {
		if r.WidthForScale(timeStr, scale) <= maxWidth {
			return scale
		}
	}
	return r.MinScale
}

func (r Renderer) WidthForScale(timeStr string, scale int) int {
	width := 0
	prevColon := false
	for i, ch := range timeStr {
		if i > 0 {
			if prevColon || ch == ':' {
				width += r.ColonSpacing
			} else {
				width += r.DigitSpacing
			}
		}
		if ch == ':' {
			width += 1
			prevColon = true
			continue
		}
		dw, _ := digitSize(scale)
		width += dw
		prevColon = false
	}
	return width
}

func (r Renderer) Render(timeStr string, maxWidth int) RenderedBlock {
	scale := r.ScaleForWidth(timeStr, maxWidth)
	width := r.WidthForScale(timeStr, scale)
	_, height := digitSize(scale)

	lines := make([]string, height)
	prevColon := false
	for i, ch := range timeStr {
		if i > 0 {
			spacing := r.DigitSpacing
			if prevColon || ch == ':' {
				spacing = r.ColonSpacing
			}
			space := strings.Repeat(" ", spacing)
			for row := range lines {
				lines[row] += space
			}
		}

		var glyph []string
		if ch == ':' {
			glyph = renderColon(scale)
			prevColon = true
		} else {
			glyph = renderDigit(scale, ch)
			prevColon = false
		}

		for row := range lines {
			lines[row] += glyph[row]
		}
	}

	return RenderedBlock{
		Lines:  lines,
		Width:  width,
		Height: height,
		Scale:  scale,
	}
}

func digitSize(scale int) (int, int) {
	h := scale + 1
	v := scale
	return h + 2, 2*v + 3
}
