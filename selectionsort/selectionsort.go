package selectionsort

// SelectionSort sortiert eine Liste von Zahlen mit dem SelectionSort-Algorithmus.
func SelectionSort(list []int) {
	// TODO
	for i := 0; i < len(list); i++ {
		smallest := SmallestPos(list[i:]) + i

		list[smallest], list[i] = list[i], list[smallest]

	}
}
