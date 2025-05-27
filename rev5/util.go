package rev5

func sliceDedupe[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	var result []T

	for _, v := range s {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func sliceToMap[T comparable](s []T) map[T]T {
	out := make(map[T]T)
	for _, v := range s {
		out[v] = v
	}
	return out
}
