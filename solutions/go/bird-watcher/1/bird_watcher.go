package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for _, b := range birdsPerDay {
        sum += b
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
    for i:= 7*(week-1); i < 7*week; i++ {
        sum += birdsPerDay[i]
    }
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, b := range birdsPerDay {
        birdsPerDay[i] = b + (i+1) % 2
    }
    return birdsPerDay
}
