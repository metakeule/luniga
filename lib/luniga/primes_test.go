package luniga

import (
	"testing"
)

func TestPhonetic(t *testing.T) {
	corpus := map[Transscriber]string{
		FOP: `fɔp`,
	}

	for tr, expected := range corpus {
		result := Phonetic(tr)
		if result != expected {
			t.Errorf("expected: %#v, got: %#v", expected, result)
		}
	}

}
