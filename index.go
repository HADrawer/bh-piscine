package piscine

func Index(s string, toFind string) int {
	lens := len(s)
	lentofind := len(toFind)

	if lens < lentofind {
		return -1
	}
	for i := 0; i <= lens-1; i++ {
		if lentofind+i < lens {
			if toFind == s[i:lentofind+i] {
				return i
			}
		} else {
			break
		}
	}
	return -1
}
