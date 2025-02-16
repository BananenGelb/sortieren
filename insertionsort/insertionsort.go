package insertionsort

// InsertionSort sortiert die gegebene Liste mit dem Insertion-Sort-Algorithmus.
func InsertionSort(list []int) {
	// TODO
	for i, _ := range list {
		MoveLeft(list, i)

	}
}
