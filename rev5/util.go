package rev5

func sliceDedupe[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	var result []T

	for _, v := range input {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
