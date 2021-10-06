package sort

type Sort struct {
	arr []int
}

func NewSort(arr []int) Sort {
	return Sort{
		arr: arr,
	}
}

func (s *Sort) swap(from, to int) {
	if from != to {
		s.arr[from], s.arr[to] = s.arr[to], s.arr[from]
	}
}
