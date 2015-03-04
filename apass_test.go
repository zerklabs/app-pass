package apass

import "testing"

func TestGenerate(t *testing.T) {
	res := Generate(0)
	expectedLength := (DefaultSegmentCount * 8) + (DefaultSegmentCount - 1)

	if len(res) != expectedLength {
		t.Errorf("Incorrect length, expected: %d, found: %d", expectedLength, len(res))
	}

	t.Log(res)
}

func TestGenerate_One(t *testing.T) {
	length := 1

	res := Generate(length)
	expectedLength := (length * 8) + (length - 1)

	if len(res) != expectedLength {
		t.Errorf("Incorrect length, expected: %d, found: %d", expectedLength, len(res))
	}

	t.Log(res)
}

func TestGenerate_Four(t *testing.T) {
	length := 4
	// segments times segment length + number of spaces in result (segments - 1)
	expectedLength := (length * 8) + (length - 1)

	res := Generate(length)

	if len(res) != expectedLength {
		t.Errorf("Incorrect length, expected: %d, found: %d", expectedLength, len(res))
	}

	t.Log(res)
}

func TestGenerate_Negative(t *testing.T) {
	length := -1

	res := Generate(length)
	expectedLength := (DefaultSegmentCount * 8) + (DefaultSegmentCount - 1)

	if len(res) != expectedLength {
		t.Errorf("Incorrect length, expected: %d, found: %d", expectedLength, len(res))
	}

	t.Log(res)
}
