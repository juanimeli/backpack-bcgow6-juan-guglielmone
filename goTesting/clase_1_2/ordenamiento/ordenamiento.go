package ordenamiento

import "sort"

func OrderSlice(incomeSlice []int) []int {

	sort.Ints(sort.IntSlice(incomeSlice))

	return incomeSlice
}
