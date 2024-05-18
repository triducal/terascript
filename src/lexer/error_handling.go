package lexer

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
	// last line
	l.lineMap = append(l.lineMap, [2]int{begin, idx + 1})
}

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

// ErrorLine (pos) returns lineNum, column, errorLine
func (l *Lexer) ErrorLine(pos int) (int, int, string) {
	lineNum, begin, end := l.linePosition(pos)
	errorLine := l.source[begin:end]
	column := pos - begin + 1
	return lineNum, column, string(errorLine)
}
