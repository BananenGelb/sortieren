package bubblesort

// BubbleSort sortiert the gegebene Liste mit dem Bubble-Sort-Algorithmus.
func BubbleSort(list []int) {
	// TODO
	max := len(list)
	for i := 0; i < len(list); i++ {
		if BubbleUp(list[:max]) == false {
			i = len(list)

		}
		max--
	}
}
