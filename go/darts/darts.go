// Darts exercise of Exercism.io Go track.
package darts

// Compute score of the dart position.
func Score(x, y float64) int {
	// Squared distance is used for optimization reason:
	// there's no point in computing square root if it's possible
	// to check against squared distance.

	scores := []struct {
		distance float64
		score    int
	}{ // order does matter
		{1.0 * 1.0, 10},
		{5.0 * 5.0, 5},
		{10.0 * 10.0, 1}}

	d := x*x + y*y

	for _, s := range scores {
		if d <= s.distance {
			return s.score
		}
	}

	return 0
}
