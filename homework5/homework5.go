package homework5

func Map(data []int, action func(int) int) []int {
	if data == nil {
		return nil
	}
	mapped := make([]int, 0, len(data))
	for _, v := range data {
		mapped = append(mapped, action(v))
	}

	return mapped
}

func Filter(data []int, action func(int) bool) []int {
	if data == nil {
		return nil
	}
	filtered := make([]int, 0, len(data))
	for _, v := range data {
		if action(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	if data == nil {
		return initial
	}
	for _, v := range data {
		initial = action(v, initial)
	}
	return initial
}
