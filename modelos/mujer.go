package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Pensar() {
	m.pensando = true
}

func (m *Mujer) Comer() {
	m.comiendo = true
}

func (m *Mujer) Genero() string {
	return "Femenino"
}

func (m *Mujer) EstaVivo() bool {
	return m.vivo
}
