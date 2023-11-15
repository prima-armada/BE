package kriteria

func KriteriaKandidat(test string) (nilai int) {
	if test == "sangat baik" {
		nilai = 4
		return nilai
	} else if test == "baik" {
		nilai = 3
		return nilai
	} else if test == "cukup" {
		nilai = 2
		return nilai
	} else if test == "kurang" {
		nilai = 1
		return nilai
	}
	nilai = 0
	return nilai
}
