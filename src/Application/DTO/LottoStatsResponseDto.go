package DTO

type LottoStats struct {
	resultAffected int
	savedFilename  string
}

func (l *LottoStats) SetResultAffected(count int) {
	l.resultAffected = count
}

func (l *LottoStats) GetResultAffected() int {
	return l.resultAffected
}

func (l *LottoStats) SetSavedFilename(filename string) {
	l.savedFilename = filename
}

func (l *LottoStats) GetSavedFilename() string {
	return l.savedFilename
}
