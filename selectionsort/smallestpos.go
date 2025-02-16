package selectionsort

// SmallestPos gibt die Position des kleinsten Elements in der Liste zurück.
func SmallestPos(list []int) int {
	smallest := 0
	// TODO
	for i := 0; i < len(list); i++ {
		if list[smallest] > list[i] {
			smallest = i
		}
	}
	return smallest
}
