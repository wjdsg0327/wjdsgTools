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
	return 2
}
