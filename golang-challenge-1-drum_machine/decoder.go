package drum

import "fmt"

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	p := &Pattern{}
	return p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	version string
	tempo   int
	tracks  []Track
}

func (p *Pattern) String() string {
	str := ""

	str += "Saved with HW Version: " + p.version
	str += fmt.Sprintf("Tempo: %d", p.tempo)

	for i, t := range p.tracks {
		if i > 0 {
			str += "\n"
		}

		str += fmt.Sprintf(t)
	}

	return str
}

type Track struct {
	id    int
	name  string
	steps []Step
}

func (p *Track) String() string {
	str := ""

	str += fmt.Sprintf("(%s) %s\t", p.id, p.name)

	for i, s := range p.steps {
		if i%4 == 0 {
			str += "|"
		}

		str += fmt.Sprint(s)
	}

	str += "|"

	return str
}

type Step struct {
	note byte
}

func (p *Step) String() string {
	str := ""

	str += string(p.note)

	return str
}
