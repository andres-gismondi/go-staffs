package practices

func SortNumbers(number []int) []int {
	nn := make([]int, 0)
	nnA := make([]int, 0)
	nn = append(nn, number[0])

	for index, n := range number[1:] {
		aux := 0
		find := false
		for j, newN := range nn {
			if n < newN && !find {
				find = true
				aux = j
				bla := nn[:j]
				nnA = append(bla, n)
			}
		}
		if !find {
			nn[index] = n
		} else {
			nn = append(nnA, nn[aux:]...)
		}
	}

	return nn
}

func SortSlice(S []int64) []int64 {
	for i := len(S); i > 0; i-- {
		for j := 1; j < i; j++ {
			if S[j-1] > S[j] {
				// swap
				intermediate := S[j]
				S[j] = S[j-1]
				S[j-1] = intermediate
			}
		}
	}
	return S
}
