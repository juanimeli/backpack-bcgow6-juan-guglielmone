package ordenamiento

import "sort"

func OrderSlice(incomeSlice []int) []int {

	sort.Sort(sort.IntSlice(incomeSlice))

	return incomeSlice
}
