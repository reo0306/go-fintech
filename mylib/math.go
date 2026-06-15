package mylib

func Average(v []int) int {
	total := 0
	for _, i := range v {
		total += i
	}
	return int(total / len(v))
}
