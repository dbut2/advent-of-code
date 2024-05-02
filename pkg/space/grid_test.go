package space

import (
	"testing"
)

func TestGrow(t *testing.T) {
	g := NewGrid[int](2, 2)

	if len(g) != 2 {
		t.Errorf("Expected len(g) == 2, got %v", len(g))
	}

	if len(g[0]) != 2 {
		t.Errorf("Expected len(g[0]) == 2, got %v", len(g[0]))
	}

	g.Set(Cell{5, 5}, 1)

	if len(g) != 6 {
		t.Errorf("Expected len(g) == 6, got %v", len(g))
	}

	if len(g[0]) != 6 {
		t.Errorf("Expected len(g[0]) == 6, got %v", len(g[0]))
	}
}
