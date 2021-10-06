package sort

func (s *Sort) QuickSort() []int {
	s.quick(0, len(s.arr)-1)
	return s.arr
}

func (s *Sort) quick(low, high int) {
	if low < high {
		p := s.partition(low, high)
		s.quick(low, p-1)
		s.quick(p+1, high)
	}
}

func (s *Sort) partition(low, high int) int {
	pivot := s.arr[high]
	i := low
	for j := low; j < high; j++ {
		if s.arr[j] < pivot {
			s.swap(i, j)
			i++
		}
	}
	s.swap(i, high)
	return i
}