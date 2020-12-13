// Insertion sort
package insort

// Int insertion sort
func IntInsort(keys []int) []int {
	for i := 1; i < len(keys); i++ {
		j := i
		for j > 0 {
			if keys[j-1] > keys[j] {
				keys[j-1], keys[j] = keys[j], keys[j-1]
			}
			j = j - 1
		}
	}

	return keys
}

// Float64 insertion sort
func Float64Insort(keys []float64) []float64 {
	for i := 1; i < len(keys); i++ {
		j := i
		for j > 0 {
			if keys[j-1] > keys[j] {
				keys[j-1], keys[j] = keys[j], keys[j-1]
			}
			j = j - 1
		}
	}

	return keys
}
