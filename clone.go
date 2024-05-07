package structure

func CloneMap[K comparable, V any](originalMap map[K]V) map[K]V {
	cloned := make(map[K]V)

	for key, value := range originalMap {
		cloned[key] = value
	}

	return cloned
}
