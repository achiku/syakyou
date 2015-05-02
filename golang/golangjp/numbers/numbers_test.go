package numbers

import "testing"

type testData struct {
	in, out int
}

var doubleTests = []testData{
	testData{1, 2},
	testData{2, 4},
	testData{-5, -10},
}

var tripleTests = []testData{
	testData{1, 3},
	testData{2, 6},
	testData{-5, -15},
}

var squareTests = []testData{
	testData{1, 1},
	testData{2, 4},
	testData{-5, 25},
}

func TestDouble(t *testing.T) {
	for _, dt := range doubleTests {
		v := Double(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d", dt.in, v, dt.out)
		}
	}
}

func TestTriple(t *testing.T) {
	for _, dt := range tripleTests {
		v := Triple(dt.in)
		if v != dt.out {
			t.Errorf("Triple(%d) = %d, want %d", dt.in, v, dt.out)
		}
	}
}

func TestSquare(t *testing.T) {
	for _, dt := range squareTests {
		v := Square(dt.in)
		if v != dt.out {
			t.Errorf("Square(%d) = %d, want %d", dt.in, v, dt.out)
		}
	}
}
