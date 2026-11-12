package clock

import "fmt"

type Clock struct {
	minutes int
}

const (
	minutesInHour = 60
	minutesInDay  = 24 * minutesInHour
)

func normalize(minutes int) int {
	if minutes < 0 {
		return (minutes % minutesInDay) + minutesInDay
	}
	if minutes >= minutesInDay {
		return minutes % minutesInDay
	}
	return minutes
}

func New(h, m int) Clock {
	return Clock{minutes: normalize(minutesInHour*h + m)}
}

func (c Clock) Add(m int) Clock {
	return Clock{minutes: normalize(c.minutes + m)}
}

func (c Clock) Subtract(m int) Clock {
	return Clock{minutes: normalize(c.minutes - m)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
