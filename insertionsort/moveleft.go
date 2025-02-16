package insertionsort

// MoveLeft bewegt das Element an der gegebenen Position nach links,
// bis es an der richtigen Position ist.
func MoveLeft(list []int, pos int) {
	// TODO
	for i := pos; i > 0; i-- {
		if list[i-1] > list[i] {

			list[i], list[i-1] = list[i-1], list[i] //vertauscht das Element der Liste an der Stelle i mit dem Element an der Stelle i-1

		} else {
			i = 0
		}

	}
}
