package clock

import "fmt"

type Clock struct {
	minutes int
}

const {
	minutesInHour = 60
	minutesInDay = 24 * minutesInHour
}

func (Clock *c) normalize()  {
	if (c.minutes < 0) {
		c.minutes = c.minutes + minutesInDay
	}
	if (c.minutes > minutesInDay) {
		c.minutes = c.minutes % minutesInDay
	}
}

func New(h, m int) Clock {
	return Clock{minutes: minutesInHour*h + m}.normalize()
}

func (c Clock) Add(m int) Clock {
	return Clock{minutes: c.minutes + m}.normalize()
}

func (c Clock) Subtract(m int) Clock {
	return Clock{minutes: c.minutes - m}.normalize()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
