package luniga

// semantic primes

/*
	The semantic primes should not change, but there meaning could and is defined inside the dictionary
*/

type Prime string

const (
	MO Prime = `@`
	MA Prime = `#`
	ME Prime = `|`
	MU Prime = `?`
	MI Prime = `&`

	SO Prime = `%`
	SA Prime = `'`
	SE Prime = `^`
	SU Prime = `[`
	SI Prime = `]`

	TO Prime = `!`
	TA Prime = `*`
	TE Prime = `=`
	TU Prime = `{`
	TI Prime = `}`

	LO Prime = `<`
	LA Prime = `>`
	LE Prime = `-`
	LU Prime = `.`
	LI Prime = `~`

	PO Prime = `:`
	PA Prime = `"`
	PE Prime = `,`
	PU Prime = `\`
	PI Prime = `/`

	KO Prime = `+`
	KA Prime = `;`
	KE Prime = `$`
	KU Prime = `(`
	KI Prime = `)`
)

var primeTransscripts = map[Prime]string{
	MO: `MO`,
	MA: `MA`,
	ME: `ME`,
	MU: `MU`,
	MI: `MI`,
	SO: `SO`,
	SA: `SA`,
	SE: `SE`,
	SU: `SU`,
	SI: `SI`,
	TO: `TO`,
	TA: `TA`,
	TE: `TE`,
	TU: `TU`,
	TI: `TI`,
	LO: `LO`,
	LA: `LA`,
	LE: `LE`,
	LU: `LU`,
	LI: `LI`,
	PO: `PO`,
	PA: `PA`,
	PE: `PE`,
	PU: `PU`,
	PI: `PI`,
	KO: `KO`,
	KA: `KA`,
	KE: `KE`,
	KU: `KU`,
	KI: `KI`,
}

func (p Prime) Transscribe() string {
	return primeTransscripts[p]
}

type Number int

const (
	FOP Number = 0
	FAP Number = 1
	FEP Number = 2
	FUP Number = 3
	FIP Number = 4
	FOS Number = 5
	FAS Number = 6
	FES Number = 7
	FUS Number = 8
	FIS Number = 9
	FOL Number = 10
	FAL Number = 100
	FEL Number = 1000
	FUL Number = 1000000
	FIL Number = 1000000000
)

var numberTransscripts = map[Number]string{
	FOP: `FOP`,
	FAP: `FAP`,
	FEP: `FEP`,
	FUP: `FUP`,
	FIP: `FIP`,
	FOS: `FOS`,
	FAS: `FAS`,
	FES: `FES`,
	FUS: `FUS`,
	FIS: `FIS`,
	FOL: `FOL`,
	FAL: `FAL`,
	FEL: `FEL`,
	FUL: `FUL`,
	FIL: `FIL`,
}

func (n Number) Transscribe() string {
	return numberTransscripts[n]
}

var pronounciation = map[rune]string{
	'M': `m`,
	'S': `ʃ`,
	'T': `t`,
	'L': `l`,
	'P': `p`,
	'K': `k`,
	'F': `f`,
	'A': `a:`,
	'O': `ɔ`,
	'E': `ɛ`,
	'U': `u:`,
	'I': `i:`,
}

// see http://de.pons.com/%C3%BCbersetzung?q=Runde&l=deen&in=ac_de&lf=de
var pronounciationAlt = map[rune]string{
	'T': `d`,
	'P': `b`,
	'K': `g`,
	'S': `ç`,
}

type Transscriber interface {
	Transscribe() string
}

func Phonetic(t Transscriber) string {
	s := t.Transscribe()

	res := ""

	for _, r := range s {
		res += pronounciation[r]
	}

	return res
}

/*

aussprache:

	M => M wie in Mama
	S => SCH wie in Schule
	T => T oder D wie in Du oder Tonne
	L => L wie in Latein
	P => P oder B wie in Papa oder Bitte
	K => K oder G wie in Kunst oder Gunst
	F => F wie in Fisch

	A => A wie in Mama
	O => O wie in Ross
  E => E wie in Recht
  U => U wie in Lust
  I => I wie in Wie


@ M O
# M A
| M E
? M U
& M I

% Sch O
' Sch A
^ Sch E
[ Sch U
] Sch I

! T O
* T A
= T E
{ T U
} T I

< L O
> L A
- L E
. L U
~ L I

: P O
“ P A
, P E
\ P U
/ P I

+ K O
; K A
( K U
) K I


was ist mit Zahlenwörter??? vielleicht alle mit F beginnen lassen

0 FOP
1 FAP
2 FEP
3 FUP
4 FIP
5 FOSch
6 FASch
7 FESch
8 FUSch
9 FISch

dann werden sie dezimal hintereinandergereiht:

112 ist FapFapFep

dann gibt es noch abkürzungen für

           10 FOL
          100 FAL
        1.000 FEL
    1.000.000 FUL
1.000.000.000 FIL

dreitausend kann man so bezeichnen

FupFopFopFop

oder

FupPoFel (3 mal tausend)


alle möglichen varianten


x PA
x PO
x PE
x PU
x PI
x TA
x TO
x TE
x TU
x TI
x KA
x KO
x KE
x KU
x KI
x MA
x MO
x ME
x MU
x MI
x SchA
x SchO
x SchE
x SchU
x SchI
x LA
x LO
x LE
o LU
x LI



P
T
K
M
Sch
L

A
O
E
U
I

*/

// grouping primes
/*
[ ]
{ }
( )
*/
