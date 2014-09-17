package v0_1

import (
	"github.com/metakeule/luniga/lib/luniga"
)

// the meanings are
var DE = &luniga.Language{
	LanguageCode: "de",
	PrimeMeanings: map[luniga.Prime]luniga.PrimeMeaning{

		luniga.MO: luniga.PrimeMeaning{
			Single: luniga.Meaning{
				Short: "MENSCH",
			},
			Double: luniga.Meaning{
				Short: "TIER",
			},
			Tripple: luniga.Meaning{
				Short: "PFLANZE",
			},
			Num1: luniga.Meaning{
				Short: "ICH",
			},
			Num2: luniga.Meaning{
				Short: "DU",
			},
			Num3: luniga.Meaning{
				Short: "ANDERE MENSCHEN",
			},
		},

		luniga.MA: luniga.PrimeMeaning{
			Single: luniga.Meaning{
				Short: "BILD",
			},
			Double: luniga.Meaning{
				Short: "ZEICHEN",
			},
			Tripple: luniga.Meaning{
				Short: "METAPHER",
			},
			Num1: luniga.Meaning{
				Short: "SKIZZE",
			},
			Num2: luniga.Meaning{
				Short: "FOTO",
			},
			Num3: luniga.Meaning{
				Short: "MODELL",
			},
		},
	},
}

func ReduceDE(m ...luniga.Meaning) (s string, wasReduced bool) {
	if len(m) == 2 {
		if m[0].Short == DE.PrimeMeanings[luniga.MO].Num1.Short && m[1].Short == DE.PrimeMeanings[luniga.MA].Double.Short {
			return "ich heiße", true
		}
	}

	r := ""

	for _, mean := range m {
		r += mean.Short
	}

	return r, false
}
