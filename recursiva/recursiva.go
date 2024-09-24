package recursiva

func Recursiva(item int) {
	// fmt.Println(item)
	if item <= 0 {
		return
	}
	Recursiva(item - 1)
}
