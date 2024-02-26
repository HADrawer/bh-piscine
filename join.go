package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			str = Concat(str, strs[i]) + sep
		} else {
			str = Concat(str, strs[i])
		}

	}
	return str
}
