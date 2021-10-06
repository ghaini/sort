package sort

func (s *Sort) merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func (s *Sort) MergeSort() []int {
	if len(s.arr) < 2 {
		return s.arr
	}

	arr := s.arr
	s.arr = arr[:len(arr)/2]
	first := s.MergeSort()
	s.arr = arr[len(arr)/2:]
	second := s.MergeSort()
	return s.merge(first, second)
}
