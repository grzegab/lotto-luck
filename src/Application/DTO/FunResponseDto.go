package DTO

type FunResponseDto struct {
	howManyUntil  int
	savedFilename string
}

func (f *FunResponseDto) HowManyUntil() int {
	return f.howManyUntil
}

func (f *FunResponseDto) SetHowManyUntil(howManyUntil int) {
	f.howManyUntil = howManyUntil
}

func (f *FunResponseDto) SavedFilename() string {
	return f.savedFilename
}

func (f *FunResponseDto) SetSavedFilename(savedFilename string) {
	f.savedFilename = savedFilename
}
