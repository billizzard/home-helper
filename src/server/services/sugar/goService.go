package sugar

func InArray(needle string, arr []string) bool {
	for _, v := range arr {
		if needle == v {
			return true
		}
	}

	return false
}
