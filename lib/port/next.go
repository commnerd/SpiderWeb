package port

func Next() Port {
	for p := MIN; p <= MAX; p++ {
		if Available(Port(p)) {
			return Port(p)
		}
	}
	return -1
}