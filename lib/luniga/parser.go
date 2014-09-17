package luniga

/*
wichtig ist, dass wir die association und greedyness definieren


klammernde zeichen

[ SU
] SI
{ TU
} TI
( KU alles was vor einem KU steht bezieht sich auf den eingeschlossenen bereich
) KI bis zum nächsten KI

operatoren (muss immer rechts und links was stehen)

= TE
- LE
: PO
, PE
. LU
> LO
< LA
| ME
& MI
/ PI

prinzipiell gilt: steht ein zeichen vor einem KU so bezieht es sich auf den eingeschlossenen bereich
bis zum nächsten KI
das ist insbesondere für folgende zeichen interessant

" PA
' SA


also "(hallo welt)
 Zitat: "Hallo Welt"


spezialfall

\ PU wenn von einem Leerzeichen gefolgt, drückt es das ende (eines satzes/abschnitts/kapitel) aus
     ansonsten den Anfang. wenn es den anfang ausdrückt beschreibt alles vorangestellte den satz/abschnitt/kapitel

z.B.

;\@+\

ein \ gefolgt von einem Leerzeichen beendet einen Satz/eine Aussage
ansonsten leitet ein \ eine Aussage ein und alles was davor steht beschreibt die Aussage.

d.h. obiges bedeutet: "Der Mensch ist gut!. Dies ist eine wahre Aussage."

das ; beschreibt das nachfolgende \, d.h. der Satz ist wahr

im Gegensatz dazu

\;@+\

"Der wahre Mensch ist gut." oder "Die Wahrheit ist: Der Mensch ist gut." oder "In Wahrheit ist der Mensch gut."
oder "Die Wahrheit des Menschen ist gut sein"

solche speziellen associationen und greedyness müssen für folgende zeichen definiert werden:


*/

//  	bracket

type Bracket struct {
	Start Prime
	End   Prime
}

/*
[ SU
] SI
{ TU
} TI
( KU alles was vor einem KU steht bezieht sich auf den eingeschlossenen bereich
) KI
*/

var SUSI = &Bracket{SU, SI}
var TUTI = &Bracket{TU, TI}
var KUKI = &Bracket{KU, KI}

var primeToBracket = map[Prime]*Bracket{
	SU: SUSI,
	SI: SUSI,
	TU: TUTI,
	TI: TUTI,
	KU: KUKI,
	KI: KUKI,
}

/*
= TE
- LE
: PO
, PE
. LU
> LO
< LA
| ME
& MI
/ PI
*/

var operators = map[Prime]struct{}{
	TE: struct{}{},
	LE: struct{}{},
	PO: struct{}{},
	PE: struct{}{},
	LU: struct{}{},
	LO: struct{}{},
	LA: struct{}{},
	ME: struct{}{},
	MI: struct{}{},
	PI: struct{}{},
}

func (p Prime) IsOperator() bool {
	_, has := operators[p]
	return has
}

func (p Prime) Bracket() *Bracket {
	b, has := primeToBracket[p]
	if !has {
		return nil
	}
	return b
}

func Parse(s string) []Transscriber {
	res := []Transscriber{}

	for _, r := range s {
		switch r {
		case '@':
			res = append(res, MO)
		case '#':
			res = append(res, MA)
		case '|':
			res = append(res, ME)
		case '?':
			res = append(res, MU)
		case '&':
			res = append(res, MI)
		case '%':
			res = append(res, SO)
		case '\'':
			res = append(res, SA)
		case '^':
			res = append(res, SE)
		case '[':
			res = append(res, SU)
		case ']':
			res = append(res, SI)
		case '!':
			res = append(res, TO)
		case '*':
			res = append(res, TA)
		case '=':
			res = append(res, TE)
		case '{':
			res = append(res, TU)
		case '}':
			res = append(res, TI)
		case '<':
			res = append(res, LO)
		case '>':
			res = append(res, LA)
		case '-':
			res = append(res, LE)
		case '.':
			res = append(res, LU)
		case '~':
			res = append(res, LI)
		case ':':
			res = append(res, PO)
		case '"':
			res = append(res, PA)
		case ',':
			res = append(res, PE)
		case '\\':
			res = append(res, PU)
		case '/':
			res = append(res, PI)
		case '+':
			res = append(res, KO)
		case ';':
			res = append(res, KA)
		case '$':
			res = append(res, KE)
		case '(':
			res = append(res, KU)
		case ')':
			res = append(res, KI)
		case '0':
			res = append(res, FOP)
		case '1':
			res = append(res, FAP)
		case '2':
			res = append(res, FEP)
		case '3':
			res = append(res, FUP)
		case '4':
			res = append(res, FIP)
		case '5':
			res = append(res, FOS)
		case '6':
			res = append(res, FAS)
		case '7':
			res = append(res, FES)
		case '8':
			res = append(res, FUS)
		case '9':
			res = append(res, FIS)
		}
	}

	return res
}
