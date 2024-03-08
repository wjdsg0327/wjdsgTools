package DateUtils

type DateField struct {
}

func (DateField) YEAR() int {
	return 1
}
func (DateField) MOUTH() int {
	return 2
}
func (DateField) DAY() int {
	return 3
}

func (DateField) FormatYMDHMS() string {
	return "2006-01-02 15:04:05"
}

func (DateField) FormatYMD() string {
	return "2006-01-02"
}
