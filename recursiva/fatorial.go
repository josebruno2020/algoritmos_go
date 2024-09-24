package recursiva

func Fatorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * Fatorial(num-1)
}
