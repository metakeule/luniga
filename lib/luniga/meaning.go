package luniga

import (
	"fmt"
)

type Meaning struct {

	// a short name that stands for the meaning,
	// should be as close as possible to the core meaning
	Short string

	// the intersection of the given words reflects the meaning
	Intersection []string

	// a longer description of the meaning that should give an exact impression
	Description string

	// the stability of the Meaning:
	// -2 questionable
	// -1 uncertain
	//  0 unstable
	//  1 stable
	//  2 fixed
	Stability int
}

type PrimeMeaning struct {
	Single  Meaning
	Double  Meaning
	Tripple Meaning
	Num1    Meaning
	Num2    Meaning
	Num3    Meaning
}

type CombinedMeaning struct {
	Primes  []Transscriber // Primes and Numbers
	Meaning Meaning
}

type Language struct {
	LanguageCode  string // according to ISO 639
	PrimeMeanings map[Prime]PrimeMeaning
	// CoreMeanings     map[string]Meaning
	// CombinedMeanings []CombinedMeaning
}

type Word struct {
	Head Prime
	Tail []Transscriber
}

func (w *Word) Transscribe() string {
	s := w.Head.Transscribe()

	for _, t := range w.Tail {
		s += t.Transscribe()
	}

	return s
}

// if an error occurs, len([]Meaning) is the position of the error
func (w *Word) Translate(l *Language) ([]Meaning, error) {

	m := []Meaning{}

	if len(w.Tail) == 0 {
		m = append(m, l.PrimeMeanings[w.Head].Double)
		return m, nil
	}

	switch x := w.Tail[0].(type) {
	case Number:
		switch int(x) {
		case 1:
			m = append(m, l.PrimeMeanings[w.Head].Num1)
			if len(w.Tail) > 1 {
				newPrime, ok := w.Tail[1].(Prime)
				if !ok {
					return m, fmt.Errorf("new inner word must begin with prime")
				}
				next := &Word{Head: newPrime}
				if len(w.Tail) > 2 {
					next.Tail = w.Tail[2:]
				}
				nextMs, err := next.Translate(l)
				if err != nil {
					return m, err
				}
				m = append(m, nextMs...)
				return m, nil
			}
			return m, nil
		case 2:
			m = append(m, l.PrimeMeanings[w.Head].Num2)
			if len(w.Tail) > 1 {
				newPrime, ok := w.Tail[1].(Prime)
				if !ok {
					return m, fmt.Errorf("new inner word must begin with prime")
				}
				next := &Word{Head: newPrime}
				if len(w.Tail) > 2 {
					next.Tail = w.Tail[2:]
				}
				nextMs, err := next.Translate(l)
				if err != nil {
					return m, err
				}
				m = append(m, nextMs...)
				return m, nil
			}
			return m, nil
		case 3:
			m = append(m, l.PrimeMeanings[w.Head].Num3)
			if len(w.Tail) > 1 {
				newPrime, ok := w.Tail[1].(Prime)
				if !ok {
					return m, fmt.Errorf("new inner word must begin with prime")
				}
				next := &Word{Head: newPrime}
				if len(w.Tail) > 2 {
					next.Tail = w.Tail[2:]
				}
				nextMs, err := next.Translate(l)
				if err != nil {
					return m, err
				}
				m = append(m, nextMs...)
				return m, nil
			}
			return m, nil
		default:
			return nil, fmt.Errorf("numbers other than 1,2,3 not allowed as prime followers")
		}
	case Prime:
		if x == w.Head {
			if len(w.Tail) > 1 {
				if w.Tail[1] == w.Head {
					m = append(m, l.PrimeMeanings[w.Head].Tripple)
					if len(w.Tail) > 2 {
						newPrime, ok := w.Tail[2].(Prime)
						if !ok {
							return m, fmt.Errorf("new inner word must begin with prime")
						}
						next := &Word{Head: newPrime}
						if len(w.Tail) > 3 {
							next.Tail = w.Tail[3:]
						}
						nextMs, err := next.Translate(l)
						if err != nil {
							return m, err
						}
						m = append(m, nextMs...)
						return m, nil
					}
					return m, nil
				}

				m = append(m, l.PrimeMeanings[w.Head].Double)
				newPrime, ok := w.Tail[1].(Prime)
				if !ok {
					return m, fmt.Errorf("new inner word must begin with prime")
				}
				next := &Word{Head: newPrime}
				if len(w.Tail) > 2 {
					next.Tail = w.Tail[2:]
				}
				nextMs, err := next.Translate(l)
				if err != nil {
					return m, err
				}
				m = append(m, nextMs...)
				return m, nil
			}
			m = append(m, l.PrimeMeanings[w.Head].Double)
			return m, nil
		}

		m = append(m, l.PrimeMeanings[w.Head].Single)

		newPrime, ok := w.Tail[0].(Prime)
		if !ok {
			return m, fmt.Errorf("new inner word must begin with prime")
		}

		next := &Word{Head: newPrime}
		if len(w.Tail) > 1 {
			next.Tail = w.Tail[1:]
		}
		nextMs, err := next.Translate(l)
		if err != nil {
			return m, err
		}
		m = append(m, nextMs...)
		return m, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", w.Tail[0])
	}

	panic("unreachable")
}

type Sentence struct {
	Type  []Transscriber
	words []*Word
}

type Paragraph []Sentence
