package port

func Next() int {
	for p := MIN; p <= MAX; p++ {
		if Available(p) {
			return p
		}
	}
	return -1
}