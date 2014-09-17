package v0_1

import (
	"github.com/metakeule/luniga/lib/luniga"

	"testing"
)

func TestDE(t *testing.T) {

	corpus := map[string]string{
		`@1`:   `ICH`,
		`@1##`: `ich heiße`,
	}

	for in, expected := range corpus {
		tr := luniga.Parse(in)

		w := &luniga.Word{Head: tr[0].(luniga.Prime), Tail: tr[1:]}

		meanings, err := w.Translate(DE)

		if err != nil {
			t.Errorf("error: %s in word ", w.Transscribe())
			continue
		}

		got, _ := ReduceDE(meanings...)

		if got != expected {
			t.Errorf("expected %#v, got %#v", expected, got)
			continue
		}
	}

}
