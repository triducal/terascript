package lexer

type Lexer struct {
	pos     int  // current position in source (points to ch)
	readPos int  // current reading position in source (position after ch)
	ch      rune // current character
	source  []rune

	// used for locating error positions in `buildLineMap()`
	lineMap [][2]int // array of [begin, end] pairs
}

func New(in string) *Lexer {
	l := &Lexer{source: []rune(in)}
}

func (l *Lexer) buildLineMap() {
	begin := 0
	idx := 0
	for i, ch := range l.source {
		idx = i
		if ch == '\n' {
			l.lineMap = append(l.lineMap, [2]int{begin, idx})
			begin = idx + 1
		}
	}
	// adds last line that doesnt have a newline character
	l.lineMap = append(l.lineMap, [2]int{begin, idx + 1})
}

func (l *Lexer) CurrentPosition() int {
	return l.pos
}

// Helper function for finding line number and start and end position of the line containing `pos`
//
//	lineNum, begin, end := linePosition(pos)
func (l *Lexer) linePosition(pos int) (int, int, int) {
	idx := 0
	begin := 0
	end := 0
	for i, tuple := range l.lineMap {
		idx = i
		begin, end = tuple[0], tuple[1]
		if pos >= begin && pos <= end {
			break
		}
	}
	lineNum := idx + 1
	return lineNum, begin, end
}

// Helper function used for reporting info of an error occuring at `pos`
//
//	lineNum, collumn, lineText := ErrorLine(pos)
func (l *Lexer) ErrorLine(pos int) (int, int, string) {
	lineNum, begin, end := l.linePosition(pos)
	errorLine := l.source[begin:end]
	column := pos - begin + 1
	return lineNum, column, string(errorLine)
}

func (l *Lexer) newToken(tokenType)
