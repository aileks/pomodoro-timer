package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRendererScaleForWidth(t *testing.T) {
	renderer := NewRenderer()
	value := "25:00"

	width1 := renderer.WidthForScale(value, 1)
	width2 := renderer.WidthForScale(value, 2)
	width3 := renderer.WidthForScale(value, 3)

	if got := renderer.ScaleForWidth(value, width2-1); got != 1 {
		t.Fatalf("expected scale 1 for width %d, got %d", width2-1, got)
	}
	if got := renderer.ScaleForWidth(value, width2); got != 2 {
		t.Fatalf("expected scale 2 for width %d, got %d", width2, got)
	}
	if got := renderer.ScaleForWidth(value, width3); got != 3 {
		t.Fatalf("expected scale 3 for width %d, got %d", width3, got)
	}
	if got := renderer.ScaleForWidth(value, width1-1); got != 1 {
		t.Fatalf("expected scale 1 for width %d, got %d", width1-1, got)
	}
}

func TestRendererGolden(t *testing.T) {
	times := []string{"00:00", "09:59", "25:00", "59:59"}
	for _, scale := range []int{1, 2, 3} {
		renderer := NewRenderer()
		renderer.MinScale = scale
		renderer.MaxScale = scale

		content := renderSamples(renderer, times)
		golden := filepath.Join("testdata", fmt.Sprintf("renderer_scale%d.golden", scale))
		assertGolden(t, golden, content)
	}
}

func renderSamples(renderer Renderer, times []string) string {
	sections := make([]string, 0, len(times))
	for _, value := range times {
		block := renderer.Render(value, 0)
		section := strings.Join([]string{
			"time " + value,
			strings.Join(block.Lines, "\n"),
		}, "\n")
		sections = append(sections, section)
	}
	return strings.Join(sections, "\n\n") + "\n"
}

func assertGolden(t *testing.T, path string, content string) {
	t.Helper()
	if os.Getenv("UPDATE_GOLDEN") == "1" {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatalf("write golden: %v", err)
		}
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	if string(data) != content {
		t.Fatalf("golden mismatch for %s", path)
	}
}
