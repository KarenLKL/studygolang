package testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	data := []struct{ a, b, c int32 }{
		{1, 2, 3},
		{2, 2, 4},
		{3, 3, 6},
		{2, 2, 5},
	}

	for _, item := range data {
		if result := Add(item.a, item.b); item.c != result {
			t.Errorf("except add(%d,%d) got %d,except %d", item.a, item.b, result, item.c)
		}
	}
}
