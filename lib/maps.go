package lib

type Map[K comparable, V comparable] struct  {}

func (m Map[K, V]) AreEqual(map1 map[K]V, map2 map[K]V) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key1, val1 := range map1 {
		val2, ok := map2[key1]
		if !ok {
			return false // Key not found in map2
		}
		if val1 != val2 {
			return false // Values are not equal
		}
	}

	return true
}
