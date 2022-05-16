package block

type block interface {
	GetChar() string
	SetChar()
	GetStatus() string
	SetStatus()
}

type boardBlock struct {
	char, status string
}

func (b boardBlock) GetChar() string {
	return b.char
}

func (b boardBlock) SetChar(char string) {
	b.char = char
}

func (b boardBlock) GetStatus() string {
	return b.status
}

func (b boardBlock) SetStatus(stat string) {
	b.status = stat
}
